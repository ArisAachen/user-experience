package config

import "pkg.deepin.io/lib/log"

var logger *log.Logger

func init() {
	logger = log.NewLogger("user-exp/config")
	logger.SetLogLevel(log.LevelDebug)
}
