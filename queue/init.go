package queue

import "pkg.deepin.io/lib/log"

var logger *log.Logger

func init() {
	logger = log.NewLogger("experience/queue")
	logger.SetLogLevel(log.LevelDebug)
}
