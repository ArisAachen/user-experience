package launch

import (
	"pkg.deepin.io/lib/log"
)

var logger *log.Logger

// init setting
func init() {
	// init log level
	logger = log.NewLogger("system/user-exp")
	logger.SetLogLevel(log.LevelDebug)

	// init global config
}
