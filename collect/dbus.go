package collect

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"os"
	"sync"
	"time"

	"github.com/ArisAachen/experience/abstract"
	"github.com/ArisAachen/experience/define"
	"github.com/godbus/dbus"
	"github.com/golang/protobuf/proto"
	networkmanager "github.com/linuxdeepin/go-dbus-factory/org.freedesktop.networkmanager"
	"pkg.deepin.io/lib/dbusutil"
)

//go:generate dbusutil-gen em -type DBusModule dbus.go

type DBusModule struct {
	// entry is a chan for post message
	entry chan define.AppEntry
	// login/out shutdown event
	logon chan define.LogInfo

	// sysBus is system bus to export dbus obj
	sysService *dbusutil.Service
	sysBus     *dbus.Conn
	sigLoop    *dbusutil.SignalLoop

	// state net available state
	state uint32

	controller abstract.BaseController

	// lock and save config
	lock sync.Mutex
	define.SysCfg
}

// NewDBusModule create dbus module
func NewDBusModule(sysService *dbusutil.Service) *DBusModule {
	bus := &DBusModule{
		// make app chan
		sysService: sysService,
		entry:      make(chan define.AppEntry),
		logon:      make(chan define.LogInfo),
	}
	return bus
}

// Init init dbus property
func (bus *DBusModule) Init() error {
	// export dbus method
	err := bus.sysService.Export(define.ServicePath, bus)
	if err != nil {
		logger.Warningf("export method failed, err: %v", err)
		return err
	}
	// request dbus name
	err = bus.sysService.RequestName(define.ServiceName)
	if err != nil {
		logger.Warningf("request name failed, err: %v", err)
		return err
	}
	// create sig loop to monitor network-manager state
	bus.sigLoop = dbusutil.NewSignalLoop(bus.sysBus, 10)
	// load config file from path, it is ok, of file cant loaded
	err = bus.LoadFromFile(bus.GetConfigPath())
	if err != nil {
		logger.Warningf("cant load file from path: %v", err)
	}
	logger.Debug("export dbus obj success")
	return nil
}

func (bus *DBusModule) GetInterfaceName() string {
	return define.ServiceName
}

func (bus *DBusModule) SetController(ctl abstract.BaseController) {
	bus.controller = ctl
}

// Collect use to collect message
func (bus *DBusModule) Collect(que abstract.BaseQueue) {
	// use collect to wait open/close message
	for {
		var req define.RequestMsg
		select {
		case entry := <-bus.entry:
			// marshal app info
			data, err := json.Marshal(&entry)
			if err != nil {
				logger.Warningf("marshal app data failed, err: %v")
				continue
			}
			// request message
			req.Msg = []string{string(data)}
			req.Pri = define.SimpleRequest
			req.Rule = define.LooseRule
			// push data to queue
			go que.Push(define.DataBaseItemQueue, bus, req)
		case log := <-bus.logon:
			// marshal app info
			data, err := json.Marshal(&log)
			if err != nil {
				logger.Warningf("marshal logon data failed, err: %v")
				continue
			}
			// request message
			req.Msg = []string{string(data)}
			req.Pri = define.LogInOutRequest
			req.Rule = define.LooseRule
			// push data to queue
			go que.Push(define.DataBaseItemQueue, bus, req)
		}
	}
}

func (bus *DBusModule) GetCollectName() string {
	return "dbus"
}

// Handler handle write to database result
// at this time, just drop data if failed
func (bus *DBusModule) Handler(base abstract.BaseQueue, crypt abstract.BaseCryptor, controller abstract.BaseController, result define.WriteResult) {
	return
}

// GetInterface write database table
func (bus *DBusModule) GetInterface() string {
	return "v2/report/unification"
}

// SendAppStateData use to collect open/close app message,
// this method has deprecated, use dde.dock to collect use message
func (bus *DBusModule) SendAppStateData(msg string, path string, name string, id string) *dbus.Error {
	return nil
}

// SendAppInstallData collect app un/install app message
func (bus *DBusModule) SendAppInstallData(msg string, path string, name string, id string) *dbus.Error {
	// create now open/close app
	entry := define.AppEntry{
		Time:    time.Now().UnixNano(),
		AppName: name,
		PkgName: id,
	}
	// in case block here
	go func() {
		bus.entry <- entry
	}()
	return nil
}

