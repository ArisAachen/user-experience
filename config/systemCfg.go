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

// SysCfg System config use to store system state info,
// check if config has been changed,
// if is, need re-check send message to writer
type SysCfg struct {
	// file lock
	lock sync.Mutex
	define.SysCfg
}

// SaveToFile save protobuf config to file
func (sys *SysCfg) SaveToFile(filename string) error {
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
func (sys *SysCfg) LoadFromFile(filename string) error {
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
func (sys *SysCfg) needUpdate() bool {
	var update bool

	return update
}

// name indicate hardware module nameS
func (sys *SysCfg) name() string {
	return "SystemConfig"
}

func (sys *SysCfg) Handler(base queue.BaseQueue, msg string) {
	base.Push(nil, msg)
}

func (sys *SysCfg) GetInterface() string {

	return ""
}
