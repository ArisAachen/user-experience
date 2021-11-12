package define

import "reflect"

type Caller struct {
	Method reflect.Value
	Args   []reflect.Value
}

type ObserveEvent string
