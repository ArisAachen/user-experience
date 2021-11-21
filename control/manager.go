package control

import (
	"github.com/ArisAachen/experience/define"
)

type Observer struct {
	events map[define.ObserveEvent]define.Caller
}

func NewObserver() *Observer {
	ob := &Observer{
		events: make(map[define.ObserveEvent]define.Caller),
	}
	return ob
}

// Register register event handler
func (ob *Observer) Register(event define.ObserveEvent, caller define.Caller) {
	ob.events[event] = caller
}

// Call call when event happens
func (ob *Observer) Call(event define.ObserveEvent) {
	caller, ok := ob.events[event]
	if !ok {
		return
	}
	// TODO: check if args is match to Call
	caller.Method.Call(caller.Args)
}

// Remove caller
func (ob *Observer) Remove(event define.ObserveEvent) {
	delete(ob.events, event)
}
