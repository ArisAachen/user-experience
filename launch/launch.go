package launch

import (
	"github.com/ArisAachen/experience/abstract"
	"github.com/ArisAachen/experience/collect"
	"github.com/ArisAachen/experience/config"
	"github.com/ArisAachen/experience/control"
	"github.com/ArisAachen/experience/crypt"
	"github.com/ArisAachen/experience/define"
	"github.com/ArisAachen/experience/queue"
	"github.com/ArisAachen/experience/writer"
	"pkg.deepin.io/lib/dbusutil"
	"reflect"
)

/*
	launch
	This module is use for start project.
	It is the main part of this project.

	1. export dbus message
*/

type Launch struct {
	// save module handler
	collector  abstract.BaseCollector
	controller abstract.BaseController
	config     abstract.BaseConfig
	writer     abstract.BaseWriter
	queue      abstract.BaseQueue
	crypt      abstract.BaseCryptor
	creator    abstract.BaseUrlCreator

	// observer
	observer *control.Observer

	// system service
	sysService *dbusutil.Service
}

func NewLaunch() *Launch {
	lch := &Launch{
	}
	return lch
}

// Init ref module
func (lau *Launch) Init(sys *dbusutil.Service) {
	lau.sysService = sys
	// TODO
	lau.collector = collect.NewCollector()
	lau.controller = control.NewController()
	lau.writer = writer.NewWriter()
	lau.queue = queue.NewQueue()
	lau.config = config.NewConfig()
	lau.observer = control.NewObserver()
}

// ModuleDisPatch dispatch module to diff manager
// some modules may be included into more than one manager
func (lau *Launch) ModuleDisPatch() {
	// create crypt
	cry := crypt.NewCryptor()
	lau.crypt = cry
	// create writer
	dbw := writer.NewDBWriter()
	wbw := writer.NewWebWriter()
	// add writer
	wrSl := []abstract.BaseWriterItem{dbw, wbw}
	lau.AddWriter(wrSl)

	// create creator
	post := config.NewPostModule()
	lau.creator = config.NewPostModule()

	// create queue item
	dbq := queue.NewDbQueue()
	wbq := queue.NewWebQueue()
	// add queue item
	queSl := []abstract.BaseQueueItem{dbq, wbq}
	lau.AddQueue(queSl)

	// create all module, modules are used more than one
	hardware := config.NewHardwareModule()
	sys := config.NewSysModule()
	dbus := collect.NewDBusModule(lau.sysService)

	// create app monitor
	app := collect.NewAppCollector()

	// add module into config manager
	cfgItems := []abstract.FileLoader{cry, hardware, sys, post, dbus}
	lau.AddConfigFileLoader(cfgItems)

	// add module collector
	colItems := []abstract.BaseCollectorItem{hardware, sys, dbus, app}
	lau.AddCollector(colItems)

	// start observer
	typ := reflect.ValueOf(dbq)
	method := typ.MethodByName("Collect")
	// collect data from db
	Args := reflect.ValueOf(lau.queue)
	caller := define.Caller{
		Method: method,
		Args:   []reflect.Value{Args},
	}
	lau.observer.Register(define.ObServerDatabase, caller)

	// create notify
	notify := control.NewNotify()
	_ = notify.Init()
	err := notify.Monitor()
	if err != nil {
		logger.Warningf("create notify failed, err: %v", err)
		return
	}
	notify.Handle(lau.observer)

}

// AddConfigFileLoader add file loader item to file
func (lau *Launch) AddConfigFileLoader(loaderSl []abstract.FileLoader) {
	// add loaders
	for _, loader := range loaderSl {
		// try to load file
		err := loader.LoadFromFile(loader.GetConfigPath())
		if err != nil {
			logger.Warningf("load file failed, err: %v", err)
		}
		// add module into file
		lau.config.AddModule(loader.GetConfigPath(), loader)
	}
}

// AddCollector add collector item to collector
func (lau *Launch) AddCollector(colSl []abstract.BaseCollectorItem) {
	// add base to collector
	for _, col := range colSl {
		// init collector item
		err := col.Init()
		if err != nil {
			logger.Warningf("init collector failed, err: %v", err)
		}
		// add queue
		go col.Collect(lau.queue)
		lau.collector.AddModule(col.GetCollectName(), col)
	}
}

// AddQueue add queue item
func (lau *Launch) AddQueue(queSl []abstract.BaseQueueItem) {
	for _, que := range queSl {
		go que.Pop(lau.crypt, lau.controller, lau.creator, lau.writer)
		lau.queue.AddModule(que.GetQueueName(), que)
	}
}

func (lau *Launch) AddWriter(wrSl []abstract.BaseWriterItem) {
	for _, wr := range wrSl {
		wr.Connect(wr.GetRemote())
		lau.writer.AddModule(wr.GetWriterItemName(), wr)
	}
}

// GetCollector get collector
func (lau *Launch) GetCollector() abstract.BaseCollector {
	// check if collector is init
	if lau.collector == nil {
		logger.Warning("cant get collector, collector is not init yet")
		return nil
	}
	return lau.collector
}

// GetController get controller
func (lau *Launch) GetController() abstract.BaseController {
	// check if controller is init
	if lau.controller == nil {
		logger.Warning("cant get controller, controller is not init yet")
		return nil
	}
	return lau.controller
}

// GetConfig get config
func (lau *Launch) GetConfig() abstract.BaseConfig {
	// check if config is init
	if lau.config == nil {
		logger.Warning("cant get controller, controller is not init yet")
		return nil
	}
	return lau.config
}

// GetWriter get writer
func (lau *Launch) GetWriter() abstract.BaseWriter {
	// check if writer is init
	if lau.writer == nil {
		logger.Warning("cant get controller, controller is not init yet")
		return nil
	}
	return lau.writer
}

// GetQueue get queue
func (lau *Launch) GetQueue() abstract.BaseQueue {
	// check if controller is init
	if lau.queue == nil {
		logger.Warning("cant get controller, controller is not init yet")
		return nil
	}
	return lau.queue
}

// StartService start service
func (lau *Launch) StartService() {
	lau.ModuleDisPatch()
}

func (lau *Launch) StopService() {

}
