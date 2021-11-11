package writer

import (
	"pkg.deepin.io/lib/log"
)

var logger *log.Logger

func init() {
	logger = log.NewLogger("writer/user-exp")
	logger.SetLogLevel(log.LevelDebug)
}
