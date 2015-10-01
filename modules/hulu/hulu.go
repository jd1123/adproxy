package hulu

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/elazarl/goproxy"
)

var filterStrings = []string{"insightexpressai.com", "imrworldwide.com", "doubleverify.com", "scorecardresearch.com",
	"ads", "rewardtv.com", "flurry.com", "doubleclick"}

type Hulu struct {
	Metadata MetaStruct
}

func (h Hulu) Init(){
}

// Block filterStrings from sending a response
func (h Hulu) FilterResponse(resp *http.Response, ctx *goproxy.ProxyCtx) *http.Response {
	for _, i := range filterStrings {
		if strings.Contains(resp.Request.URL.String(), i) {
			fmt.Println("Adserver found in ", h.MetaData.ModuleName, "... blocking: ", resp.Request.URL.String())
			bb := ClosingBuffer{bytes.NewBufferString("0")}
			resp.Body = bb
		}
	}
	return resp
}

func (h Hulu) FilterRequest(req *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {
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

func LoadFakeResponse() string {
	f, err := os.OpenFile("/etc/adproxy/dat/response", os.O_RDONLY, 0660)
	if err != nil {
		fmt.Println("File Error: ", err)
		os.Exit(1)
	}
	defer f.Close()
	buff, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println("File Error: ", err)
		os.Exit(1)
	}
	return string(buff)
}

func LoadFilterList() []string {
	f, err := os.OpenFile("/etc/adproxy/dat/filterlist.txt", os.O_RDONLY, 0660)
	if err != nil {
		fmt.Println("File Errror: ", err)
		os.Exit(1)
	}
	defer f.Close()
	buff, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println("File Error: ", err)
		os.Exit(1)
	}
	ns := make([]string, 0)
	s := strings.Split(string(buff), "\n")
	for i := range s {
		if s[i] != "" {
			ns = append(ns, s[i])
		}
	}
	return ns

}

func CreateResponse(req *http.Request) http.Response {
	n := time.Now()
	layout := "Sun, 09 Aug 1999 18:20:22 GMT"
	finalDate := n.Format(layout)
	nresp := goproxy.NewResponse(req, ContentTypeJSON, http.StatusOK, FAKERESPONSE)
	nresp.Status = "200 OK"
	nresp.ProtoMajor = 1
	nresp.ProtoMinor = 1
	nresp.Header.Add("Cache-Control", "max-age=0, no-cache, no-store")
	nresp.Header.Add("Connection", "keep-alive")
	nresp.Header.Add("Date", finalDate)
	nresp.Header.Add("Server", "nginx/1.0.12")
	nresp.Header.Add("Vary", "Accept-Encoding")
	return *nresp
}
