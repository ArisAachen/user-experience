package writer

import (
	"database/sql"
	"errors"
	"fmt"
	"net/url"
	"strconv"
	"sync"

	"github.com/ArisAachen/experience/abstract"
	"github.com/ArisAachen/experience/define"
)

// database writer use to write data to database
/*
	head:       type      data    nano

	@type: type of data, some data is important, these data should lazy deleted
	@data: encrypted data,
	@nano: data happens time, use to delete old data, if database size if out of range
*/

type dbSender struct {
	client *sql.DB

	lock sync.Mutex
}

func newDBWriter() *dbSender {
	db := &dbSender{
	}
	return db
}

// Connect connect to database
func (db *dbSender) Connect(dbPath string) error {
	handle, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return err
	}
	db.client = handle
	return nil
}

// Disconnect disconnect from database
func (db *dbSender) Disconnect() error {
	// check if client exist
	if db.client == nil {
		return errors.New("database is not init")
	}
	// close database
	err := db.client.Close()
	if err != nil {
		return err
	}
	return nil
}

// Write write data to ref table
func (db *dbSender) Write(crypt abstract.BaseCryptor, table string, msg string) define.WriteResult {
	// create lock
	db.lock.Lock()
	defer db.lock.Unlock()
	// write database result
	var result define.WriteResult
	value, err := url.ParseQuery(msg)
	if err != nil {
		result.ResultCode = define.WriteParseQueryFailed
		logger.Warningf("parse query message failed, err: %v", err)
		return result
	}
	// get tid from params
	str := value.Get(string(define.Tid))
	id, err := strconv.Atoi(str)
	if err != nil {
		result.ResultCode = define.WriteResultUnknown
		logger.Warningf("convert tid to int failed, err: %v", err)
		return result
	}
	// get data happens time from params
	str = value.Get(string(define.DataTime))
	time, err := strconv.Atoi(str)
	if err != nil {
		logger.Warningf("convert data time to int failed, err: %v", err)
	}
	// check if database already opened
	if db.client == nil {
		result.ResultCode = define.WriteResultUnknown
		logger.Warning("database is not opened yet")
		return result
	}
	//

	return result
}

// createTable create table if table not exist
func (db *dbSender) createTable(table string) error {
	// check if database is init
	if db.client == nil {
		return errors.New("database is not opened yet")
	}
	// try to find table
	req := fmt.Sprintf(".table %v;", table)
	result, err := db.client.ExecContext()
	if err != nil {
		return err
	}

}
