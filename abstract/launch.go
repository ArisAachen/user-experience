package abstract

// BaseLaunch use to launch all ap
type BaseLaunch interface {
	GetCollector() BaseCollector
	GetController() BaseController
	GetConfig() BaseConfig
	GetWriter() BaseWriter
	GetQueue() BaseQueue

	StartService()
	StopService()
}
