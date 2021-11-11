package abstract

import "github.com/ArisAachen/experience/abstract"

// BaseCfgItem the abstract config, indicate the abstract methods
// all config should realize
type BaseCfgItem interface {
	// GetConfigPath get config path
	GetConfigPath() string

	// NeedUpdate indicate if need to post request to web server
	NeedUpdate() bool

	// Push push data to ref writer
	Push(que abstract.BaseQueue)

	// SaveToFile and LoadFromFile save and load config from file
	SaveToFile(filename string) error
	LoadFromFile(filename string) error
}

type BaseConfig interface {
	Load()
	Update(que abstract.BaseQueue)
	Module
}
