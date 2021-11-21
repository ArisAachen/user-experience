package common

import (
	"strings"

	"github.com/ArisAachen/experience/define"
)

type gpuParser struct {
}

// parse used to parse gpu command line
func (gpu *gpuParser) parse(info *define.BaseInfo, buf []byte) {
	if info == nil {
		return
	}
	msgSl := strings.Split(string(buf), ":")
	// check length, in case of panic
	if len(msgSl) < 3 {
		return
	}
	// get key VGA compatible controller
	key := msgSl[1]
	if strings.Contains(key, define.VgaController.String()) {
		info.Model = strings.TrimLeft(msgSl[2], " ")
	}
	return
}

// tool memory info use dmidecode tool
func (gpu *gpuParser) tool() string {
	return define.LsPci.String()
}

// module get current module param as dmi param
func (gpu *gpuParser) module() string {
	return define.GpuModule.Module()
}

// param indicate parser param to exec
func (gpu *gpuParser) param() string {
	args := []string{gpu.tool(), "|", "grep", "-i", gpu.module()}
	return strings.Join(args, " ")
}

type networkParser struct {
}

// parse used to parse gpu command line
func (net *networkParser) parse(info *define.BaseInfo, buf []byte) {
	if info == nil {
		return
	}
	msgSl := strings.Split(string(buf), ":")
	// check length, in case of panic
	if len(msgSl) < 2 {
		return
	}
	// get key
	key := msgSl[0]
	if strings.Contains(key, define.NetworkController.String()) {
		info.Model = strings.TrimLeft(msgSl[1], " ")
	}
}

// tool memory info use dmidecode tool
func (net *networkParser) tool() string {
	return define.LsPci.String()
}

// module get current module param as dmi param
func (net *networkParser) module() string {
	return define.NetModule.Module()
}

// param indicate parser param to exec
func (net *networkParser) param() string {
	args := []string{net.tool(), "|", "grep", "-i", net.module()}
	return strings.Join(args, " ")
}

type etherParser struct {
}

// parse used to parse ethernet command line
func (ether *etherParser) parse(info *define.BaseInfo, buf []byte) {
	if info == nil {
		return
	}
	msgSl := strings.Split(string(buf), ":")
	// check length, in case of panic
	if len(msgSl) < 3 {
		return
	}
	// get key
	key := msgSl[1]
	if strings.Contains(key, define.EthernetController.String()) {
		info.Model = strings.TrimLeft(msgSl[2], " ")
	}
}

// tool memory info use lspci tool
func (ether *etherParser) tool() string {
	return define.LsPci.String()
}

// module get current module param as dmi param
func (ether *etherParser) module() string {
	return define.EtherModule.Module()
}

// param indicate parser param to exec
func (ether *etherParser) param() string {
	args := []string{ether.tool(), "|", "grep", "-i", ether.module()}
	return strings.Join(args, " ")
}
