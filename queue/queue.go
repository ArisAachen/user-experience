package queue

import "github.com/ArisAachen/experience/writer"

// ContainerQueue contains database queue and webserver queue
// caller use name to decide which queue to push
type ContainerQueue struct {
	items map[string]BaseQueueItem
}

func (cn *ContainerQueue) Push(name string, base BaseQueueHandler, msg string) {
	// find item to push data
	item, ok := cn.items[name]
	if !ok {
		logger.Warningf("push data failed, queue %s not exist", name)
		return
	}
	// push data
	item.Push(base, msg)
}

// Pop begin pop data to
func (cn *ContainerQueue) Pop(name string, sender writer.BaseWriter) {
	// find item to pop data
	item, ok := cn.items[name]
	if !ok {
		logger.Warningf("push data failed, queue %s not exist", name)
		return
	}
	// push data
	item.Pop(sender)
}
