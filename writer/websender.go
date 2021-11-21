package writer

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/ArisAachen/experience/abstract"
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

// WebWriterItem use to send data to server, call handler callback to handle result.
// When post data failed, retry 3 times, if none success, drop data.
// System data should send immediately, so these data dont need save to database, only if data send failed
type WebWriterItem struct {
	client http.Client
}

// NewWebWriter create one web writer
func NewWebWriter() *WebWriterItem {
	clt := &WebWriterItem{
		client: http.Client{
			Timeout: 5 * time.Second,
		},
	}
	return clt
}

// Connect connect to web
func (web *WebWriterItem) Connect(url string) {
	return
}

// Disconnect disconnect from web
func (web *WebWriterItem) Disconnect() {

}

// GetRemote get remote url path, maybe use for check web state later
func (web *WebWriterItem) GetRemote() string {
	return ""
}

// Write write message to web
func (web *WebWriterItem) Write(crypt abstract.BaseCryptor, url string, msg []string) define.WriteResult {
	var result define.WriteResult
	req := define.PostRequest{
		ReqTime: time.Now().UnixNano() / 1e6,
		Uni:     "",
		Data:    msg,
	}
	// marshal request
	buf, err := json.Marshal(&req)
	if err != nil {
		result.ResultCode = define.WriteResultUnknown
		logger.Warningf("marshal request message failed, err: %v", err)
		return result
	}
	logger.Debugf("post request message is %v", req)
	// use Cryptor to crypt data
	cryResult, err := crypt.Encode(string(buf))
	if err != nil {
		// when data encrypt failed, just drop this data
		// also this module can return to handler, if some special handle is needed
		logger.Warningf("failed to crypt data, err: %v", err)
		return result
	}
	reader := strings.NewReader(cryResult.Data)
	// post data
	// TODO should optimize
	url = url + "&key=" + cryResult.Key
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
	buf, err = ioutil.ReadAll(resp.Body)
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
	// webserver response success, data is valid and accepted by server successfully
	case define.RespSuccess:
		result.ResultCode = define.WriteResultSuccess
		result.Msg = rcv.Data
		logger.Debug("data has sent successfully")
	// web verification failed
	case define.RespVfnInvalid:
		result.ResultCode = define.WriteResultParamInvalid
		logger.Warning("data drop by web server, verification is not valid")
	// webserver has receive data, but data is dropped because param is not valid
	case define.RespParamInvalid:
		result.ResultCode = define.WriteResultParamInvalid
		logger.Warning("data drop by web server, param is not valid")
	// not sure what happens, maybe webserver response unexpected message
	default:
		result.ResultCode = define.WriteResultUnknown
		logger.Warningf("web server response unknown code: %v", rcv.Code)
	}
	return result
}

func (web *WebWriterItem) GetWriterItemName() define.WriterItemModule {
	return define.WebItemWriter
}
