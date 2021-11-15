package abstract

import "github.com/ArisAachen/experience/define"

// BaseController controller to control module's behavior
// now hve two controller,
// 1. sender controller is use to limit send data to web,
// sometimes, we need to wait for op, or dont allow to send all data, so use sender controller
// 2. database controller is use to limit the size of db file
// when data is out of range, should delete old data, then add new ones.
type BaseController interface {
	// Invoke and Release is use to add and remove limit condition
	// also low level release rule cant release high invoke rule
	Invoke(invoke define.Rule)
	Release(release define.Rule)

	// Monitor use to monitor current op, sometimes would be block
	Monitor(monitor define.Rule)
}

// BaseObserver observe event happens, and call ref func
type BaseObserver interface {
	Call(event define.ObserveEvent)
	Register(event define.ObserveEvent, caller define.Caller)
	Remove(event define.ObserveEvent)
}
