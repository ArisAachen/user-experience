package launch

import (
	"github.com/ArisAachen/experience/collect"
	"github.com/ArisAachen/experience/config"
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
	basic exportBasic

	// save module handler
	collector collect.BaseCollector
	config    config.BaseConfig
	writer    writer.BaseWriter

	// queue offer data queue to write data to database or webserver
	queue queue.BaseQueue
}

func NewLaunch() *Launch {
	lch := &Launch{
		basic: new(experience),
	}
	return lch
}

// GetCollector collector use for collect data
func (lau *Launch) GetCollector() collect.BaseCollector {
	return lau.collector
}

// GetConfig config to manage all config
func (lau *Launch) GetConfig() config.BaseConfig {
	return lau.config
}

// GetWriter writer to write data to web or data base
func (lau *Launch) GetWriter() writer.BaseWriter {
	return lau.writer
}

// GetQueue que to store and pop data into database
func (lau *Launch) GetQueue() queue.BaseQueue {
	return lau.queue
}
