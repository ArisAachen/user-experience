package writer

import "github.com/ArisAachen/experience/queue"

// BaseWriter the abstract writer, indicate the interface methods
// all writer handler should realize
type BaseWriter interface {
	Write(handler queue.BaseQueueHandler, msg string)
}
