package control

import (
	"sync"

	"github.com/ArisAachen/experience/define"
)

type State int

const (
	// Strict the hardest rule, in this case, no data will be sent
	// this state has 2 conditions: 1. post update request 2. post hardware uni id request
	Strict State = iota
	// Simple in this state, only special priority data will be sent
	// this state has 1 condition: 1. must
	Simple
	// Loose in this state, all data can be sent
	Loose
)

// WebController controller if should send data immediately or wait
// current rule is following:
// 1. no data will be sent, conditions : 1. update post interface 2. update uni id
// 2. only part data state allow sent, conditions: 1.
// 2. all data should not be sent if now network is online
// 3. all data will not sent or collected when user-experience is not on
type WebController struct {
	state State
	cond  sync.Cond
}

// Wait decide is should block here
func (wb *WebController) Wait(priority define.RequestPriority) {
	if wb.state == Strict || (wb.state == Simple && priority != define.Must) {
		wb.cond.Wait()
	}
}
