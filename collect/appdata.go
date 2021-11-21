package collect

import (
	"encoding/json"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/ArisAachen/experience/abstract"
	"github.com/ArisAachen/experience/define"
	"github.com/godbus/dbus"
	"github.com/linuxdeepin/go-dbus-factory/com.deepin.dde.daemon.dock"
	"pkg.deepin.io/lib/dbusutil"
)

type appCollectorItem struct {
	// items use to save dock map[entry]desktop
	items map[string]define.AppEntry
	// lock
	lock sync.Mutex
	// dbus
	ddeDock dock.Dock
	sesBus  *dbus.Conn
	sigLoop *dbusutil.SignalLoop
}

// NewAppCollector create app collector
func NewAppCollector() *appCollectorItem {
	collector := &appCollectorItem{
		items: make(map[string]define.AppEntry),
	}
	return collector
}

// Init init dbus property
func (app *appCollectorItem) Init() error {
	// create dbus
	sesBus, err := dbus.SessionBus()
	if err != nil {
		logger.Warningf("get session bus failed, err: %v", err)
		return err
	}
	// save session bus
	app.sesBus = sesBus
	// save sig loop
	app.sigLoop = dbusutil.NewSignalLoop(sesBus, 10)
	// init dock
	app.ddeDock = dock.NewDock(sesBus)
	app.ddeDock.InitSignalExt(app.sigLoop, true)
	app.sigLoop.Start()
	logger.Debug("app collect init successfully")
	return nil
}

// Collect use to collect message
func (app *appCollectorItem) Collect(que abstract.BaseQueue) {
	// check if dock obj exist
	if app.ddeDock == nil {
		logger.Warning("dock init failed")
		return
	}
	// use to monitor items added
	// sometimes user send some quick-start app to dock
	// will also monitor this message
	// so should also monitor each entry's active state
	_, err := app.ddeDock.ConnectEntryAdded(func(path dbus.ObjectPath, index int32) {
		go app.monitor(que, path)
	})
	// use to monitor items close,
	// sometimes some app is closed by command,
	// item will not inactive, but remove directly
	_, err = app.ddeDock.ConnectEntryRemoved(func(entryId string) {
		// TODO: if app entry is removed, need stop goroutine to save memory, use context
		entry := app.getEntry(entryId)
		if entry.AppName == "" {
			return
		}
		// try to remove entry
		if !app.removeEntry(entryId) {
			return
		}
		// post close data
		app.post(que, entry, false)
	})
	// get all entries
	entries, err := app.ddeDock.Entries().Get(0)
	if err != nil {
		return
	}
	// monitor all entry
	for _, entry := range entries {
		go app.monitor(que, entry)
	}
	return
}

// GetCollectName get app collector
func (app *appCollectorItem) GetCollectName() string {
	return "app"
}

// Handler handle write to database result
// at this time, just drop data if failed
func (app *appCollectorItem) Handler(base abstract.BaseQueue, controller abstract.BaseController, result define.WriteResult) {
	return
}

// GetInterface write database table
func (app *appCollectorItem) GetInterface() string {
	return "v2/report/unification"
}

// monitor use to monitor app entry state
func (app *appCollectorItem) monitor(que abstract.BaseQueue, entryPath dbus.ObjectPath) {
	// create entry item
	entryObj, err := dock.NewEntry(app.sesBus, entryPath)
	// it is ok, of just cant get such entry
	if err != nil {
		logger.Warningf("create new entry failed, err: %v", err)
		return
	}
	// get desktop file from entry obj,
	// desktop file wont be changed, so dont care about the file
	desktop, err := entryObj.DesktopFile().Get(0)
	if err != nil {
		logger.Warningf("get desktop file failed, err: %v", err)
		return
	}
	// get app name from entry obj
	name, err := entryObj.Name().Get(0)
	if err != nil {
		logger.Warningf("get app name failed, err: %v", err)
	}
	// run dpkg to get package name
	dpkg := []string{"dpkg", "-S", desktop}
	cmd := exec.Command("/bin/bash", "-c", strings.Join(dpkg, " "))
	buf, err := cmd.CombinedOutput()
	if err != nil {
		logger.Warningf("get package name failed, err: %v", err)
	}
	pkg := strings.Trim(string(buf), "\n")
	// get id from dbus path,
	//such as /com/deepin/dde/daemon/Dock/entries/e0T61978f3b
	entryId := filepath.Base(string(entryPath))
	// TODO: should get entry desktop here
	entry := define.AppEntry{
		PkgName: pkg,
		AppName: name,
	}
	// monitor entry active state change
	err = entryObj.IsActive().ConnectChanged(func(hasValue bool, value bool) {
		// check if value is valid
		if !hasValue {
			return
		}
		// TODO: this code can use state machine to optimize
		// if now entry is active, should add this to map
		var open bool
		if value {
			// check if need post open app data
			post := app.addEntry(entryId, entry)
			if !post {
				return
			}
			// now is open app
			open = true
		} else {
			// check if need post close app data
			post := app.removeEntry(entryId)
			if !post {
				return
			}
			// now is close app
			open = false
		}
		// post app open/close data
		app.post(que, entry, open)
	})
	if err != nil {
		return
	}
	// get now active state
	active, err := entryObj.IsActive().Get(0)
	if err != nil {
		return
	}
	// now app is not active yet, dont need store
	// if add entry success, need post data
	if !active && app.addEntry(entryId, entry) {
		return
	}
	// post open app data
	app.post(que, entry, true)
}

// post write data to database
func (app *appCollectorItem) post(que abstract.BaseQueue, desktop define.AppEntry, open bool) {
	// store time
	desktop.Time = time.Now().UnixNano()
	// marshal app info
	data, err := json.Marshal(&desktop)
	if err != nil {
		logger.Warningf("marshal app data failed, err: %v")
		return
	}
	// request message
	req := define.RequestMsg{
		Pri:  define.SimpleRequest,
		Rule: define.LooseRule,
		Msg:  []string{string(data)},
	}
	// push data to queue
	que.Push(define.DataBaseItemQueue, app, req)
}

// getEntry get entry message from map
func (app *appCollectorItem) getEntry(entryId string) define.AppEntry {
	// lock
	app.lock.Lock()
	defer app.lock.Unlock()
	return app.items[entryId]
}

// addEntry add entry message to map,
// if this entry has already exist, dont need to post app open message
func (app *appCollectorItem) addEntry(entryId string, entry define.AppEntry) bool {
	// lock
	app.lock.Lock()
	defer app.lock.Unlock()
	// entry has already app, and entry is the same,
	// so entry has already exist, dont need to post data
	if item, ok := app.items[entryId]; ok && item.Id == entry.Id {
		return false
	}
	// if entryId not exist, or entry is new obj, need post data
	app.items[entryId] = entry
	return true
}

// removeEntry remove entry message from map,
// if removed app is now not exist
func (app *appCollectorItem) removeEntry(entryId string) bool {
	// lock
	app.lock.Lock()
	defer app.lock.Unlock()
	// get entry if exist
	if _, ok := app.items[entryId]; ok {
		return false
	}
	// delete entry from items
	delete(app.items, entryId)
	return true
}
