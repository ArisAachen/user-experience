package common

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"os/exec"
	"strings"

	"github.com/ArisAachen/experience/define"
	"pkg.deepin.io/lib/keyfile"
)

// comm Module
// use command to get system info

// GetCpuInfo use dmidecode get cpu info
func GetCpuInfo() (define.BaseInfo, error) {
	return general(define.CpuModule)
}

// GetBaseBoardInfo use dmidecode get base board info
func GetBaseBoardInfo() (define.BaseInfo, error) {
	return general(define.BoardModule)
}

// GetMemoryInfo use dmidecode to get memory
func GetMemoryInfo() (define.BaseInfo, error) {
	return general(define.MemoryModule)
}

// GetDiskInfo use lsblk to get disk info
func GetDiskInfo() (define.BaseInfo, error) {
	return general(define.DiskModule)
}

// GetGpuInfo use lspci to get gpu info
func GetGpuInfo() (define.BaseInfo, error) {
	return general(define.GpuModule)
}

// GetNetworkInfo use lspci to get network info
func GetNetworkInfo() (define.BaseInfo, error) {
	return general(define.NetModule)
}

// GetEtherInfo use lspci to get network info
func GetEtherInfo() (define.BaseInfo, error) {
	return general(define.EtherModule)
}

// GetMachineId get machine id from file
func GetMachineId() (string, error) {
	// read machine id file
	machine, err := ioutil.ReadFile(define.MachineFile)
	if err != nil {
		return "", err
	}
	result := strings.Trim(string(machine), "\n")
	return result, nil
}

// GetAptToken get apt token from file
func GetAptToken() (string, error) {
	// get token file
	buf, err := ioutil.ReadFile(define.AptTokenFile)
	if err != nil {
		return "", err
	}
	// split
	msgSl := strings.Split(string(buf), " ")
	if len(msgSl) != 2 {
		return "", errors.New("apt token file length invalid")
	}
	return msgSl[1], nil
}

// GetEdition get os edition
func GetEdition() (string, error) {
	// read os-version file to find
	kf := keyfile.NewKeyFile()
	err := kf.LoadFromFile(define.SysTypFile)
	if err != nil {
		return "", err
	}
	// read edition name
	typ, err := kf.GetString("Version", "EditionName")
	if err != nil {
		return "", err
	}
	return typ, nil
}

// general the general func to get system info
func general(file define.SysModule) (define.BaseInfo, error) {
	// init cpu info
	var info define.BaseInfo
	var factory parserFactory

	// check if file is valid, and construct base parser
	parser := factory.createParser(file)
	if parser == nil {
		return info, errors.New("file module is not exist")
	}

	// use dmi command read cpu info
	cmd := exec.Command("/bin/bash", "-c", parser.param())
	log.Println(cmd.Args)
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
