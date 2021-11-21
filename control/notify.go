package control

import (
	"os"

	"github.com/ArisAachen/experience/define"
	"gopkg.in/fsnotify.v1"
)

type Notify struct {
	watch *fsnotify.Watcher
}

// NewNotify notify
func NewNotify() *Notify {
	ny := &Notify{

	}
	return ny
}

func (ny *Notify) Init() error {
	// create notify
	var err error
	ny.watch, err = fsnotify.NewWatcher()
	if err != nil {
		return err
	}

	return nil
}

// Monitor use to monitor message
func (ny *Notify) Monitor() error {
	// add monitor file
	err := ny.watch.Add(define.SqlitePath)
	if err != nil {
		return err
	}
	return nil
}

// Handle to call
func (ny *Notify) Handle(observer *Observer) {
	defer ny.watch.Close()
	for {
		select {
		case event, ok := <-ny.watch.Events:
			if !ok {
				continue
			}
			if event.Op != fsnotify.Write {
				continue
			}
			// check file
			stat, err := os.Stat(define.SqlitePath)
			if err != nil {
				continue
			}
			if stat.Size() < 50*1024 {
				continue
			}
			observer.Call(define.ObServerDatabase)
		case _ = <-ny.watch.Errors:
			return
		}
	}
}
