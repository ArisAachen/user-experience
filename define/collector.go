package define

// dbus object
const (
	ServiceName   = "com.deepin.UserExperience.Daemon"
	ServicePath   = "/com/deepin/UserExperience/Daemon"
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
