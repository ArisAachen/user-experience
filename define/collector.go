package define

// dbus object
const (
	ServiceName   = "com.deepin.userexperience.Daemon"
	ServicePath   = "/com/deepin/userexperience/Daemon"
	DbusInterface = ServiceName
)

// AppEntry use to store current open/close app message
type AppEntry struct {
	Tid     TidTyp
	Time    int64
	Id      int
	PkgName string
	AppName string
}

// LogInfo login logout
type LogInfo struct {
	Tid  TidTyp
	Time int64
}
