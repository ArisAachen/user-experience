package abstract

type BaseConfig interface {
	Load()
	AddModule(module string, loader FileLoader)
}
