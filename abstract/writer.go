package abstract

import (
	"github.com/ArisAachen/experience/define"
)

// BaseWriter indicate which writer to use
// after write, call handler here
type BaseWriter interface {
	AddModule(name define.WriterItemModule, item BaseWriterItem)
	Write(name define.WriterItemModule, crypt BaseCryptor, controller BaseController, handler BaseQueueHandler, creator BaseUrlCreator, msg []string)
	Connect()
}

// BaseWriterItem the abstract writer, indicate the abstract methods
// all writer handler should realize
// path is the url of post web server or table name of database
type BaseWriterItem interface {
	GetWriterItemName() define.WriterItemModule
	Write(crypt BaseCryptor, path string, msg []string) define.WriteResult
	// Connect connect to web or database
	Connect(url string)
	// GetRemote get connect path
	GetRemote() string
}
