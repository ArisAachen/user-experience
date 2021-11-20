package abstract

import (
	"github.com/ArisAachen/experience/define"
)

// BaseQueueHandler push message to queue and handle writer callback
// push data to message queue, attach handler to msg,
// after msg is sent by writer, call handler to handler result
// queue handler could be config, collector and database(not sure now)
type BaseQueueHandler interface {
	// Handler handle result
	Handler(base BaseQueue, controller BaseController, result define.WriteResult)
}

// BaseQueue use to push data to queue, pop queue to writer
// name
type BaseQueue interface {
	Push(module define.QueueItemModule, base BaseQueueHandler, msg define.RequestMsg)
	Pop(module define.QueueItemModule, controller BaseController, crypt BaseCryptor, sender BaseWriter)
	AddModule(name define.QueueItemModule, item BaseQueueItem)
}

// BaseQueueItem include may items, each item write data to diff place
type BaseQueueItem interface {
	Push(base BaseQueueHandler, msg define.RequestMsg)
	Pop(crypt BaseCryptor, controller BaseController, creator BaseUrlCreator, writer BaseWriter)
	GetQueueName() define.QueueItemModule
}
