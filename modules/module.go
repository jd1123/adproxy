package modules

import (
	"net/http"

	"github.com/elazarl/goproxy"
)

/*
The module interface
*/

type MetaData struct {
	ModuleName    string
	VersionNumber string
	Service       string
}

type Module interface {
	FilterRequest(req *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response)
	FilterResponse(resp *http.Response, ctx *goproxy.ProxyCtx) *http.Response
}
