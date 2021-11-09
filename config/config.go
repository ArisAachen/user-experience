package config

import (
	"bytes"
	"io/ioutil"
	"os"
	"sync"

	"github.com/ArisAachen/experience/define"
	"github.com/golang/protobuf/proto"
)

// SysCfg System config use to store system info
// check if hardware is diff with last store,
// if is, need re-check uni id from web
type SysCfg struct {
	// file lock
	lock sync.Mutex
	define.HardWareInfo
}

// SaveToFile save protobuf config to file
func (cfg *SysCfg) SaveToFile(filename string) error {
	// lock op
	cfg.lock.Lock()
	defer cfg.lock.Unlock()

	// check and create file
	fObj, err := os.Create(filename)
	if err != nil {
		return err
	}

	// marshal config to buf
	buf, err := proto.Marshal(&cfg.HardWareInfo)
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
func (cfg *SysCfg) LoadFromFile(filename string) error {
	// lock op
	cfg.lock.Lock()
	defer cfg.lock.Unlock()

	// read file and unmarshal
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	err = proto.Unmarshal(buf, &cfg.HardWareInfo)
	if err != nil {
		return err
	}
	return nil
}
