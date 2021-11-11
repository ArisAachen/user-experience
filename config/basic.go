package config

import "github.com/ArisAachen/experience/queue"

// baseCfgElem the abstract config, indicate the interface methods
// all config should realize
type baseCfgElem interface {
	// needUpdate
	needUpdate() bool
	//
	push(queue queue.BaseQueue)
	// name
	name() string

	// SaveToFile and LoadFromFile save and load config from file
	SaveToFile(filename string) error
	LoadFromFile(filename string) error
}

type BaseConfig interface {
}
