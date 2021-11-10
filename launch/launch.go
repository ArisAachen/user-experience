package launch

import (
	"github.com/ArisAachen/experience/config"
	"pkg.deepin.io/lib/log"
)

/*
	launch
	This module is use for start project.
	It is the main part of this project.

	1. export dbus message
*/

var logger *log.Logger

type Launch struct {
	basic exportBasic
}

func NewLaunch() *Launch {
	lch := &Launch{
		basic: new(experience),
	}
	return lch
}

// init setting
func init() {
	// init log level
	logger = log.NewLogger("system/user-exp")
	logger.SetLogLevel(log.LevelDebug)

	// init global config
	_ = config.GetInstance()
}
