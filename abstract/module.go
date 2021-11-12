package abstract

type Module interface {
	AddModule(name string)
}

type FileLoader interface {
	// SaveToFile and LoadFromFile save and load config from file
	SaveToFile(filename string) error
	LoadFromFile(filename string) error
}
