package collect

import (
	"github.com/ArisAachen/experience/abstract"
)

type Collector struct {
	items map[string]abstract.BaseCollectorItem
}

func NewCollector() *Collector {
	col := &Collector{

	}
	return col
}

// AddModule add module
func (col *Collector) AddModule(name string, item abstract.BaseCollectorItem) {
	// check if already exist collector
	if _, ok := col.items[name]; ok {
		return
	}
	// save collector
	col.items[name] = item
}

// Collect begin to collect message
func (col *Collector) Collect(que abstract.BaseQueue) {
	// loop all item to collect data
	for _, item := range col.items {
		// use go routine to collect data
		go item.Collect(que)
	}
}
