package common

import "github.com/ArisAachen/experience/define"

// baseParser is the basic system file parse
type baseParser interface {
	parse(info *define.BaseInfo, buf []byte)
	//exe() string
	param() string
}

// parserFactory create diff obj to parse buf
type parserFactory struct {
}

// createParser create parser
func (com *parserFactory) createParser(module define.SysModule) baseParser {
	// check module name to create parser
	var parser baseParser
	switch module {
	case define.CpuModule:
		parser = &cpuParser{}
	case define.BoardModule:
		parser = &boardParser{}
	case define.MemoryModule:
		parser = &memoryParser{}
	default:
	}
	return parser
}