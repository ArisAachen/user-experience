package queue

import (
	"sync"

	"github.com/ArisAachen/experience/abstract"
	"github.com/ArisAachen/experience/define"
)

// DbQueueItem use to write database
type DbQueueItem struct {
	// queue store saved data and send web failed data
	queue queue
	// cond is signal if current queue is empty
	cond sync.Cond
}

// NewDbQueue create database queue
func NewDbQueue() *DbQueueItem {
	// create web queue
	wq := &DbQueueItem{
		queue: queue{},
		cond: sync.Cond{
			L: new(sync.Mutex),
		},
	}
	return wq
}

// Push data to queue
func (db *DbQueueItem) Push(handler abstract.BaseQueueHandler, msg define.RequestMsg) {
	// push data to queue
	db.queue.push(handler, msg)

	// notify this queue is not empty
	db.cond.Signal()
}

// Pop pop data to writer
func (db *DbQueueItem) Pop(crypt abstract.BaseCryptor, controller abstract.BaseController, creator abstract.BaseUrlCreator, writer abstract.BaseWriter) {
	// check if writer is valid
	if writer == nil {
		logger.Warning("writer failed, writer is nil")
		return
	}

	for {
		// if queue if empty, wait until has at last one elem
		if db.queue.empty() {
			db.cond.L.Lock()
			db.cond.Wait()
			db.cond.L.Unlock()
		}
		// at these time queue is not empty, should call writer to send message
		elem := db.queue.pop()
		// if body is nil, queue is empty as well
		if elem == nil {
			continue
		}
		// write data to database writer
		writer.Write(define.DataBaseItemWriter, crypt, controller, nil, nil, elem.msg.Msg)
	}
}

// GetQueueName get queue name
func (db *DbQueueItem) GetQueueName() define.QueueItemModule {
	return define.DataBaseItemQueue
}
