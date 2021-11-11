package writer

import (
	"github.com/ArisAachen/experience/define"
	"github.com/ArisAachen/experience/launch"
	"github.com/ArisAachen/experience/queue"
)

// Writer include web sender and db writer
type Writer struct {
	// all writer is store here
	items map[define.WriterItemModule]BaseWriterItem

	// store launch here
	launch launch.Launch
}

// Write write data to diff writer item according to name
// TODO: name can be cover by self defined type
func (wr *Writer) Write(name define.WriterItemModule, handler queue.BaseQueueHandler, msg string) {
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
	for {
		// when circle arrive 3, should go to failed
		if circle == 3 {
			break
		}
		// write data
		logger.Debugf("begin to write data, circle: %v", circle)
		result = item.Write(handler.GetInterface(), msg)
		// TODO: here can be written by state module
		// only sent failed can active retry write
		if result.ResultCode != define.WriteResultWriteFailed {
			circle++
			break
		}
	}
	// handler write result, now only write web server failed case should be handle
	handler.Handler(wr.launch.GetQueue(), result)
}

// AddItem add ref item to item map
func (wr *Writer) AddItem(module define.WriterItemModule) {
	// check if module exist, add ref module
	switch module {
	case define.WebItemWriter:
		wr.items[define.WebItemWriter] = newWebWriter()
	case define.DataBaseItemWriter:

	default:
		logger.Warningf("add unknown writer item, module: %v", module)
	}
}
