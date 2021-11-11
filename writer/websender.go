package writer

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/ArisAachen/experience/define"
)

// receive message from web server
type rcvMsg struct {
	Code int
	Msg  string
	Data json.RawMessage
}

const (
	postSuccess = 200
)

// webWriterItem use to send data to server, call handler callback to handle result.
// When post data failed, retry 3 times, if none success, drop data.
// System data should send immediately, so these data dont need save to database, only if data send failed
type webWriterItem struct {
	client http.Client
}

// NewWebWriter create one web writer
func newWebWriter() *webWriterItem {
	clt := &webWriterItem{
		client: http.Client{
			Timeout: 500 * time.Millisecond,
		},
	}
	return clt
}

// Write write message to web
func (web *webWriterItem) Write(url string, msg string) define.WriteResult {
	var result define.WriteResult
	reader := strings.NewReader(msg)
	// post data
	resp, err := web.client.Post(url, "application/json", reader)
	// post data failed at this time
	if err != nil {
		result.ResultCode = define.WriteResultWriteFailed
		logger.Warningf("post data failed, err: %v", err)
		// at this while, should write data back to database
		return result
	}
	defer resp.Body.Close()
	// post data failed, due to network env
	if resp.StatusCode != postSuccess {
		result.ResultCode = define.WriteResultWriteFailed
		logger.Warningf("post data failed, status code: %v", resp.StatusCode)
		return result
	}
	// read message from body
	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		result.ResultCode = define.WriteResultReadBodyFailed
		logger.Warningf("read body failed, err: %v", err)
		return result
	}
	// unmarshal receive data
	var rcv rcvMsg
	err = json.Unmarshal(buf, &rcv)
	if err != nil {
		result.ResultCode = define.WriteResultReadBodyFailed
		logger.Warningf("unmarshal received body failed, err: %v", err)
		return result
	}
	logger.Debugf("receive response web server, code: %v, msg: %v", rcv.Code, rcv.Msg)
	// parse response code
	// it is ok here, dont use state model at all
	switch define.RespCode(rcv.Code) {
	case define.RespSuccess:
		result.ResultCode = define.WriteResultSuccess
		result.Msg = rcv.Data
		logger.Debug("data has sent successfully")
	case define.RespVfnInvalid:
		result.ResultCode = define.WriteResultParamInvalid
		logger.Warning("data drop by web server, verification is not valid")
	case define.RespParamInvalid:
		result.ResultCode = define.WriteResultParamInvalid
		logger.Warning("data drop by web server, param is not valid")
	default:
		result.ResultCode = define.WriteResultUnknown
		logger.Warningf("web server response unknown code: %v", rcv.Code)
	}
	return result
}