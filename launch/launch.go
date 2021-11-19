package launch

import (
	"github.com/ArisAachen/experience/abstract"
	"github.com/ArisAachen/experience/config"
	"github.com/ArisAachen/experience/crypt"
	"github.com/ArisAachen/experience/define"
	"github.com/ArisAachen/experience/queue"
	"github.com/ArisAachen/experience/writer"
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
}

func NewLaunch() *Launch {
	lch := &Launch{
	}
	return lch
}

// Init ref module
func (lau *Launch) Init() {
	// TODO
	lau.writer = writer.NewWriter(nil)
	lau.queue = queue.NewQueue(nil)
	lau.config = config.NewConfig(nil)
	lau.crypt = crypt.NewCryptor(nil)
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

// AddWriterItemModules add writer item to module
// now only has two module: web sender and database writer
func (lau *Launch) AddWriterItemModules() {
	if lau.writer == nil {
		logger.Warningf("cant add writer modules, write hasn't been init")
		return
	}
	// define need add modules
	modules := []define.WriterItemModule{
		define.WebItemWriter, define.DataBaseItemWriter,
	}
	// add modules
	for _, module := range modules {
		lau.config.AddModule(string(module))
	}
	logger.Debugf("writer modules add success, modules: %v", modules)
}

// AddQueueItemModules add queue item to queue
// now only has two module: web queue and database queue
func (lau *Launch) AddQueueItemModules() {
	if lau.queue == nil {
		logger.Warningf("cant add queue modules, queue hasn't been init")
		return
	}
	// define need add modules
	modules := []define.QueueItemModule{
		define.WebItemQueue, define.DataBaseItemQueue,
	}
	// add modules
	for _, module := range modules {
		lau.config.AddModule(string(module))
	}
	logger.Debugf("queue modules add success, modules: %v", modules)
}

// AddConfigItemModules add config item to config
// now only has three module: post system hardware
func (lau *Launch) AddConfigItemModules() {
	if lau.queue == nil {
		logger.Warningf("cant add queue modules, queue hasn't been init")
		return
	}
	// define need add modules
	modules := []define.ConfigItemModule{
		define.PostItemConfig, define.SystemItemConfig, define.HardwareItemConfig,
	}
	// add module
	for _, module := range modules {
		lau.config.AddModule(string(module))
	}
	logger.Debugf("queue modules add success, modules: %v", modules)
}

// AddCryptItemModules add crypt item to crypt data
// since now crypt only has one item,
// add this method is mention to notice maintainer,
// this module can support more parts of crypt
func (lau *Launch) AddCryptItemModules() {

}

// StartService start service
func (lau *Launch) StartService() {
	// start writer first
	lau.launchWriter()
	// start queue
	lau.launchQueue()
	//	launch config
	lau.launchConfig()
}

// launchWriter to make sure data can be sent and write,
// writer module should be init at beginning
func (lau *Launch) launchWriter() {
	lau.writer.Connect()
}

// launchQueue should be start
// queue start second time after writer is started
func (lau *Launch) launchQueue() {
	// TODO can optimize here
	// start pop data to webserver once data is push into queue
	lau.queue.Pop(define.WebItemQueue, lau.controller, lau.crypt, lau.writer)
	// start pop data to database once data is push into queue
	lau.queue.Pop(define.DataBaseItemQueue, lau.controller, lau.crypt, lau.writer)
}

// refreshConfig decide if need to update config message
func (lau *Launch) launchConfig() {

}
