package queue

import (
	"github.com/ArisAachen/experience/abstract"
	"github.com/ArisAachen/experience/define"
	"github.com/ArisAachen/experience/launch"
)

// Queue contains database queue and webserver queue
// caller use name to decide which queue to push
type Queue struct {
	launch *launch.Launch
	items  map[define.QueueItemModule]abstract.BaseQueueItem
}

// NewQueue create queue
func NewQueue(launch *launch.Launch) *Queue {
	que := &Queue{
		launch: launch,
	}
	return que
}

func (que *Queue) Push(module define.QueueItemModule, base abstract.BaseQueueHandler, msg string) {
	// find item to push data
	item, ok := que.items[module]
	if !ok {
		logger.Warningf("push data failed, queue %s not exist", module)
		return
	}
	// push data
	item.Push(base, msg)
}

// Pop begin pop data to
func (que *Queue) Pop(module define.QueueItemModule, sender abstract.BaseWriter) {
	// find item to pop data
	item, ok := que.items[module]
	if !ok {
		logger.Warningf("pop data failed, queue %s not exist", module)
		return
	}
	// push data
	item.Pop(sender)
}

// AddModule add queue module item into queue
func (que *Queue) AddModule(module string) {
	// check if module exist, add ref module
	switch define.QueueItemModule(module) {
	case define.WebItemQueue:
		que.items[define.WebItemQueue] = newWebQueue()
	case define.DataBaseItemQueue:

	default:
		logger.Warningf("add unknown queue item, module: %v", module)
	}
}
