package config

import (
	"github.com/ArisAachen/experience/abstract"
	"github.com/ArisAachen/experience/define"
	"github.com/ArisAachen/experience/launch"
)

// Config global management config
// the global config offer method to operate all config
type Config struct {
	// launch save all handler
	// so diff module can be called any module
	launch *launch.Launch

	// all config in this package,
	items map[define.ConfigItemModule]abstract.BaseCfgItem
}

// NewConfig create global config to manager all configs
func NewConfig(launch *launch.Launch) *Config {
	// launcher must exist
	if launch == nil {
		return nil
	}
	// create obj
	cfg := &Config{
		launch: launch,
	}
	return cfg
}

// AddModule add config item into config
func (cfg *Config) AddModule(module string) {
	switch define.ConfigItemModule(module) {
	case define.HardwareItemConfig:
		cfg.items[define.HardwareItemConfig] = new(hardwareCfg)
	case define.PostItemConfig:
		cfg.items[define.PostItemConfig] = new(postCfg)
	case define.SystemItemConfig:
		cfg.items[define.SystemItemConfig] = new(sysCfg)
	default:
	}
}

// Load begin to load all config from file path
func (cfg *Config) Load() {
	// each part load config from file
	for _, item := range cfg.items {
		// when first init, config may not exist
		// so it is ok if load file failed
		_ = item.LoadFromFile(item.GetConfigPath())
	}
}

// Update update modules update, the order is required
// 1. update post interface
// 2. update hardware uni id
// 3. update develop-enabled login logout
func (cfg *Config) Update(que abstract.BaseQueue) {
	// TODO now load config order is const, but need more flexible realization
	// push data
	modules := []define.ConfigItemModule{define.PostItemConfig, define.HardwareItemConfig, define.SystemItemConfig}
	for _, module := range modules {
		item, ok := cfg.items[module]
		if ok && item.NeedUpdate() {
			item.Push(que)
		}
	}
	logger.Debug("config update request finished")
}
