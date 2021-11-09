package define

// dbus object
const (
	ServiceName   = "com.deepin.UserExperience.Daemon"
	ServicePath   = "/com/deepin/UserExperience/Daemon"
	DbusInterface = ServiceName
)

// SysModule indicate system info file
type SysModule string

const (
	// CpuModule save all cpu info file
	CpuModule   SysModule = "CpuModule"
	BoardModule SysModule = "BoardModule"
)

const (
	// Processor indicate using processor as param to read cpu info
	processor = "processor"
	baseboard = "baseboard"

)

// String check if now file path is valid and convert to string
func (info SysModule) String() string {
	// should check file here
	switch info {
	case CpuModule:
		return string(info)
	default:
	}
	return ""
}

// Module get ref module name as param of dmidecode
func (info SysModule) Module() string {
	switch info {
	case CpuModule:
		return processor
	case BoardModule:
		return baseboard
	}
	return ""
}

// SysInfoKey indicate system hardware info key, such as "Processor"
type SysInfoKey string

const (
	Gene SysModule = "Version"

	// ProcessorVersion and ProcessorId is read key of cpu info
	ProcessorVersion SysInfoKey = "Version"
	ProcessorId      SysInfoKey = "ID"

	// BoardProductName  and BoardSerialNumber is read key of base board
	BoardProductName  SysModule = "ProductName"
	BoardSerialNumber SysModule = "SerialNumber"
)

// String check if system key is valid and convert to string
func (key SysInfoKey) String() string {
	switch key {
	case ProcessorVersion, ProcessorId:
		return string(key)

	default:
	}
	return ""
}

// CpuInfo Cpu Info
type CpuInfo struct {
	Model string
	Id    string
}

type BaseInfo struct {
	Model string
	Id    string
}
