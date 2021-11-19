package writer

import (
	"github.com/ArisAachen/experience/abstract"
	"github.com/ArisAachen/experience/define"
	"github.com/ArisAachen/experience/launch"
)

// Writer include web sender and db writer
type Writer struct {
	// all writer is store here
	items map[define.WriterItemModule]abstract.BaseWriterItem

	// store launch here
	launch *launch.Launch
}

// NewWriter create writer to manager all web sender and database writer
func NewWriter(launch *launch.Launch) *Writer {
	wr := &Writer{
		launch: launch,
		items:  make(map[define.WriterItemModule]abstract.BaseWriterItem),
	}
	return wr
}

// Write write data to diff writer item according to name
func (wr *Writer) Write(name define.WriterItemModule, crypt abstract.BaseCryptor, controller abstract.BaseController, handler abstract.BaseQueueHandler, msg string) {
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

	// write data to writer 3 times
	// TODO these code can be optimize, using "for and circle" seems no good design because lack of flexibility
	for {
		// when circle arrive 3, should go to failed
		if circle == 3 {
			break
		}
		// write data
		logger.Debugf("begin to write data, circle: %v", circle)
		result = item.Write(crypt, handler.GetInterface(), msg)
		// only sent failed can active retry write
		if result.ResultCode != define.WriteResultWriteFailed {
			circle++
			break
		}
	}
	// handler write result, now only write web server failed case should be
	// TODO
	go handler.Handler(nil, controller, result)
}

// Connect each module try to connect self target
func (wr *Writer) Connect() {
	for _, item := range wr.items {
		err := item.Connect(item.GetRemote())
		if err != nil {
			logger.Warningf("connect to remote failed")
		}
	}
}

func (wr *Writer) AddModule(module string) {
	// check if module exist, add ref module
	switch define.WriterItemModule(module) {
	case define.WebItemWriter:
		wr.items[define.WebItemWriter] = newWebWriter()
	case define.DataBaseItemWriter:
		wr.items[define.DataBaseItemWriter] = newDBWriter()
	default:
		logger.Warningf("add unknown writer item, module: %v", module)
	}
}
