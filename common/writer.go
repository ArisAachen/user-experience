package common

import (
	"github.com/ArisAachen/experience/define"
	"io/ioutil"
	"math/rand"
	"os/exec"
	"pkg.deepin.io/lib/keyfile"
	"strings"
	"time"
)

// QueryLevel query tid level
// when collect data from database, should covert data to request level
func QueryLevel(tid define.TidTyp) define.RequestLevel {
	// TODO should optimize here
	if isTidExp(tid) {
		return define.ExpStateRequest
	} else if isTidLogOp(tid) {
		return define.LogInOutRequest
	}
	return define.SimpleRequest
}

// isTidLogOp check if tid type is login/out operation
func isTidLogOp(tid define.TidTyp) bool {
	if tid == define.LogoutTid || tid == define.LoginTid || tid == define.ShutDownTid || tid == define.RebootTid {
		return true
	}
	return false
}

// isTidExp check if type is user experience enabled state
func isTidExp(tid define.TidTyp) bool {
	if tid == define.ExpPlanTid {
		return true
	}
	return false
}

// UpdatePackage update package
func UpdatePackage(pkg string) (string, error) {
	// first of all run apt update
	req := []string{"apt", "update"}
	cmd := exec.Command("/bin/bash", "-c", strings.Join(req, " "))
	// run apt update
	buf, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}
	// run apt install
	req = []string{"apt", "install", pkg}
	cmd = exec.Command("/bin/bash", "-c", strings.Join(req, " "))
	// run apt install
	buf, err = cmd.CombinedOutput()
	if err != nil {
		return "", err
	}
	return string(buf), nil
}

// Shuffle random list
func Shuffle(list []string) []string {
	// rand seed
	rand.Seed(int64(time.Now().Unix()))
	temptList := list
	// shuffle
	for key, value := range temptList {
		// get random key
		nRand := getRandomInt(0, key)
		// exchange elem
		tempt := value
		temptList[key] = temptList[nRand]
		temptList[nRand] = tempt
	}
	return temptList
}

// getRandomInt random seed
func getRandomInt(min, max int) int {
	// check if params is valid
	if max == 0 {
		return 0
	}
	// random
	nRound := min + rand.Intn(max-min)
	return nRound
}

// GetMachineId get machine id from file
func GetMachineId() (string, error) {
	// read machine id file
	machine, err := ioutil.ReadFile(define.MachineFile)
	if err != nil {
		return "", err
	}
	return string(machine), nil
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
