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

// Update check all module if need update
func (cfg *Config) Update() {
	// TODO now load config order is const, but need more flexible realization

	// now post update request to refresh post interface
	item, ok := cfg.items[define.PostItemConfig]
	if ok {
		item.GetConfigPath()
	}

	// check if hardware message has changed, if changed, need to request update uni id
	item, ok = cfg.items[define.HardwareItemConfig]
	if ok {

	}

	// post current develop-enabled
	item, ok = cfg.items[define.SystemItemConfig]
	if ok {

	}

}
