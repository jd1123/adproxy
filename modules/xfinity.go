package modules

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/elazarl/goproxy"
)

/*
	Xfinity module
*/

var FILTER_STRINGS = []string{"adserver"}

type Xfinity struct {
	Name string
}

func (x Xfinity) filterResponse(resp *http.Response, ctx *goproxy.ProxyCtx) *http.Response {
	for _, i := range FILTER_STRINGS {
		if strings.Contains(resp.Request.URL.String(), i) {
			fmt.Println("Adserver found... blocking: ", resp.Request.URL.String())
			bb := ClosingBuffer{bytes.NewBufferString("0")}
			resp.Body = bb
		}
	}
	return resp
}

func (x Xfinity) filterRequest(req *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {
	flag := 0
	for _, i := range FILTER_STRINGS {
		if strings.Contains(req.URL.String(), i) {
			flag = 1
		}
	}
	if flag == 0 {
		log.Println("Req: ", req.Method, ": ", req.URL.String())
	}

	if strings.Contains(req.URL.String(), "analytics.xcal.tv") {
		fmt.Println(req.URL.String(), "Analytics Request Intercepted...")
		return req, goproxy.NewResponse(req, goproxy.ContentTypeText, http.StatusOK, "0")
	}

	b := checkBody(&req.Body)

	return req, nil
}
