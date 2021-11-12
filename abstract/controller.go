package abstract

import "github.com/ArisAachen/experience/define"

// BaseController controller to control module's behavior
type BaseController interface {
	// Wake and Release allow
	Wake()
	Check()
	Release()
}

// BaseObserver observe event happens, and call ref func
type BaseObserver interface {
	Call(event define.ObserveEvent)
	Register(event define.ObserveEvent, caller define.Caller)
	Remove(event define.ObserveEvent)
}
