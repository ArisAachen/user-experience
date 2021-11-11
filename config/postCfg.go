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

// postCfg System config use to store system state info,
// check if config has been changed,
// if is, need re-check send message to writer
type postCfg struct {
	// file lock
	lock sync.Mutex
	define.SysCfg
}

// SaveToFile save protobuf config to file
func (st *postCfg) SaveToFile(filename string) error {
	// lock op
	st.lock.Lock()
	defer st.lock.Unlock()

	// check and create file
	fObj, err := os.Create(filename)
	if err != nil {
		return err
	}

	// marshal config to buf
	buf, err := proto.Marshal(&st.SysCfg)
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
func (st *postCfg) LoadFromFile(filename string) error {
	// lock op
	st.lock.Lock()
	defer st.lock.Unlock()

	// read file and unmarshal
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	err = proto.Unmarshal(buf, &st.SysCfg)
	if err != nil {
		return err
	}
	return nil
}

// check if hardware is changed, if is, should update
func (st *postCfg) needUpdate() bool {
	var update bool

	return update
}

// name indicate hardware module nameS
func (st *postCfg) name() string {
	return "SystemConfig"
}

func (st *postCfg) Handler(base queue.BaseQueue, msg string) {
	// base.Push(nil, msg)
}

func (st *postCfg) GetInterface() string {

	return ""
}

func (st *postCfg) push(queue queue.BaseQueue) {

}