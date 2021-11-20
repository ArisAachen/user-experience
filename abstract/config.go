package abstract

type BaseConfig interface {
	Load()
	Update(que BaseQueue)
	Module
}
