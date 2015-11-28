package hulu

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httputil"
	"strings"

	"github.com/elazarl/goproxy"
	"github.com/jd1123/adproxy/modules"
)

var resplog = "/etc/adproxy/log/resplog.log"
var debuglog = "/etc/adproxy/log/huludebug.log"

// URI strings to filter
// took out "ads"
var filterStrings = []string{"insightexpressai.com", "imrworldwide.com", "doubleverify.com", "scorecardresearch.com", "rewardtv.com", "flurry.com", "doubleclick", "adServer"}

type Hulu struct {
	Metadata modules.MetaStruct
	Quiet    bool
}

func NewHulu(quiet bool) *Hulu {
	ms := modules.NewMetaStruct("Hulu Filter", "0.1", "Hulu")
	return &Hulu{*ms, quiet}
}

func (h Hulu) Init() {
	h.Metadata = modules.MetaStruct{"Hulu Filter", "", ""}
	h.Metadata.ModuleName = "Hulu Filter"
	h.Metadata.VersionNumber = "0.1"
	h.Metadata.Service = "Hulu"
}

// Block filterStrings from sending a response
func (h Hulu) FilterResponse(resp *http.Response, ctx *goproxy.ProxyCtx) *http.Response {
	for _, i := range filterStrings {
		if strings.Contains(resp.Request.URL.String(), i) {
			if !h.Quiet {
				fmt.Println("RESPONSE FILTER: Adserver found in", h.Metadata.ModuleName, "...blocking:", resp.Request.URL.String())
			}
			bb := ClosingBuffer{bytes.NewBufferString("0")}
			resp.Body = bb
		}
	}

	if strings.Contains(resp.Request.URL.String(), "s.hulu.com") {
		fmt.Println("The Response filter found the s.hulu.com domain")
		//b := checkBody(&resp.Body)
		//fmt.Println(b)
		//PrintResponse(*resp)
		b, _ := httputil.DumpResponse(resp, true)
		//fmt.Println(string(b))
		go LogToFile(resplog, b)
	}
	b, _ := httputil.DumpResponse(resp, false)
	go LogToFile(debuglog, b)

	return resp
}

func (h Hulu) FilterRequest(req *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {
	for _, i := range filterStrings {
		if strings.Contains(req.URL.String(), i) {
			if !h.Quiet {
				fmt.Println("REQUEST FILTER: Adserver found in", h.Metadata.ModuleName, "...blocking:", req.URL.String())
			}
			c := CreateResponse(req)
			return req, &c
		}
	}

	b := checkBody(&req.Body)
	if strings.Contains(b, "adstate") {
		buf, _ := httputil.DumpRequest(req, true)
		fmt.Println("The request filter found a request body with an adstate string")
		//fmt.Println(string(buf))
		go LogToFile(resplog, buf)
	}
	bf, _ := httputil.DumpRequest(req, true)
	go LogToFile(debuglog, bf)
	return req, nil
}
