package config

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"os"
	"sync"

	"github.com/ArisAachen/experience/abstract"
	"github.com/ArisAachen/experience/crypt"
	"github.com/ArisAachen/experience/define"
	"github.com/ArisAachen/experience/queue"
	"github.com/golang/protobuf/proto"
)

// hardwareCfg System config use to store hardware info
// check if hardware is diff with last store,
// if is, need re-check uni id from web
type hardwareCfg struct {
	// file lock
	lock sync.Mutex
	tmp  define.HardwareInfo
	define.HardwareInfo

	tunnel chan int
}

// SaveToFile save protobuf config to file
func (hc *hardwareCfg) SaveToFile(filename string) error {
	// lock op
	hc.lock.Lock()
	defer hc.lock.Unlock()

	// check and create file
	fObj, err := os.Create(filename)
	if err != nil {
		return err
	}

	// marshal config to buf
	buf, err := proto.Marshal(&hc.HardwareInfo)
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
func (hc *hardwareCfg) LoadFromFile(filename string) error {
	// lock op
	hc.lock.Lock()
	defer hc.lock.Unlock()

	// read file and unmarshal
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	err = proto.Unmarshal(buf, &hc.HardwareInfo)
	if err != nil {
		return err
	}
	return nil
}

// name indicate hardware module nameS
func (hc *hardwareCfg) name() string {
	return "HardwareConfig"
}

// Handler use to handle write result
func (hc *hardwareCfg) Handler(base abstract.BaseQueue, controller abstract.BaseController, result define.WriteResult) {
	// hardware update uni id is strict rule
	defer controller.Release(define.StrictRule)
	// for hardware config, write data to web sender failed,
	// should write data to database
	// should marshal encrypt msg here
	var cryptData define.CryptResult
	err := json.Unmarshal(result.Msg, &cryptData)
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
	// get uni id from remote
	var uni string
	err = json.Unmarshal([]byte(data), &uni)
	if err != nil {
		logger.Warningf("unmarshal receive uni id failed, err: %v", err)
		return
	}
	// check if uni id is the same
	if hc.GetUniId() == uni {
		logger.Debugf("uni id is the same, dont need to update: %v", err)
		return
	}
	// save uni id
	hc.UniId = uni
	// save file to config file
	err = hc.SaveToFile(hc.GetConfigPath())
	if err != nil {
		logger.Warningf("save hardware config file failed, err: %v", err)
		return
	}
	logger.Debug("write hardware config file success")
}

// Init init current module
func (hc *hardwareCfg) Init() {
	// load file first, it is ok if file not exist first time
	err := hc.LoadFromFile(hc.GetConfigPath())
	if err != nil {
		logger.Debugf("get load file from config, err: %v", err)
	}
}

// Collect to collect message
func (hc *hardwareCfg) Collect(que queue.Queue) {
	// check if need update hardware uni
	if hc.updateHardware() {

	} else if hc.updateUni() {

	}
	logger.Debug("update uni id end")
}

// hardware use to check if need update hard ware
func (hc *hardwareCfg) updateHardware() bool {

	return false
}

// uni use to check if need update uni id
func (hc *hardwareCfg) updateUni() bool {
	// check if exist uni id
	if hc.GetUniId() == "" {
		logger.Debug("uni id not exist, need update package")
		return true
	}
	// if already exist, dont need to update
	logger.Debug("uni id is already exist")
	return false
}

// update request update uni id
func (hc *hardwareCfg) update(que queue.Queue) {
	// marshal data
	buf, err := proto.Marshal(hc)
	if err != nil {
		logger.Warningf("marshal proto to buf failed, err: %v", err)
		return
	}
	logger.Debugf("marshal proto success, message: %v", string(buf))
	// create request message
	req := define.RequestMsg{
		Rule: define.StrictRule,
		Pri:  define.UpdateUniRequest,
		Msg:  string(buf),
	}
	// push data to message queue
	que.Push(define.WebItemQueue, hc, req)
}

// GetInterface get post interface
func (hc *hardwareCfg) GetInterface() string {

	return ""
}

// GetConfigPath get hardware config file
func (hc *hardwareCfg) GetConfigPath() string {
	return define.HwCfgFile
}
