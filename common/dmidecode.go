package common

import (
	"strings"

	"github.com/ArisAachen/experience/define"
)

// dmidecode module
// use dmidecode command to get system info message
// man doc: https://linux.die.net/man/8/dmidecode
// file dir: /sys/class/dmi/id

// cpuParser use to parse cpu info file
type cpuParser struct {
}

// parse use to parse file line, search for version and id
func (cp *cpuParser) parse(info *define.BaseInfo, buf []byte) {
	// check if info is valid
	if info == nil {
		return
	}
	msgSl := strings.Split(string(buf), ":")
	// check length, in case of panic
	if len(msgSl) < 2 {
		return
	}
	// get key
	key := strings.TrimSpace(msgSl[0])
	// check key to fix value
	switch key {
	// get processor version
	case define.ProcessorVersion.String():
		// "Version: Intel(R) Core(TM) i7-10700 CPU @ 2.90GHz"
		info.Model = strings.TrimSpace(msgSl[1])
	// get processor id
	case define.ProcessorId.String():
		// "ID: 55 06 0A 00 FF FB EB BF"
		// trim left and right
		id := strings.TrimLeft(msgSl[1], " ")
		id = strings.TrimRight(id, " ")
		// replace all space to -
		info.Id = strings.Replace(id, " ", "-", -1)
	default:
		return
	}
}

// tool cpu info use dmidecode tool
func (cp *cpuParser) tool() string {
	return define.DmiDecode.String()
}

// module get current module param as dmi param
func (cp *cpuParser) module() string {
	return define.CpuModule.Module()
}

// param indicate parser param to exec
func (cp *cpuParser) param() string {
	args := []string{cp.tool(), "-t", cp.module()}
	return strings.Join(args, " ")
}

// boardParser use to parse base board file
type boardParser struct {
}

// parse use to parse file line, search for product-name and serial-number
func (bc *boardParser) parse(info *define.BaseInfo, buf []byte) {
	// check if info is valid
	if info == nil {
		return
	}
	msgSl := strings.Split(string(buf), ":")
	// check length, in case of panic
	if len(msgSl) < 2 {
		return
	}
	// get key
	key := strings.TrimSpace(msgSl[0])
	// check key to fix value
	switch key {
	// Product Name: B460M-HDV(RD) ; Serial Number: M80-D5014001427
	case define.BoardProductName.String():
		info.Model = strings.TrimSpace(msgSl[1])
	// Serial Number: M80-D5014001427
	case define.BoardSerialNumber.String():
		info.Id = strings.TrimSpace(msgSl[1])
	default:
		return
	}
}

// tool board info use dmidecode tool
func (bc *boardParser) tool() string {
	return define.DmiDecode.String()
}

// module get current module param as dmi param
func (bc *boardParser) module() string {
	return define.BoardModule.Module()
}

// param indicate parser param to exec
func (bc *boardParser) param() string {
	args := []string{bc.tool(), "-t", bc.module()}
	return strings.Join(args, " ")
}

// memoryParser
type memoryParser struct {
}

// parse use to parse file line, search for product-name and serial-number
func (mem *memoryParser) parse(info *define.BaseInfo, buf []byte) {
	// check if info is valid
	if info == nil {
		return
	}
	msgSl := strings.Split(string(buf), ":")
	// check length, in case of panic
	if len(msgSl) < 2 {
		return
	}
	// get key
	key := strings.TrimSpace(msgSl[0])
	switch key {
	case define.MemoryMaximumCapacity.String():
		info.Model = strings.TrimSpace(msgSl[1])
	default:
		return
	}
}

// tool memory info use dmidecode tool
func (mem *memoryParser) tool() string {
	return define.DmiDecode.String()
}

// module get current module param as dmi param
func (mem *memoryParser) module() string {
	return define.MemoryModule.Module()
}

// param indicate parser param to exec
func (mem *memoryParser) param() string {
	args := []string{mem.tool(), "-t", mem.module()}
	return strings.Join(args, " ")
}
