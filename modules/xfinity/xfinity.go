package xfinity

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
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
	flag := 0
	for _, i := range filterStrings {
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

	//b := checkBody(&req.Body)

	return req, nil
}

func checkBody(rc *io.ReadCloser) string {
	var buf []byte
	buf, _ = ioutil.ReadAll(*rc)
	b := string(buf)
	*rc = ioutil.NopCloser(bytes.NewBuffer(buf))
	return b
}

// struct to implement io.ReadCloser
type ClosingBuffer struct {
	io.Reader
}

func (cb ClosingBuffer) Close() (err error) {
	return
}
