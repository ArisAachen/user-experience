package config

import (
	"bytes"
	"io/ioutil"
	"os"
	"sync"

	"github.com/ArisAachen/experience/abstract"
	"github.com/ArisAachen/experience/define"
	"github.com/golang/protobuf/proto"
)

// hardwareCfg System config use to store hardware info
// check if hardware is diff with last store,
// if is, need re-check uni id from web
type hardwareCfg struct {
	// file lock
	lock sync.Mutex
	define.HardwareInfo
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

// check if hardware is changed, if is, should update
func (hc *hardwareCfg) needUpdate() bool {
	var update bool

	return update
}

// name indicate hardware module nameS
func (hc *hardwareCfg) name() string {
	return "HardwareConfig"
}

// Handler use to handle write result
func (hc *hardwareCfg) Handler(base abstract.BaseQueue, result define.WriteResult) {
	// for hardware config, write data to web sender failed,
	// should write data to database
	base.Push("", nil, "")
}

func (hc *hardwareCfg) GetInterface() string {

	return ""
}

func (hc *hardwareCfg) push(queue abstract.BaseQueue) {

}

func (hc *hardwareCfg) GetConfigPath() string {

	return ""
}

// NeedUpdate only when hardware message changed, should call update
// TODO
func (hc *hardwareCfg) NeedUpdate() bool {
	return true
}


// Push for update interface, should push data to webserver,
func (hc *hardwareCfg) Push(que abstract.BaseQueue) {

}