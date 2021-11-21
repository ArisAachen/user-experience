package abstract

import "github.com/ArisAachen/experience/define"

// BaseUrlCreator create url paths and interface
type BaseUrlCreator interface {
	GetRandomPostUrls() []string
	GetInterface(tid define.TidTyp) string
	GetAid() string
}
