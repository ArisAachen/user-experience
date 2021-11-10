package queue

import "github.com/ArisAachen/experience/writer"

// BaseQueueHandler push message to queue and handle writer callback
// push data to message queue, attach handler to msg,
// after msg is sent by writer, call handler to handler result
// queue handler could be config, collector and database(not sure now)
type BaseQueueHandler interface {
	GetInterface() string
	Handler(base BaseQueue, msg string)
}

// BaseQueue use to push data to queue, pop queue to writer
type BaseQueue interface {
	Push(base BaseQueueHandler, msg string)
	Pop(sender writer.BaseWriter)
}
