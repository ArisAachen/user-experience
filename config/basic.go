package config

// BaseCfg the abstract config, indicate the interface methods
// all config should realize
type BaseCfg interface {
	// needUpdate
	needUpdate() bool
	// name
	name() string

	// SaveToFile and LoadFromFile save and load config from file
	SaveToFile(filename string) error
	LoadFromFile(filename string) error
}
