package collect

import "pkg.deepin.io/lib/log"

var logger *log.Logger

func init() {
	logger = log.NewLogger("user-exp/collect")
	logger.SetLogLevel(log.LevelDebug)
}
