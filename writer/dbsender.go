package writer

import (
	"github.com/ArisAachen/experience/define"
)

type dbSender struct {
}

func newDBWriter() *dbSender {
	db := &dbSender{
	}
	return db
}

func (db *dbSender) Write(url string, msg string) define.WriteResult {
	var result define.WriteResult

	return result
}
