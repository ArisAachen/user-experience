package config

import (
	"bytes"
	"encoding/json"
	"github.com/ArisAachen/experience/crypt"
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
	define.PostInterface
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
	buf, err := proto.Marshal(&st.PostInterface)
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
	err = proto.Unmarshal(buf, &st.PostInterface)
	if err != nil {
		return err
	}
	return nil
}

// name indicate hardware module nameS
func (st *postCfg) name() string {
	return "SystemConfig"
}

// Handler post interface is a tmp request,
// so these request will not save to database even sent failed
func (st *postCfg) Handler(base abstract.BaseQueue, controller abstract.BaseController, result define.WriteResult) {
	// update interface is strict rule
	defer controller.Release(define.StrictRule)

	// actually when post update interface not success, dont store it in database
	if result.ResultCode != define.WriteResultSuccess {
		logger.Warningf("update interface failed, reason code: %v", result.ResultCode)
		return
	}

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

	// unmarshal update interface
	var ifcSl []define.RcvInterface
	err = json.Unmarshal([]byte(data), &ifcSl)
	if err != nil {
		logger.Warningf("unmarshal decrypted post interface failed, err: %v", err)
		return
	}

	// save update url to config
	var domains []*define.PostDomain
	for _, ifc := range ifcSl {
		domain := &define.PostDomain{
			Time:    uint64(ifc.Update),
			UrlPath: ifc.Ip,
		}
		domains = append(domains, domain)
	}
	st.Domains = domains

	// save file to config
	err = st.SaveToFile(st.GetConfigPath())
	if err != nil {
		logger.Warningf("save post interface config failed, err: %v", err)
		return
	}
}

func (st *postCfg) GetInterface() string {

	return ""
}

// Push for update interface, should push data to webserver,
func (st *postCfg) Push(que abstract.BaseQueue) {

}

// GetConfigPath post interface config path
func (st *postCfg) GetConfigPath() string {
	return define.PostInterfacePath
}

// NeedUpdate post config should be call in every boot
func (st *postCfg) NeedUpdate() bool {
	return true
}
