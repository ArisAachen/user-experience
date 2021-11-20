package config

import (
	"bytes"
	"encoding/json"
	"github.com/ArisAachen/experience/common"
	"github.com/ArisAachen/experience/crypt"
	"io/ioutil"
	"os"
	"sync"

	"github.com/ArisAachen/experience/abstract"
	"github.com/ArisAachen/experience/define"
	"github.com/golang/protobuf/proto"
)

// sysCfg System config use to store system state info,
// check if config has been changed,
// if is, need re-check send message to writer
type sysCfg struct {
	// file lock
	lock sync.Mutex
	define.SysCfg
}

// SaveToFile save protobuf config to file
func (sys *sysCfg) SaveToFile(filename string) error {
	// lock op
	sys.lock.Lock()
	defer sys.lock.Unlock()
	// check and create file
	fObj, err := os.Create(filename)
	if err != nil {
		return err
	}
	// marshal config to buf
	buf, err := proto.Marshal(&sys.SysCfg)
	if err != nil {
		return err
	}
	// bytes writer
	byBuf := bytes.NewBuffer(buf)
	_, err = byBuf.WriteTo(fObj)
	if err != nil {
		return err
	}
	return nil
}

// LoadFromFile load protobuf config from file
func (sys *sysCfg) LoadFromFile(filename string) error {
	// lock op
	sys.lock.Lock()
	defer sys.lock.Unlock()

	// read file and unmarshal
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	err = proto.Unmarshal(buf, &sys.SysCfg)
	if err != nil {
		return err
	}
	return nil
}

// check if hardware is changed, if is, should update
func (sys *sysCfg) needUpdate() bool {
	var update bool

	return update
}

// name indicate hardware module nameS
func (sys *sysCfg) name() string {
	return "SystemConfig"
}

// Handler handle web sender result
func (sys *sysCfg) Handler(base abstract.BaseQueue, controller abstract.BaseController, result define.WriteResult) {
	defer controller.Release(define.LooseRule)
	// decode msg to find tid
	msg := result.Origin
	var origin define.WriteOrigin
	// unmarshal origin data, to decide which request has been sent
	err := json.Unmarshal([]byte(msg), &origin)
	if err != nil {
		logger.Warningf("unmarshal origin data failed, err: %v", err)
		return
	}
	logger.Debugf("system config req tid %v receive response", origin.Tid)
	// for system config, write data to web sender failed,
	// should write data to database
	// should marshal encrypt msg here
	var cryptData define.CryptResult
	err = json.Unmarshal(result.Msg, &cryptData)
	if err != nil {
		logger.Warningf("unmarshal encrypted post interface failed, err: %v", err)
		return
	}
	// decrypt data
	decry := crypt.NewCryptor(nil)
	data, err := decry.Decode(cryptData)
	if err != nil {
		logger.Warningf("decode post interface message failed, err: %v", err)
		return
	}
	// TODO this code can be optimize
	// create request
	var req define.RequestMsg
	// check if current type is update package,
	// dont store this message event send failed
	if origin.Tid == define.NewCheckUpdateTid {
		// check if post update and receive response success
		if result.ResultCode != define.WriteResultSuccess {
			logger.Warningf("request update package failed, reason code: %v", result.ResultCode)
			return
		}
		// unmarshal update state
		var update int
		err = json.Unmarshal([]byte(data), &update)
		if err != nil {
			logger.Warningf("unmarshal update state failed, err: %v", err)
			return
		}
		// check if need update
		if update != 1 {
			logger.Debug("receive remote dont need to update package")
			return
		}
		// request update package
		out, err := common.UpdatePackage(define.PkgName)
		if err != nil {
			logger.Warningf("request update package failed, err: %v", err)
			return
		}
		logger.Debugf("request update package success, out: %v", out)
		return
	}
	// when cant write data into database, dont need to handle again, just drop this message
	base.Push(define.DataBaseItemQueue, nil, req)
}

func (sys *sysCfg) GetInterface() string {

	return ""
}

func (sys *sysCfg) GetConfigPath() string {

	return ""
}

// NeedUpdate when system config changed, should call update
func (sys *sysCfg) NeedUpdate() bool {
	return true
}

// Push for update interface, should push data to webserver,
func (sys *sysCfg) Push(que abstract.BaseQueue) {

}
