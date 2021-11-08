package launch

import (
	"errors"

	"github.com/ArisAachen/experience/define"
	"github.com/godbus/dbus"
)

type experience struct {
	svc *dbus.Conn
}

// create experience obj
func newExp(svc *dbus.Conn) (*experience,error) {
	// check if svc is nil
	if svc == nil {
		return nil,errors.New("dbus service is nil")
	}
	// create service
	exp := &experience{
		svc: svc,
	}
	return exp,nil
}

// GetInterfaceName dbus implement
func (ex *experience)GetInterfaceName()string{
	return define.DbusInterface
}

// export dbus method
func (ex *experience)export()error {
	// check if system service is nil
	if ex.svc == nil {
		return errors.New("dbus service is nil")
	}
	// export dbus path
	err := ex.svc.Export(ex,define.ServicePath,define.DbusInterface)
	if err != nil {
		return err
	}
	// request name
	_, err = ex.svc.RequestName(define.ServicePath,dbus.NameFlagDoNotQueue)
	if err != nil {
		return err
	}
	return nil
}