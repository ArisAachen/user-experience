package config

import (
	"bytes"
	"encoding/json"
	"github.com/ArisAachen/experience/abstract"
	"github.com/ArisAachen/experience/common"
	"github.com/ArisAachen/experience/define"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
	"os"
	"sync"
)

// SysModule System config use to store system state info,
// check if config has been changed,
// if is, need re-check send message to writer
type SysModule struct {
	// file lock
	lock sync.Mutex
	define.SysCfg
}

// NewSysModule create system module
func NewSysModule() *SysModule {
	sys := &SysModule{
	}
	return sys
}

// Init init current module
func (sys *SysModule) Init() error {
	return nil
}

// Collect collect message
func (sys *SysModule) Collect(que abstract.BaseQueue) {
	// check if need update package
	if sys.updatePkg() {
		// get machine id from file
		machine, err := common.GetMachineId()
		if err != nil {
			logger.Warningf("get machine id failed, err: %v", err)
			return
		}
		// get edition from file
		edition, err := common.GetEdition()
		if err != nil {
			logger.Warningf("get edition failed, err: %v", err)
		}
		version := "1.1.0"
		// create update request
		body := define.WriteUpdateReq{
			Tid:     define.NewCheckUpdateTid,
			Type:    edition,
			Machine: machine,
			Version: version,
		}
		// marshal req to json
		buf, err := json.Marshal(body)
		if err != nil {
			logger.Warningf("marshal update request failed, err: %v", err)
			return
		}
		// create request
		req := define.RequestMsg{
			Rule: define.LooseRule,
			Pri:  define.SimpleRequest,
			Msg:  []string{string(buf)},
		}
		// push data to queue
		que.Push(define.WebItemQueue, sys, req)
	}
}

// GetCollectName collect name
func (sys *SysModule) GetCollectName() string {
	return "system"
}

// SaveToFile save protobuf config to file
func (sys *SysModule) SaveToFile(filename string) error {
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
func (sys *SysModule) LoadFromFile(filename string) error {
	// lock op
	sys.lock.Lock()
	defer sys.lock.Unlock()

	// read file and unmarshal
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	// unmarshal system config file
	err = proto.Unmarshal(buf, &sys.SysCfg)
	if err != nil {
		return err
	}
	return nil
}

// Handler handle web sender result
func (sys *SysModule) Handler(base abstract.BaseQueue, crypt abstract.BaseCryptor, controller abstract.BaseController, result define.WriteResult) {
	defer controller.Release(define.LooseRule)
	// decode msg to find tid
	var origin define.WriteOrigin
	// unmarshal origin data, to decide which request has been sent
	origin.Tid = define.NewCheckUpdateTid

	// for system config, write data to web sender failed,
	// should write data to database
	// should marshal encrypt msg here
	var cryptData define.CryptResult
	err := json.Unmarshal(result.Msg, &cryptData)
	if err != nil {
		logger.Warningf("unmarshal encrypted post interface failed, err: %v", err)
		return
	}
	// decrypt data
	data, err := crypt.Decode(cryptData)
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

// GetInterface get post interface
func (sys *SysModule) GetInterface() string {
	return ""
}

// GetConfigPath get config file
func (sys *SysModule) GetConfigPath() string {
	return define.SysCfgFile
}

// updatePkg now update package should always
func (sys *SysModule) updatePkg() bool {
	return true
}
