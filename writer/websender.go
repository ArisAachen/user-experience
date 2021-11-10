package writer

import (
	"net/http"
	"strings"
	"time"

	"github.com/ArisAachen/experience/queue"
)

type WebWriter struct {
	client http.Client
}

// NewWebWriter create one web writer
func NewWebWriter() *WebWriter {
	clt := &WebWriter{
		client: http.Client{
			Timeout: 500 * time.Millisecond,
		},
	}

	return clt
}

// Write write message to web
func (web *WebWriter) Write(handler queue.BaseQueueHandler, msg string) {
	reader := strings.NewReader(msg)
	// post data
	resp, err := web.client.Post(handler.GetInterface(), "application/json", reader)
	// post data failed at this time
	if err != nil {
		logger.Warningf("post data failed, err: %v", err)
		// at this while, should write data back to database
		handler.Handler(nil, msg)
		return
	}
	defer resp.Body.Close()
}
