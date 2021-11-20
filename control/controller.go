package control

import (
	"sync"

	"github.com/ArisAachen/experience/define"
)

// Controller controller if should send data immediately or wait
// current rule is following:
// 1. no data will be sent, conditions : 1. update post interface 2. update uni id
// 2. only part data state allow sent, conditions: 1.
// 2. all data should not be sent if now network is online
// 3. all data will not sent or collected when user-experience is not on
type Controller struct {
	rule define.Rule
	cond sync.Cond
}

// NewController create controller
func NewController() *Controller {
	col := &Controller{
		rule: define.NoneRule,
		cond: sync.Cond{
			L: new(sync.Mutex),
		},
	}
	return col
}

// Invoke just set rule state, rule will block until monitor level check
func (wb *Controller) Invoke(invoke define.Rule) {
	// check if coming rule is more strictly
	if wb.rule >= invoke {
		return
	}
	// save state
	wb.rule = invoke
}

// Release if release level if strict than current, should release wait
func (wb *Controller) Release(release define.Rule) {
	// check if coming rule is more strictly
	if wb.rule >= release {
		return
	}
	// release current rule, and emit signal
	wb.rule = define.NoneRule
	wb.cond.Signal()
}

// Monitor if current rule is more strictly, should wait until release
func (wb *Controller) Monitor(monitor define.Rule) {
	// should wait
	if monitor <= wb.rule {
		wb.cond.L.Lock()
		wb.cond.Wait()
		wb.cond.L.Unlock()
	}
}
