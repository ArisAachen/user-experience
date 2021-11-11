package define

import "encoding/json"

// dbus object
const (
	ServiceName   = "com.deepin.UserExperience.Daemon"
	ServicePath   = "/com/deepin/UserExperience/Daemon"
	DbusInterface = ServiceName
)

const (
	BaseCfgFile = "/var/lib/deepin-user-experience"
	HwCfgFile   = BaseCfgFile + "hardware"
)


// WriteResultCode code indicate diff write result
type WriteResultCode int

const (
	WriteResultSuccess WriteResultCode = iota

	WriteResultWriteFailed
	WriteResultVfnFailed
	WriteResultParamInvalid
	WriteResultReadBodyFailed

	WriteResultUnknown
)

type WriteResult struct {
	ResultCode WriteResultCode
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