// SendLogonData collect logon data
func (bus *DBusModule) SendLogonData(msg string) *dbus.Error {
	// save time
	var log define.LogInfo
	log.Time = time.Now().UnixNano()
	// check message type
	switch define.LogEvent(msg) {
	case define.LoginEvent:
		log.Tid = define.LoginTid
	case define.LogOutEvent:
		log.Tid = define.LogoutTid
	default:
		logger.Warningf("unknown login event: %v", msg)
		return nil
	}
	// in case block
	go func() {
		bus.logon <- log
	}()
	return nil
}

// IsEnabled check now collector state
func (bus *DBusModule) IsEnabled() (bool, *dbus.Error) {
	return bus.GetUserExp(), nil
}

// Enable enable user-exp state
func (bus *DBusModule) Enable(enabled bool) *dbus.Error {
	// check if now state is the same
	if enabled == bus.GetUserExp() {
		logger.Debugf("user exp state is now already %v", enabled)
	}
	// when open user-exp, should release rule
	// or when close user-exp, should invoke rule
	if enabled {
		bus.controller.Release(define.StrictRule)
	} else {
		bus.controller.Invoke(define.StrictRule)
	}
	// save user exp
	bus.UserExp = enabled
	// TODO
	go func() {
		err := bus.SaveToFile(bus.GetConfigPath())
		if err != nil {
			logger.Warningf("save user-exp state to file failed, err: %v", err)
			return
		}
	}()
	return nil
}

// SaveToFile save protobuf config to file
func (bus *DBusModule) SaveToFile(filename string) error {
	// lock op
	bus.lock.Lock()
	defer bus.lock.Unlock()

	// check and create file
	fObj, err := os.Create(filename)
	if err != nil {
		return err
	}
	// marshal config to buf
	buf, err := proto.Marshal(&bus.SysCfg)
	if err != nil {
		return err
	}
	// bytes writer
	byBuf := bytes.NewBuffer(buf)
	_, err = byBuf.WriteTo(fObj)
	if err != nil {
		return err
	}
	return nil
}

// GetConfigPath get config file name
func (bus *DBusModule) GetConfigPath() string {
	return define.SysCfgFile
}

// LoadFromFile load protobuf config from file
func (bus *DBusModule) LoadFromFile(filename string) error {
	// lock op
	bus.lock.Lock()
	defer bus.lock.Unlock()
	// read file and unmarshal
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	err = proto.Unmarshal(buf, &bus.SysCfg)
	if err != nil {
		return err
	}
	return nil
}

// Wait use to control web sender op
// when user-exp is not up, or current web is unavailable
// dont try to send data, so use strict rule to block send
func (bus *DBusModule) Wait(controller abstract.BaseController) {
	// if user-exp state is not true, cant post data
	if !bus.GetUserExp() {
		// use strict rule to disable post
		bus.controller.Invoke(define.StrictRule)
	}
	// create network manager use to block data
	nm := networkmanager.NewManager(bus.sysBus)
	nm.InitSignalExt(bus.sigLoop, true)
	bus.sigLoop.Start()
	// monitor connectivity state
	err := nm.Connectivity().ConnectChanged(func(hasValue bool, value uint32) {
		// check if has value
		if !hasValue {
			logger.Warning("connectivity has no value")
			return
		}
		// save state and decide rule
		bus.save(value, controller)
	})
	// check if can monitor, if cant also ok
	if err != nil {
		logger.Warningf("monitor connectivity state failed, err: %v", err)
	}
	// get first connectivity to start rule
	state, err := nm.Connectivity().Get(0)
	if err != nil {
		logger.Warningf("get connectivity failed, err: %v", err)
		return
	}
	// when first time is not 4, block sent
	if state != 4 {
		bus.save(state, controller)
	}
}

// save save current state, and control rule
func (bus *DBusModule) save(state uint32, controller abstract.BaseController) {
	// check if state is the same
	if bus.state == state {
		return
	}
	// save current state
	bus.state = state
	// check state, 4 means network ok
	if state == 4 {
		controller.Release(define.StrictRule)
	} else {
		controller.Invoke(define.StrictRule)
	}
}
