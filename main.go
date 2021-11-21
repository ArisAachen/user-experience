package main

import (
	"github.com/ArisAachen/experience/launch"
	"pkg.deepin.io/lib/dbusutil"
)

/*
	launch
	This module is use for start project.
	It is the main part of this project.
*/

func main() {
	// create launch
	lch := launch.NewLaunch()

	// create system service
	sysService, err := dbusutil.NewSystemService()
	if err != nil {
		return
	}

	// init launch and add modules
	lch.Init(sysService)

	// start service
	lch.StartService()
	sysService.Wait()
}
