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

// Handler post interface is a tmp request,
// so these request will not save to database even sent failed
func (st *postCfg) Handler(base abstract.BaseQueue, result define.WriteResult) {
	if result.ResultCode != define.WriteResultSuccess {
		logger.Warningf("update interface failed, reason code: %v", result.ResultCode)
		return
	}



	// base.Push(define.DataBaseItemQueue, st, "")
}

func (st *postCfg) GetInterface() string {

	return ""
}

// Push for update interface, should push data to webserver,
func (st *postCfg) Push(que abstract.BaseQueue) {



}

func (st *postCfg) GetConfigPath() string {

	return ""
}

// NeedUpdate post config should be call in every boot
func (st *postCfg) NeedUpdate() bool {
	return true
}
