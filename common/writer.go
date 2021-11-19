package common

import "github.com/ArisAachen/experience/define"

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