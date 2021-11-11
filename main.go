package launch

import "github.com/ArisAachen/experience/launch"

/*
	launch
	This module is use for start project.
	It is the main part of this project.
*/

func main() {

	// create launch
	lch := launch.NewLaunch()

	// init launch and add modules
	lch.Init()
	lch.AddWriterItemModules()
	lch.AddQueueItemModules()
	lch.AddConfigItemModules()

	// start service
	lch.StartService()
}
