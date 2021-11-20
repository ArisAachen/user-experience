package abstract

type BaseCollector interface {
}

type BaseCollectorItem interface {
	Init() error
	Collect(que BaseQueue)
}
