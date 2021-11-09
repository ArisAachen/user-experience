package common

import (
	"strings"

	"github.com/ArisAachen/experience/define"
)

// baseCommand is the basic system file parse
type baseCommand interface {
	parse(info *define.BaseInfo, buf []byte)
}

// cpuCommand use to parse cpu info file
type cpuCommand struct {
}

// parse use to parse file line, search for version and id
func (cp *cpuCommand) parse(info *define.BaseInfo, buf []byte) {
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

// boardCommand use to parse base board file
type boardCommand struct {
}

// parse use to parse file line, search for product-name and serial-number
func (bc *boardCommand) parse(info *define.BaseInfo, buf []byte) {
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
	switch strings.TrimSpace(key) {
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


