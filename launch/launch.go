package launch

/*
	launch
	This module is use for start project.
	It is the main part of this project.

	1. export dbus message
*/


type Launch struct {
	basic exportBasic
}


func NewLaunch() *Launch {
	lch := &Launch{
		basic: new(experience),
	}
	return lch
}