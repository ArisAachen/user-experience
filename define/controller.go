package define

import "reflect"

type Caller struct {
	Method reflect.Value
	Args   []reflect.Value
}

type ObserveEvent string

const (
	ObServerDatabase ObserveEvent = "start push database"
)

// Rule use to controller writers behavior
// in strict rule, no data will be written
// in gentle rule, only special data would be sent
// in loose rule, all data can be sent
type Rule int

// strict: update interface and hardware uni id, also update interface should be sent first
// gentle: experience state, this data should sent after interface update, always should be sent
// loose: login and logout should be sent at head

const (
	NoneRule Rule = iota
	LooseRule
	GentleRule
	StrictRule
)
