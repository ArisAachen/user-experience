package config

import (
	"github.com/ArisAachen/experience/launch"
)

// GlobalConfig global management config
// the global config offer method to operate all config
type GlobalConfig struct {
	// launch save all handler
	// so diff module can be called any module
	launch *launch.Launch

	// all config in this package,
	hardware baseCfgElem
	post     baseCfgElem
	system   baseCfgElem
}

// NewGlobalConfig create global config to manager all configs
func NewGlobalConfig(launch *launch.Launch) *GlobalConfig {
	// launcher must exist
	if launch == nil {
		return nil
	}
	// create obj
	cfg := &GlobalConfig{
		launch:   launch,
		hardware: new(hardwareCfg),
		post:     new(postCfg),
		system:   new(sysCfg),
	}
	return cfg
}

// Load begin to load all config from file path
func (gl *GlobalConfig) Load() {
	// load dont care about failed, file may not exist
	// In the beginning, it is ok, just create file later
	_ = gl.hardware.LoadFromFile("hardware")
	_ = gl.post.LoadFromFile("post")
	_ = gl.system.LoadFromFile("system")
}

// Update check all module if need update
func (gl *GlobalConfig) Update() {
	// should update post interface at first start
	// these message has highest priority
	// also these request dont need to store into database if request failed
	if gl.post != nil && gl.post.needUpdate() {
		// push post update to queue
		gl.post.push(gl.launch.GetQueue())
	}

	//
	if gl.system != nil && gl.system.needUpdate() {

	}

	//
	if gl.hardware != nil && gl.hardware.needUpdate() {

	}
}


