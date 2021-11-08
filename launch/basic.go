package launch

/*
	basic is the interface of dbus obj in this module,
	in case more experience module should be export.
*/

type exportBasic interface {
	export() error
}

