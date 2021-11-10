package queue

import (
	"sync"

	"github.com/ArisAachen/experience/writer"
)

// WebQueue collect messages from
type WebQueue struct {
	// queue store callback and message
	queue queue
	// cond is signal if current queue is empty
	cond sync.Cond
}

// create web queue
func newWebQueue() *WebQueue {
	// create web queue
	wq := &WebQueue{
		queue: queue{},
	}
	return wq
}

// Push data to queue
func (web *WebQueue) Push(handler BaseQueueHandler, msg string) {
	// push data to queue
	web.queue.push(handler, msg)

	// notify this queue is not empty
	web.cond.Signal()
}

// Pop pop data to writer
func (web *WebQueue) Pop(writer writer.BaseWriter) {
	// check if writer is valid
	if writer == nil {
		logger.Warning("writer failed, writer is nil")
		return
	}

	for {
		// if queue if empty, wait until has at last one elem
		if web.queue.empty() {
			web.cond.Wait()
		}
		// at these time queue is not empty, should call writer to send message
		elem := web.queue.pop()
		// if body is nil, queue is empty as well
		if elem == nil {
			continue
		}
		// write msg to writer
		writer.Write(nil, elem.msg)
	}
}
