package abstract

type BaseCollector interface {
	AddModule(name string, item BaseCollectorItem)
}

type BaseCollectorItem interface {
	Init() error
	Collect(que BaseQueue)
	GetCollectName() string
}
