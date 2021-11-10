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

// PostCfg System config use to store system state info,
// check if config has been changed,
// if is, need re-check send message to writer
type PostCfg struct {
	// file lock
	lock sync.Mutex
	define.SysCfg
}

// SaveToFile save protobuf config to file
func (st *PostCfg) SaveToFile(filename string) error {
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
func (st *PostCfg) LoadFromFile(filename string) error {
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
func (st *PostCfg) needUpdate() bool {
	var update bool

	return update
}

// name indicate hardware module nameS
func (st *PostCfg) name() string {
	return "SystemConfig"
}

func (st *PostCfg) Handler(base queue.BaseQueue, msg string) {
	base.Push(nil, msg)
}

func (st *PostCfg) GetInterface() string {

	return ""
}