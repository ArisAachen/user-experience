package queue

import (
	"sync"

	"github.com/ArisAachen/experience/abstract"
	"github.com/ArisAachen/experience/define"
	"github.com/ArisAachen/experience/launch"
)

// WebQueueItem collect messages from
type WebQueueItem struct {
	// queue store callback and message
	queue queue
	// cond is signal if current queue is empty
	cond sync.Cond

	// save launcher
	launcher *launch.Launch
}

// create web queue
func newWebQueue() *WebQueueItem {
	// create web queue
	wq := &WebQueueItem{
		queue: queue{},
	}
	return wq
}

// Push data to queue
func (web *WebQueueItem) Push(handler abstract.BaseQueueHandler, msg define.RequestMsg) {
	// push data to queue
	web.queue.push(handler, msg.Msg)

	// notify this queue is not empty
	web.cond.Signal()
}

// Pop pop data to writer
func (web *WebQueueItem) Pop(crypt abstract.BaseCryptor, writer abstract.BaseWriter) {
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
		// use Cryptor to crypt data
		result, err := crypt.Encode(elem.msg)
		if err != nil {
			// when data encrypt failed, just drop this data
			// also this module can return to handler, if some special handle is needed
			logger.Warningf("failed to crypt data, err: %v", err)
			continue
		}
		// write msg to writer
		writer.Write(define.WebItemWriter, elem.handler, result)
	}
}
