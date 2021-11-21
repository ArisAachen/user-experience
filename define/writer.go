package define

import (
	"encoding/json"
)

// File path
const (
	BaseCfgDir  = "/var/lib/deepin-user-experience"
	HwCfgFile   = BaseCfgDir + "/" + "hardware"
	SysCfgFile  = BaseCfgDir + "/" + "system"
	PostCfgFile = BaseCfgDir + "/" + "post"
	SqlitePath  = BaseCfgDir + "/" + "exp.db"
	EncryptPath = BaseCfgDir + "/" + "pem"

	BaseEtcDir  = "/etc"
	MachineFile = BaseEtcDir + "/" + "machine-id"
	SysTypFile  = BaseEtcDir + "/" + "os-version"

	AptTokenFile = BaseEtcDir + "/" + "apt/apt.conf.d/99lastore-token.conf"

	PkgName = "deepin-user-experience-daemon"
)

// WriteResultCode code indicate diff write result
type WriteResultCode int

const (
	WriteResultSuccess WriteResultCode = iota

	WriteResultWriteFailed
	WriteResultVfnFailed
	WriteResultParamInvalid
	WriteResultReadBodyFailed

	WriteParseQueryFailed

	WriteResultUnknown
)

// WriteUpdateReq write update request
type WriteUpdateReq struct {
	Tid     TidTyp `json:"tid"`
	Machine string `json:"machine"`
	Type    string `json:"sys_type"`
	Version string `json:"version"`
}

// WriteOrigin origin data, use to decode
type WriteOrigin struct {
	Tid TidTyp
	Msg json.RawMessage
}

type WriteResult struct {
	ResultCode WriteResultCode
	Origin     string
	Msg        json.RawMessage
}

// RespCode indicate if data has write success by web server or database
type RespCode int

const (
	RespSuccess RespCode = iota
	RespVfnInvalid
	RespParamInvalid
)

// WriterItemModule use to select add writer item into writer
type WriterItemModule string

const (
	WebItemWriter      WriterItemModule = "web sender"
	DataBaseItemWriter WriterItemModule = "database writer"
)

// QueueItemModule use to select add queue item into queue
type QueueItemModule string

const (
	WebItemQueue      QueueItemModule = "web queue"
	DataBaseItemQueue QueueItemModule = "database queue"
)

// RequestMsg the message to write,
// message has priority, highest priority must sent at first
// now priority is
// 1. update interface
// 2. update uni id
// 3. exp-enabled state
// 4. loin and logout
// 5. simple data

type RequestMsg struct {
	Rule Rule
	Pri  RequestLevel
	Msg  []string
}

// RequestLevel priority to send data
type RequestLevel int

const (
	UpdateIfcRequest RequestLevel = iota
	UpdateUniRequest
	ExpStateRequest
	LogInOutRequest
	SimpleRequest
	NoneRequest
)

// Web Message

// RcvInterface receive update post interface message from remote
type RcvInterface struct {
	Ip     string `json:"ip"`
	Update int64  `json:"update"`
}

// RcvUni receive uni id and
type RcvUni struct {
	Uni string `json:"uuid"`
}

// LogEvent login shutdown
type LogEvent string

const (
	LoginEvent    LogEvent = "login"
	LogOutEvent   LogEvent = "logout"
	ShutDownEvent LogEvent = "shutdown"
)

// PostRequest post request
type PostRequest struct {
	ReqTime int64    `json:"rt"`
	Uni     string   `json:"unid"`
	Data    []string `json:"data"`
}

// PostByte Post Byte Define
// define the format of post interface,
// now use 'Byte' to divide different post event
type PostByte string

const (
	/*
		This module include general message
		@Tid:      post byte only id
		@Kind:     post byte kind
		@Version:  posy byte version
		@DataTime: event happens time
		@ReqTime:  request send byte time
		@DevId:    device id
		@UserId:   user id
	*/
	Tid      PostByte = "tid"
	Kind     PostByte = "k"
	Version  PostByte = "v"
	DataTime PostByte = "dt"
	ReqTime  PostByte = "rt"
	DevId    PostByte = "did"
	UserId   PostByte = "uid"

	// Ip post update interface need ip
	Ip PostByte = "ip"

	// HwInfo post hardware info need info
	HwInfo PostByte = "info"

	// Order post user state, such as user-experience-enabled
	Order PostByte = "oder"

	/*
		This module include un/install open/close app message
		@AppPkgName: app package name
		@AppId:      app id
		@AppPkgSize: app package size
		@AppVersion: app version
		@AppOther:   app other info
	*/
	AppPkgName PostByte = "an"
	AppId      PostByte = "ai"
	AppPkgSize PostByte = "as"
	AppVersion PostByte = "av"
	AppOther   PostByte = "ao"

	/*
		@AppDlStartTime:   app download start time
		@AppDlEndTime:     app download end time
		@AppInsStartTime:  app install start time
		@AppInsEndTime:    app install end time
	*/
	AppDlStartTime  PostByte = "dst"
	AppDlEndTime    PostByte = "det"
	AppInsStartTime PostByte = "ist"
	AppInsEndTime   PostByte = "iet"
)

// TidTyp define unique post type id
type TidTyp int

const (
	GeneralTid TidTyp = iota
	MotherBoardTid
	ProcessorTid
	DisplayCardTid
	VideoCardTid
	MachineBranchTid
	MemoryTid
	DiskTid
	NetworkCardTid
	WirelessCardTid
	MouseTid
	KeyboardTid
	UosEdition
	UosMajorTid
	IpAddrTid
	MacAddrTid
	MachineIdTid

	LoginTid
	LogoutTid

	UpgradeTid
	DevelopModeTid
	ExpPlanTid

	InstallAppTid TidTyp = iota + 26
	UninstallAppTid
	OpenAppTid
	CloseAppTid

	UpdateNodeTid

	ShutDownTid
	RebootTid

	UosMinorTid
	UosBuild
	UosProduct

	NewSystemInfoTid  TidTyp = 1000
	NewCheckUpdateTid TidTyp = 1001
	NewLoginTid       TidTyp = 1002
	NewLogoutTid      TidTyp = 1003
	NewAppOpenTid     TidTyp = 1004
	NewAppCloseTid    TidTyp = 1005
)

const (
	Sqlite3Driver = "sqlite3"
)
