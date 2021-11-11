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

func (sys *sysCfg) Handler(base queue.BaseQueue, msg string) {
	// base.Push(nil, msg)
}

func (sys *sysCfg) GetInterface() string {

	return ""
}

func (sys *sysCfg) push(queue queue.BaseQueue) {

}
