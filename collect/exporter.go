package collect

import (
	"bytes"
	"encoding/json"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
	"os"
	"sync"
	"time"

	"github.com/ArisAachen/experience/abstract"
	"github.com/ArisAachen/experience/define"
	"github.com/ArisAachen/experience/launch"
	"github.com/godbus/dbus"
)

type DBusCollector struct {
	// entry is a chan for post message
	entry chan define.AppEntry
	// login/out shutdown event
	logon chan define.LogInfo
	// sysBus is system bus to export dbus obj
	sysBus *dbus.Conn
	// lau save launch
	lau *launch.Launch

	// lock and save config
	lock sync.Mutex
	define.SysCfg
}

// newDBusCollector create dbus collector
func newDBusCollector(lau *launch.Launch) *DBusCollector {
	bus := &DBusCollector{
		// make app chan
		entry: make(chan define.AppEntry),
		logon: make(chan define.LogInfo),
		// launch
		lau: lau,
	}
	return bus
}

// Init init dbus property
func (bus *DBusCollector) Init() error {
	// create system bus
	var err error
	bus.sysBus, err = dbus.SystemBus()
	if err != nil {
		logger.Warningf("create system bus failed, err: %v", err)
		return err
	}
	// export dbus method
	err = bus.sysBus.Export(bus, define.ServicePath, define.DbusInterface)
	if err != nil {
		logger.Warningf("export method failed, err: %v", err)
		return err
	}
	// request dbus name
	_, err = bus.sysBus.RequestName(define.ServiceName, dbus.NameFlagDoNotQueue)
	if err != nil {
		logger.Warningf("request name failed, err: %v", err)
		return err
	}
	// load config file from path, it is ok, of file cant loaded
	err = bus.LoadFromFile(bus.GetFileName())
	if err != nil {
		logger.Warningf("cant load file from path: %v", err)
	}
	// if user-exp state is not true, cant post data
	if !bus.GetUserExp() {
		// use strict rule to disable post
		bus.lau.GetController().Invoke(define.StrictRule)
	}

	logger.Debug("export dbus obj success")
	return nil
}

// Collect use to collect message
func (bus *DBusCollector) Collect(que abstract.BaseQueue) error {
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
			req.Msg = string(data)
			req.Pri = define.SimpleRequest
			req.Rule = define.LooseRule
		case log := <-bus.logon:
			// marshal app info
			data, err := json.Marshal(&log)
			if err != nil {
				logger.Warningf("marshal logon data failed, err: %v")
				continue
			}
			// request message
			req.Msg = string(data)
			req.Pri = define.LogInOutRequest
			req.Rule = define.LooseRule
		}
		// push data to queue
		go que.Push(define.DataBaseItemQueue, bus, req)
	}
}

// Handler handle write to database result
// at this time, just drop data if failed
func (bus *DBusCollector) Handler(base abstract.BaseQueue, controller abstract.BaseController, result define.WriteResult) {
	return
}

// GetInterface write database table
func (bus *DBusCollector) GetInterface() string {
	return "v2/report/unification"
}

// SendAppStateData use to collect open/close app message,
// this method has deprecated, use dde.dock to collect use message
func (bus *DBusCollector) SendAppStateData(msg string, path string, name string, id string) {
	return
}

// SendAppInstallData collect app un/install app message
func (bus *DBusCollector) SendAppInstallData(msg string, path string, name string, id string) {
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
}

// SendLogonData collect logon data
func (bus *DBusCollector) SendLogonData(msg string) {
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
		return
	}
	// in case block
	go func() {
		bus.logon <- log
	}()
}

// IsEnabled check now collector state
func (bus *DBusCollector) IsEnabled() bool {
	return bus.GetUserExp()
}

// Enable enable user-exp state
func (bus *DBusCollector) Enable(enabled bool) *dbus.Error {
	// check if now state is the same
	if enabled == bus.GetUserExp() {
		logger.Debugf("user exp state is now already %v", enabled)
	}
	// if user-exp was closed, and now open this file, should release rule
	if enabled {
		bus.lau.GetController().Release(define.StrictRule)
	}
	// save user exp
	bus.UserExp = enabled
	// TODO
	go func() {
		err := bus.SaveToFile(bus.GetFileName())
		if err != nil {
			logger.Warningf("save user-exp state to file failed, err: %v", err)
			return
		}
	}()
	return nil
}

// SaveToFile save protobuf config to file
func (bus *DBusCollector) SaveToFile(filename string) error {
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

// GetFileName get config file name
func (bus *DBusCollector) GetFileName() string {
	return define.SysCfgFile
}

// LoadFromFile load protobuf config from file
func (bus *DBusCollector) LoadFromFile(filename string) error {
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
