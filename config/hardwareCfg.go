package config

import (
	"bytes"
	"io/ioutil"
	"os"
	"sync"

	"github.com/ArisAachen/experience/define"
	"github.com/ArisAachen/experience/queue"
	"github.com/golang/protobuf/proto"
)

// HardwareCfg System config use to store hardware info
// check if hardware is diff with last store,
// if is, need re-check uni id from web
type HardwareCfg struct {
	// file lock
	lock sync.Mutex
	define.HardwareInfo
}

// SaveToFile save protobuf config to file
func (hc *HardwareCfg) SaveToFile(filename string) error {
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
func (hc *HardwareCfg) LoadFromFile(filename string) error {
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

// check if hardware is changed, if is, should update
func (hc *HardwareCfg) needUpdate() bool {
	var update bool

	return update
}

// name indicate hardware module nameS
func (hc *HardwareCfg) name() string {
	return "HardwareConfig"
}

func (hc *HardwareCfg) Handler(base queue.BaseQueue, msg string) {
	base.Push(nil, msg)
}

func (hc *HardwareCfg) GetInterface() string {

	return ""
}
