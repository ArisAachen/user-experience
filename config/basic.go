package config

// baseCfg the abstract config, indicate the interface methods
// all config should realize
type baseCfg interface {
	needUpdate() bool
	SaveToFile(filename string) error
	LoadFromFile(filename string) error
}
