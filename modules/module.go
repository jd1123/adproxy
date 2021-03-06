package modules

import (
	"fmt"
	"net/http"

	"github.com/elazarl/goproxy"
)

/*
The module interface
*/

type MetaStruct struct {
	ModuleName    string
	VersionNumber string
	Service       string
}

func NewMetaStruct(name, ver, service string) *MetaStruct {
	return &MetaStruct{name, ver, service}
}

func (m MetaStruct) PrintMetaData() {
	fmt.Println("Module", m.ModuleName, "loaded...")
	fmt.Println("Version", m.VersionNumber)
	fmt.Println("Service", m.Service)
}

type Module interface {
	Init()
	FilterRequest(req *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response)
	FilterResponse(resp *http.Response, ctx *goproxy.ProxyCtx) *http.Response
}
