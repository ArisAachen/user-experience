package writer

import (
	"github.com/ArisAachen/experience/define"
	"github.com/ArisAachen/experience/queue"
)

// BaseWriter indicate which writer to use
// after write, call handler here
type BaseWriter interface {
	Write(name define.WriterItemModule, handler queue.BaseQueueHandler, msg string)
	AddItem(module define.WriterItemModule)
}

// BaseWriterItem the abstract writer, indicate the interface methods
// all writer handler should realize
// path is the url of post web server or table name of database
type BaseWriterItem interface {
	Write(path string, msg string) define.WriteResult
}
