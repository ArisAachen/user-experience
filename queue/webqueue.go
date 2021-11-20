package queue

import (
	"sync"

	"github.com/ArisAachen/experience/abstract"
	"github.com/ArisAachen/experience/define"
)

// WebQueueItem collect messages from
type WebQueueItem struct {
	// queue store callback and message
	queue queue
	// cond is signal if current queue is empty
	cond sync.Cond
}

// NewWebQueue create web queue
func NewWebQueue() *WebQueueItem {
	// create web queue
	wq := &WebQueueItem{
		queue: queue{},
		cond: sync.Cond{
			L: new(sync.Mutex),
		},
	}
	return wq
}

// Push data to queue
func (web *WebQueueItem) Push(handler abstract.BaseQueueHandler, msg define.RequestMsg) {
	// push data to queue
	web.queue.push(handler, msg)

	// notify this queue is not empty
	web.cond.Signal()
}

// Pop pop data to writer
func (web *WebQueueItem) Pop(crypt abstract.BaseCryptor, controller abstract.BaseController, creator abstract.BaseUrlCreator, writer abstract.BaseWriter) {
	// check if writer is valid
	if writer == nil {
		logger.Warning("writer failed, writer is nil")
		return
	}

	for {
		// if queue if empty, wait until has at last one elem
		if web.queue.empty() {
			web.cond.L.Lock()
			web.cond.Wait()
			web.cond.L.Unlock()
		}
		// at these time queue is not empty, should call writer to send message
		elem := web.queue.pop()
		// if body is nil, queue is empty as well
		if elem == nil {
			continue
		}
		// monitor rule
		controller.Monitor(elem.msg.Rule)
		// set current rule
		controller.Invoke(elem.msg.Rule)
		// write msg to writer
		writer.Write(define.WebItemWriter, crypt, controller, elem.handler, creator, elem.msg.Msg)
	}
}

// GetQueueName get queue name
func (web *WebQueueItem) GetQueueName() define.QueueItemModule {
	return define.WebItemQueue
}
