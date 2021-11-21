package define

const (
	// Processor indicate using processor as param to read cpu info
	processor = "processor"
	baseboard = "baseboard"
	memory    = "memory"
	disk      = "disk"
	vga       = "VGA"
	network   = "Network"
	ether     = "Ethernet"
)

// SysModule indicate system info file
type SysModule string

const (
	// CpuModule save all cpu info file
	CpuModule    SysModule = "CpuModule"
	BoardModule  SysModule = "BoardModule"
	MemoryModule SysModule = "MemoryModule"
	DiskModule   SysModule = "DiskModule"
	GpuModule    SysModule = "GpuModule"
	NetModule    SysModule = "NetworkModule"
	EtherModule  SysModule = "EtherModule"
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
	case MemoryModule:
		return memory
	case DiskModule:
		return disk
	case GpuModule:
		return vga
	case NetModule:
		return network
	case EtherModule:
		return ether
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
	BoardProductName  SysInfoKey = "ProductName"
	BoardSerialNumber SysInfoKey = "SerialNumber"

	// MemoryMaximumCapacity is read key of memory
	MemoryMaximumCapacity SysInfoKey = "MaximumCapacity"

	// VgaController is read key of VGA
	VgaController SysInfoKey = "VGA compatible controllerl"

	// EthernetController is read key of Ethernet
	EthernetController SysInfoKey = "Ethernet controller"

	// NetworkController is read key of Network
	NetworkController SysInfoKey = "Network controller"
)

// String check if system key is valid and convert to string
func (key SysInfoKey) String() string {
	switch key {
	case ProcessorVersion, ProcessorId, BoardProductName,
		BoardSerialNumber, MemoryMaximumCapacity, VgaController,
		NetworkController, EthernetController:
		return string(key)
	default:
	}
	return ""
}

// Tool indicate exec tool to get system info
type Tool string

const (
	DmiDecode Tool = "dmidecode"
	LsBlk     Tool = "lsblk"
	LsPci     Tool = "lspci"
)

func (t Tool) String() string {
	switch t {
	case DmiDecode, LsBlk, LsPci:
		return string(t)
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

// ConfigItemModule config item
type ConfigItemModule string

const (
	HardwareItemConfig ConfigItemModule = "hardware"
	PostItemConfig     ConfigItemModule = "post"
	SystemItemConfig   ConfigItemModule = "system"
)
