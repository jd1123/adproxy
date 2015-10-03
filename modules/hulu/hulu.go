package hulu

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/elazarl/goproxy"
	"github.com/jd1123/adproxy/modules"
)

// URI strings to filter
var filterStrings = []string{"insightexpressai.com", "imrworldwide.com", "doubleverify.com", "scorecardresearch.com",
	"ads", "rewardtv.com", "flurry.com", "doubleclick"}

type Hulu struct {
	Metadata modules.MetaStruct
}

func NewHulu() *Hulu {
	ms := modules.NewMetaStruct("Hulu Filter", "0.1", "Hulu")
	return &Hulu{*ms}
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
			fmt.Println("Adserver found in", h.Metadata.ModuleName, "...blocking:", resp.Request.URL.String())
			bb := ClosingBuffer{bytes.NewBufferString("0")}
			resp.Body = bb
		}
	}
	return resp
}

func (h Hulu) FilterRequest(req *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {
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

// struct to implement io.ReadCloser
type ClosingBuffer struct {
	io.Reader
}

func (cb ClosingBuffer) Close() (err error) {
	return
}
