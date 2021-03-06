package config

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"os"
	"sync"

	"github.com/ArisAachen/experience/abstract"
	"github.com/ArisAachen/experience/common"
	"github.com/ArisAachen/experience/define"
	"github.com/ArisAachen/experience/queue"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
)

// HardwareModule System config use to store hardware info
// check if hardware is diff with last store,
// if is, need re-check uni id from web
type HardwareModule struct {
	// file lock
	lock sync.Mutex
	define.HardwareInfo

	tunnel chan int
}

// NewHardwareModule create hardware config module
func NewHardwareModule() *HardwareModule {
	hw := &HardwareModule{}
	return hw
}

// SaveToFile save protobuf config to file
func (hc *HardwareModule) SaveToFile(filename string) error {
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
func (hc *HardwareModule) LoadFromFile(filename string) error {
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

// Handler use to handle write result
func (hc *HardwareModule) Handler(base abstract.BaseQueue, crypt abstract.BaseCryptor, controller abstract.BaseController, result define.WriteResult) {
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
	data, err := crypt.Decode(cryptData)
	if err != nil {
		logger.Warningf("decode post interface message failed, err: %v", err)
		return
	}
	// check length
	if len(data) < 17 {
		logger.Warning("receive uni length if invalid")
	}
	data = data[16:]
	// get uni id from remote
	var uni define.RcvUni
	err = json.Unmarshal([]byte(data), &uni)
	if err != nil {
		logger.Warningf("unmarshal receive uni id failed, err: %v", err)
		return
	}
	// check if uni id is the same
	if hc.GetUniId() == uni.Uni {
		logger.Debugf("uni id is the same, dont need to update: %v", err)
		return
	}
	// save uni id
	hc.UniId = uni.Uni
	// save file to config file
	err = hc.SaveToFile(hc.GetConfigPath())
	if err != nil {
		logger.Warningf("save hardware config file failed, err: %v", err)
		return
	}
	logger.Debug("write hardware config file success")
}

// Init init current module
func (hc *HardwareModule) Init() error {
	return nil
}

// Collect to collect message
func (hc *HardwareModule) Collect(que abstract.BaseQueue) {
	// check if need update hardware uni
	if !hc.updateHardware() && !hc.updateUni() {
		return
	}
	hc.Tid = int32(define.NewSystemInfoTid)
	// create request
	var req define.RequestMsg
	marshal := jsonpb.Marshaler{}
	buf, err := marshal.MarshalToString(hc)
	if err != nil {
		logger.Warningf("marshal hardware message failed, err: %v", err)
		return
	}
	// create request message
	req.Rule = define.StrictRule
	req.Pri = define.UpdateUniRequest
	req.Msg = []string{string(buf)}
	logger.Debugf("hardware marshal %v", req.Msg)
	// push data
	que.Push(define.WebItemQueue, hc, req)
	logger.Debug("update uni id end")
}

// GetCollectName collect name
func (hc *HardwareModule) GetCollectName() string {
	return "hardware"
}

// hardware use to check if need update hard ware
func (hc *HardwareModule) updateHardware() bool {
	var update bool
	// cpu module
	info, err := common.GetCpuInfo()
	if err != nil {
		logger.Warningf("cant get cpu info", err)
	}
	// TODO these code can be optimize
	// check if need update
	if hc.GetCpu().GetModel() != info.Model || hc.GetCpu().GetId() != info.Id {
		update = true
		hc.Cpu = new(define.Obj)
		hc.GetCpu().Id = info.Id
		hc.GetCpu().Model = info.Model
	}

	// board module
	info, err = common.GetBaseBoardInfo()
	if err != nil {
		logger.Warningf("cant get board info, err: %v", err)
	}
	// check if need update
	if hc.GetBoard().GetModel() != info.Model || hc.GetBoard().GetId() != info.Id {
		update = true
		hc.Board = new(define.Obj)
		hc.GetBoard().Id = info.Id
		hc.GetBoard().Model = info.Model
	}

	// gpu module
	info, err = common.GetGpuInfo()
	if err != nil {
		logger.Warningf("cant get gpu info, err: %v", err)
	}
	// check if need update
	if hc.GetGpu().GetModel() != info.Model || hc.GetGpu().GetId() != info.Id {
		update = true
		hc.Gpu = new(define.Obj)
		hc.GetGpu().Id = info.Id
		hc.GetGpu().Model = info.Model
	}

	// memory
	info, err = common.GetMemoryInfo()
	if err != nil {
		logger.Warningf("cant get memory info, err: %v", err)
	}
	// check if need update
	if hc.GetMemory().GetModel() != info.Model {
		update = true
		hc.Memory = new(define.Obj)
		hc.GetMemory().Model = info.Id
	}

	// disk
	info, err = common.GetDiskInfo()
	if err != nil {
		logger.Warningf("cant get disk info, err: %v", err)
	}
	// check if need update
	if hc.GetDisk().GetModel() != info.Model {
		update = true
		hc.Disk = new(define.Obj)
		hc.GetDisk().Model = info.Id
	}

	// network
	info, err = common.GetNetworkInfo()
	if err != nil {
		logger.Warningf("cant get network info, err: %v", err)
	}
	info, err = common.GetEtherInfo()
	if err != nil {
		logger.Warningf("cant get ethernet info, err: %v", err)
	}
	// check if need update
	if hc.GetNetwork().GetModel() != info.Model {
		update = true
		hc.Network = new(define.Obj)
		hc.GetNetwork().Model = info.Id
	}

	// get machine id
	machine, err := common.GetMachineId()
	if err != nil {
		logger.Warningf("read machine id failed, err: %v", err)
	}
	// get token
	token, err := common.GetAptToken()
	if err != nil {
		logger.Warningf("read apt token failed, err: %v", err)
	}

	// get machine
	if hc.GetOther().GetMachine() != machine || hc.GetOther().GetApt() != token {
		hc.Other = new(define.HardwareOther)
		hc.Other.Machine = machine
		hc.Other.Apt = token
	}

	if update {
		go hc.SaveToFile(hc.GetConfigPath())
	}

	return update
}

// uni use to check if need update uni id
func (hc *HardwareModule) updateUni() bool {
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
func (hc *HardwareModule) update(que queue.Queue) {
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
		Msg:  []string{string(buf)},
	}
	// push data to message queue
	que.Push(define.WebItemQueue, hc, req)
}

// GetConfigPath get hardware config file
func (hc *HardwareModule) GetConfigPath() string {
	return define.HwCfgFile
}
