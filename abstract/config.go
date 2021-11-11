package abstract

// BaseCfgItem the abstract config, indicate the abstract methods
// all config should realize
type BaseCfgItem interface {
	// GetConfigPath get config path
	GetConfigPath() string

	// NeedUpdate indicate if need to post request to web server
	NeedUpdate() bool

	// SaveToFile and LoadFromFile save and load config from file
	SaveToFile(filename string) error
	LoadFromFile(filename string) error
}

type BaseConfig interface {
	Load()
	Update()
	Module
}
