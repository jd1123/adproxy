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
	filterRequest(req *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response)
	filterResponse(resp *http.Response, ctx *goproxy.ProxyCtx) *http.Response
}

func RegisterModule(m Module) {
	mods = append(mods, m)
}
