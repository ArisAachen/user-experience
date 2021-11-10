package config

import "sync"

var globalCfgMgr *CfgMgr
var once sync.Once

// CfgMgr Config Manager to manager all config
type CfgMgr struct {
	cfg map[string]BaseCfg
}

// GetInstance create global config once
func GetInstance() *CfgMgr {
	// init global config once
	once.Do(func() {
		globalCfgMgr = &CfgMgr{
			// create config mapS
			cfg: make(map[string]BaseCfg),
		}
	})
	return globalCfgMgr
}

// Init read all config file from config
func (mgr *CfgMgr) Init() {
	// check if map is nil
	if mgr.cfg == nil {
		return
	}
	// add post config module
	psc := new(PostCfg)
	_ = psc.LoadFromFile("post")
	mgr.cfg[psc.name()] = psc

	// add hardware config module
	hdc := new(HardwareCfg)
	_ = hdc.LoadFromFile("hardware")
	mgr.cfg[hdc.name()] = hdc

	// add system config module
	syc := new(SysCfg)
	_ = syc.LoadFromFile("system")
	mgr.cfg[syc.name()] = syc
}

// GetRandomDomain get random url to post each time
func (mgr *CfgMgr) GetRandomDomain() string {


	return ""
}
