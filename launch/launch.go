package launch

import (
	"github.com/ArisAachen/experience/abstract"
	"github.com/ArisAachen/experience/config"
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
	collector abstract.BaseCollector
	config    abstract.BaseConfig
	writer    abstract.BaseWriter
	queue     abstract.BaseQueue
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
}

// AddWriterItemModules add writer item to module
// now only has two module: web sender and database writer
func (lau *Launch) AddWriterItemModules() {
	if lau.writer == nil {
		logger.Warningf("cant add writer modules, write hasn't been init")
		return
	}
	// add modules
	lau.writer.AddModule(string(define.WebItemWriter))
	lau.writer.AddModule(string(define.DataBaseItemWriter))
	logger.Debugf("writer modules add success, modules: %v", []define.WriterItemModule{define.WebItemWriter,
		define.DataBaseItemWriter})
}

// AddQueueItemModules add queue item to queue
// now only has two module: web queue and database queue
func (lau *Launch) AddQueueItemModules() {
	if lau.queue == nil {
		logger.Warningf("cant add queue modules, queue hasn't been init")
		return
	}
	// add modules
	lau.queue.AddModule(string(define.WebItemQueue))
	lau.queue.AddModule(string(define.DataBaseItemQueue))
	logger.Debugf("queue modules add success, modules: %v", []define.QueueItemModule{define.WebItemQueue,
		define.DataBaseItemQueue})
}

// AddConfigItemModules add config item to config
// now only has three module: post system hardware
func (lau *Launch) AddConfigItemModules() {
	if lau.queue == nil {
		logger.Warningf("cant add queue modules, queue hasn't been init")
		return
	}
	// add modules
	lau.config.AddModule(string(define.PostItemConfig))
	lau.config.AddModule(string(define.SystemItemConfig))
	lau.config.AddModule(string(define.HardwareItemConfig))
	logger.Debugf("queue modules add success, modules: %v", []define.ConfigItemModule{define.PostItemConfig,
		define.SystemItemConfig, define.HardwareItemConfig})
}

// StartService start service
func (lau *Launch) StartService() {

}

// refreshConfig decide if need to update config message
func (lau *Launch) refreshConfig() {
	// TODO this part can be more flexible, but since now, it is ok


}
