package writer

import (
	"github.com/ArisAachen/experience/abstract"
	"github.com/ArisAachen/experience/define"
)

// Writer include web sender and db writer
type Writer struct {
	// all writer is store here
	items map[define.WriterItemModule]abstract.BaseWriterItem
}

// NewWriter create writer to manager all web sender and database writer
func NewWriter() *Writer {
	wr := &Writer{
		items: make(map[define.WriterItemModule]abstract.BaseWriterItem),
	}
	return wr
}

// Write write data to diff writer item according to name
func (wr *Writer) Write(name define.WriterItemModule, crypt abstract.BaseCryptor, controller abstract.BaseController,
	handler abstract.BaseQueueHandler, creator abstract.BaseUrlCreator, msg []string) {
	// find item to write
	item, ok := wr.items[name]
	if !ok {
		logger.Warningf("write data failed, writer %s not exist", name)
		return
	}
	// write data to writer
	// each write should retry 3 times, including web and database
	var circle int
	var result define.WriteResult
	// TODO should optimize
	var urlPath string
	if name == define.DataBaseItemWriter {
		urlPath = "exp"
	} else if name == define.DataBaseItemWriter {
		urlPath = creator.GetRandomPostUrls()[0] + creator.GetInterface(define.GeneralTid) + "?aid=uospc"
	}

	// write data to writer 3 times
	// TODO these code can be optimize, using "for and circle" seems no good design because lack of flexibility
	for {
		// when circle arrive 3, should go to failed
		if circle == 3 {
			break
		}
		// write data
		logger.Debugf("begin to write data, circle: %v", circle)
		result = item.Write(crypt, urlPath, msg)
		// only sent failed can active retry write
		if result.ResultCode != define.WriteResultWriteFailed {
			break
		}
		circle++
	}
	// handler write result, now only write web server failed case should be
	// TODO
	if name == define.WebItemWriter {
		go handler.Handler(nil, crypt, controller, result)
	}
}

func (wr *Writer) AddModule(name define.WriterItemModule, item abstract.BaseWriterItem) {
	// check if already exist
	if _, ok := wr.items[name]; ok {
		return
	}
	// save file loader
	wr.items[name] = item
}

// Connect each module try to connect self target
func (wr *Writer) Connect() {
	for _, item := range wr.items {
		item.Connect(item.GetRemote())
	}
}
