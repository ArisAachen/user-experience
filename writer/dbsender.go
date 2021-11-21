package writer

import (
	"database/sql"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"
	"sync"

	"github.com/ArisAachen/experience/abstract"
	"github.com/ArisAachen/experience/define"
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

func (db *DBSender) Init() error {
	return nil
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

func (db *DBSender) GetCollectName() string {
	return ""
}

// Write write data to ref table
func (db *DBSender) Write(crypt abstract.BaseCryptor, table string, msg []string) define.WriteResult {
	// write database result
	var result define.WriteResult
	// get tid from params
	// id := 1
	// get data happens time from params
	// now := time.Now().UnixNano() / 1e6
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
	err := db.createTable(table)
	if err != nil {
		result.ResultCode = define.WriteResultUnknown
		logger.Warningf("create table %v failed, err: %v", table, err)
		return result
	}
	// create insert request
	input := strings.Join(msg, "")
	input = base64.StdEncoding.EncodeToString([]byte(input))
	insert := fmt.Sprintf(`insert into %v(Data) values("%v");`, table, input)
	logger.Debugf("insert table %v command: %v", table, insert)
	_, err = db.client.Exec(insert)
	if err != nil {
		result.ResultCode = define.WriteResultWriteFailed
		logger.Warningf("insert data into database failed, err: %v", err)
		return result
	}
	logger.Debug("insert data into database success")
	return result
}

// Collect database push data into queue
func (db *DBSender) Collect(que abstract.BaseQueue) {
	if db.client == nil {
		logger.Warning("database is not opened yet")
		return
	}
	// select data from table
	req := fmt.Sprintf("select Data from %v", "exp")
	row, err := db.client.Query(req)
	if err != nil {
		logger.Warningf("get data from database failed, err: %v", err)
		return
	}
	// create table header
	var dataMsg []string
	// row to end
	for row.Next() {
		var tmp string
		// scan result
		err = row.Scan(&tmp)
		if err != nil {
			logger.Warningf("scan failed, err: %v", err)
			continue
		}
		msg, err := base64.StdEncoding.DecodeString(tmp)
		if err != nil {
			logger.Warningf("decode database base64 failed, err: %v", err)
			continue
		}
		// convert data level
		dataMsg = append(dataMsg, string(msg))
	}
	// TODO rule
	msg := define.RequestMsg{
		Rule: define.LooseRule,
		Pri:  define.SimpleRequest,
		Msg:  dataMsg,
	}
	// push current item to web queue
	que.Push(define.WebItemQueue, db, msg)
}

func (db *DBSender) GetInterface() string {
	return ""
}

// Handler handle sent data result
func (db *DBSender) Handler(base abstract.BaseQueue, crypt abstract.BaseCryptor, controller abstract.BaseController, result define.WriteResult) {
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
	_, err := db.client.Exec(req)
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
	Type integer,
	Data Text,
	Nano timestamp
	`
	// create table typ
	req := fmt.Sprintf("create table if not exists %v(%v)", table, typ)
	logger.Debugf("sql create table req: %v", req)
	_, err := db.client.Exec(req)
	if err != nil {
		return err
	}
	logger.Debugf("table %v exist", table)
	return nil
}

func (db *DBSender) GetWriterItemName() define.WriterItemModule {
	return define.DataBaseItemWriter
}
