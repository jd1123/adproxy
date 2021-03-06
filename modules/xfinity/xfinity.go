package xfinity

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/elazarl/goproxy"
	"github.com/jd1123/adproxy/modules"
)

/*
	Xfinity module
*/

var filterStrings = []string{"adserver"}

type Xfinity struct {
	Metadata modules.MetaStruct
}

func NewXfinity() *Xfinity {
	ms := modules.NewMetaStruct("Xfinity Filter", "0.1", "Xfinity onDemand")
	return &Xfinity{*ms}
}

func (x Xfinity) Init() {
	x.Metadata.ModuleName = "Xfinity Filter"
	x.Metadata.VersionNumber = "0.1"
	x.Metadata.Service = "Xfinity onDemand"
}

func (x Xfinity) FilterResponse(resp *http.Response, ctx *goproxy.ProxyCtx) *http.Response {
	for _, i := range filterStrings {
		if strings.Contains(resp.Request.URL.String(), i) {
			fmt.Println("Adserver found... blocking: ", resp.Request.URL.String())
			bb := ClosingBuffer{bytes.NewBufferString("0")}
			resp.Body = bb
		}
	}
	return resp
}

func (x Xfinity) FilterRequest(req *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {
	// Block analytics requests
	if strings.Contains(req.URL.String(), "analytics.xcal.tv") {
		fmt.Println(req.URL.String(), "Analytics Request Intercepted...")
		return req, goproxy.NewResponse(req, goproxy.ContentTypeText, http.StatusOK, "0")
	}
	return req, nil
}

// struct to implement io.ReadCloser
type ClosingBuffer struct {
	io.Reader
}

func (cb ClosingBuffer) Close() (err error) {
	return
}
