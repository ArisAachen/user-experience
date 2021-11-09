package config

import "sync"

var globalCfgMgr *CfgMgr
var once sync.Once

// CfgMgr Config Manager to manager all config
type CfgMgr struct {
	sysCfg *SysCfg
}

// GetInstance create global config once
func GetInstance() *CfgMgr {
	once.Do(func() {
		globalCfgMgr = &CfgMgr{
			sysCfg: new(SysCfg),
		}
	})
	return globalCfgMgr
}

func (mgr *CfgMgr) SysCfg() *SysCfg {
	// make sure sysconfig exist
	if mgr.sysCfg == nil {
		mgr.sysCfg = &SysCfg{}
	}
	return mgr.sysCfg
}
