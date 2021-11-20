package config

import (
	"github.com/ArisAachen/experience/abstract"
)

// Config global management config
// the global config offer method to operate all config
type Config struct {
	// all config in this package,
	items map[string]abstract.FileLoader
}

// NewConfig create global config to manager all configs
func NewConfig() *Config {
	// create obj
	cfg := &Config{
	}
	return cfg
}

// AddModule add config item into config
func (cfg *Config) AddModule(module string, loader abstract.FileLoader) {
	// check if already exist
	if _, ok := cfg.items[module]; ok {
		return
	}
	// save file loader
	cfg.items[module] = loader
}

// Load begin to load all config from file path
func (cfg *Config) Load() {
	// each part load config from file
	for _, item := range cfg.items {
		// load file
		err := item.LoadFromFile(item.GetConfigPath())
		if err != nil {
			logger.Debugf("read config file failed, err: %v", err)
		}
	}
}
