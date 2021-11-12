package abstract

import "github.com/ArisAachen/experience/define"

// BaseWriter indicate which writer to use
// after write, call handler here
type BaseWriter interface {
	Write(name define.WriterItemModule, handler BaseQueueHandler, msg define.CryptResult)
	Module
}

// BaseWriterItem the abstract writer, indicate the abstract methods
// all writer handler should realize
// path is the url of post web server or table name of database
type BaseWriterItem interface {
	Write(path string, msg define.CryptResult) define.WriteResult
}
