package common

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"os/exec"
	"strings"

	"github.com/ArisAachen/experience/define"
)

// comm Module
// use dmidecode command to get system info message
// man doc: https://linux.die.net/man/8/dmidecode
// file dir: /sys/class/dmi/id

// GetCpuInfo use dmidecode get cpu info
func GetCpuInfo() (define.BaseInfo, error) {
	return general(define.CpuModule)
}

// GetBaseBoardInfo use dmidecode get base board info
func GetBaseBoardInfo() (define.BaseInfo, error) {
	return general(define.BoardModule)
}

func general(file define.SysModule) (define.BaseInfo, error) {
	// init cpu info
	var info define.BaseInfo
	var parser baseCommand

	// check if file is valid, and construct base parser
	switch file {
	case define.CpuModule:
		parser = &cpuCommand{}
	case define.BoardModule:
		parser = &boardCommand{}
	default:
		return info, errors.New("file is not valid")
	}

	// use dmi command read cpu info
	// var byErr bytes.
	args := []string{"dmidecode", "-t", file.Module()}
	cmd := exec.Command("/bin/bash", "-c", strings.Join(args, " "))
	// run command
	buffer, err := cmd.CombinedOutput()
	if err != nil {
		return info, err
	}
	// read line search for cpu
	by := bytes.NewBuffer(buffer)
	reader := bufio.NewReader(by)

	// read line to find cpu info
	for {
		// read line file
		buf, _, err := reader.ReadLine()
		if err != nil {
			// read end file, means info is not fully fill until file end,
			// since now, it is error
			if err == io.EOF {
				return info, errors.New("cant fill info until file end")
			}
			// read file end
			return info, err
		}
		// parse buf to info
		parser.parse(&info, buf)

		// either model or id is not fix, continue to read line
		if info.Model == "" || info.Id == "" {
			continue
		}
		// all message has read successfully
		break
	}
	return info, nil
}
