package common

import (
	"encoding/json"
	"strings"

	"github.com/ArisAachen/experience/define"
)

// lsblk module
// use lsblk command to get block info message
// man doc: https://linux.die.net/man/8/lsblk
// file dir: /sys/class/block

type blkMsg struct {
	Serial   string `json:"serial"`
	Type     string `json:"type"`
	Vendor   string `json:"vendor"`
	Model    string `json:"model"`
	Children string `json:"children"`
}

// blkParser use to parse block info
type blkParser struct {
}

func (blk *blkParser) parse(info *define.BaseInfo, buf []byte) {
	// check if info is valid
	if info == nil {
		return
	}
	msgSl := strings.Split(string(buf), ":")
	// check length, in case of panic
	if len(msgSl) < 2 {
		return
	}
	// for block module, buf type is json
	// try to unmarshal file
	var block blkMsg
	err := json.Unmarshal(buf, &block)
	if err != nil {
		return
	}
	info.Model = block.Model
	info.Id = block.Serial
}

// tool block info use lsblk tool
func (blk *blkParser) tool() string {
	return define.LsBlk.String()
}

func (blk *blkParser) module() string {
	return define.DiskModule.Module()
}

func (blk *blkParser) param() string {
	// "lsblk -J -no SERIAL,TYPE,SIZE,VENDOR,MODEL"
	args := []string{define.LsBlk.String(), "-J", "-no", "SERIAL,TYPE,SIZE,VENDOR,MODEL"}
	return strings.Join(args, " ")
}
