package abstract

// BaseCfgItem the abstract config, indicate the abstract methods
// all config should realize
type BaseCfgItem interface {
	// GetConfigPath get config path
	GetConfigPath() string

	// NeedUpdate indicate if need to post request to web server
	NeedUpdate() bool

	// Push push data to ref writer
	Push(que BaseQueue)

	// FileLoader the load config file interface
	FileLoader
}

type BaseConfig interface {
	Load()
	Update(que BaseQueue)
	Module
}
