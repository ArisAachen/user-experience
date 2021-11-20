package writer

import (
	"database/sql"
	"errors"
	"fmt"
	"net/url"
	"sync"

	"github.com/ArisAachen/experience/abstract"
	"github.com/ArisAachen/experience/common"
	"github.com/ArisAachen/experience/define"
	"github.com/ArisAachen/experience/queue"
	_ "github.com/mattn/go-sqlite3"
)

// database writer use to write data to database
/*
	head:       type      data    nano

	@type: type of data, some data is important, these data should lazy deleted
	@data: encrypted data,
	@nano: data happens time, use to delete old data, if database size if out of range
*/

type DBSender struct {
	client *sql.DB

	lock sync.Mutex
}

func NewDBWriter() *DBSender {
	db := &DBSender{
	}
	return db
}

// Connect connect to database
func (db *DBSender) Connect(dbPath string) {
	handle, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		logger.Warningf("open database %v failed, err: %v", dbPath, err)
		return
	}
	db.client = handle
	return
}

// GetRemote get local database path
func (db *DBSender) GetRemote() string {
	return define.SqlitePath
}

// Disconnect disconnect from database
func (db *DBSender) Disconnect() error {
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
func (db *DBSender) Write(crypt abstract.BaseCryptor, table string, msg string) define.WriteResult {
	// write database result
	var result define.WriteResult
	value, err := url.ParseQuery(msg)
	if err != nil {
		result.ResultCode = define.WriteParseQueryFailed
		logger.Warningf("parse query message failed, err: %v", err)
		return result
	}
	// get tid from params
	id := value.Get(string(define.Tid))
	// get data happens time from params
	time := value.Get(string(define.DataTime))
	// check if database already opened
	if db.client == nil {
		result.ResultCode = define.WriteResultUnknown
		logger.Warning("database is not opened yet")
		return result
	}
	// create lock
	db.lock.Lock()
	defer db.lock.Unlock()
	// check if handle exist
	if db.client == nil {
		result.ResultCode = define.WriteResultUnknown
		logger.Warning("database is not opened yet")
		return result
	}
	// create table if table not exist
	err = db.createTable(table)
	if err != nil {
		result.ResultCode = define.WriteResultUnknown
		logger.Warningf("create table %v failed, err: %v", table, err)
		return result
	}
	// create insert request
	insert := fmt.Sprintf("insert into %v(Type Data Nano) values(%v %v %v)", table, id, msg, time)
	logger.Debugf("insert table %v command: %v", table, insert)
	_, err = db.client.Query(insert)
	if err != nil {
		result.ResultCode = define.WriteResultWriteFailed
		logger.Warningf("insert data into database failed, err: %v", err)
		return result
	}
	logger.Debug("insert data into database success")
	return result
}

// Push database push data into queue
func (db *DBSender) Push(que queue.Queue) {
	if db.client == nil {
		logger.Warning("database is not opened yet")
		return
	}
	// select data from table
	req := "select Type ,Data from table order by Type,Nano"
	row, err := db.client.Query(req)
	if err != nil {
		logger.Warningf("get data from database failed, err: %v", err)
		return
	}
	// create table header
	var id int
	var msg string
	// row to end
	for row.Next() {
		// scan result
		err = row.Scan(&id, &msg)
		if err != nil {
			logger.Warningf("scan failed, err: %v", err)
			continue
		}
		// convert data level
		level := common.QueryLevel(define.TidTyp(id))
		// TODO rule
		req := define.RequestMsg{
			Rule: define.StrictRule,
			Pri:  level,
			Msg:  msg,
		}
		// push current item to web queue
		que.Push(define.WebItemQueue, db, req)
	}
}

func (db *DBSender) GetInterface() string {
	return ""
}

// Handler handle sent data result
func (db *DBSender) Handler(base abstract.BaseQueue, controller abstract.BaseController, result define.WriteResult) {
	// check if database is opened
	if db.client == nil {
		logger.Warning("database is not opened yet")
		return
	}
	// if if write failed, dont delete data from db
	if result.ResultCode == define.WriteResultWriteFailed {
		return
	}
	// TODO table
	req := "delete from table where data=" + result.Origin
	_, err := db.client.Query(req)
	if err != nil {
		logger.Warningf("delete data from table %v failed, err: %v", err)
		return
	}
	// data is deleted successfully
	logger.Debugf("delete data from database success, data: %v", result.Origin)
}

// createTable create table if table not exist
func (db *DBSender) createTable(table string) error {
	// check if database is init
	if db.client == nil {
		return errors.New("database is not opened yet")
	}
	// create table is table not exist
	typ := `
	Type integer NOT NULL,
	Data Text NOT NULL,
	Nano timestamp NOT NULL
	`
	// create table typ
	req := fmt.Sprintf("create table if not exists %v(%v)", table, typ)
	logger.Debugf("sql create table req: %v", req)
	_, err := db.client.Query(req)
	if err != nil {
		return err
	}
	logger.Debugf("table %v exist", table)
	return nil
}
