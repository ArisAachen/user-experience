package queue

import (
	"github.com/ArisAachen/experience/abstract"
	"github.com/ArisAachen/experience/define"
)

// Queue contains database queue and webserver queue
// caller use name to decide which queue to push
type Queue struct {
	items map[define.QueueItemModule]abstract.BaseQueueItem
}

// NewQueue create queue
func NewQueue() *Queue {
	que := &Queue{
		make(map[define.QueueItemModule]abstract.BaseQueueItem),
	}
	return que
}

func (que *Queue) Push(module define.QueueItemModule, base abstract.BaseQueueHandler, msg define.RequestMsg) {
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
func (que *Queue) Pop(module define.QueueItemModule, controller abstract.BaseController, crypt abstract.BaseCryptor, sender abstract.BaseWriter) {
	// find item to pop data
	item, ok := que.items[module]
	if !ok {
		logger.Warningf("pop data failed, queue %s not exist", module)
		return
	}
	// push data
	item.Pop(crypt, controller, nil, sender)
}

// AddModule module
func (que *Queue) AddModule(name define.QueueItemModule, item abstract.BaseQueueItem) {
	if _, ok := que.items[name]; ok {
		return
	}
	que.items[name] = item
}
