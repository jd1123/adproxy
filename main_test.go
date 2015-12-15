package main

import (
	"bytes"
	"fmt"
	"net"
	"net/http"
	"os"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/elazarl/goproxy"
	"github.com/hydrogen18/stoppableListener"
	"github.com/jd1123/adproxy/modules"
)

type TestModule struct {
	Metadata      modules.MetaStruct
	FilterStrings []string
}

func (tm TestModule) Init() {
	tm.Metadata.ModuleName = "Test Filter"
	tm.Metadata.VersionNumber = "0.0"
	tm.Metadata.Service = "Testing"
	tm.FilterStrings = append(tm.FilterStrings, "google.com")
}

func (tm TestModule) FilterResponse(resp *http.Response, ctx *goproxy.ProxyCtx) *http.Response {
	for _, i := range tm.FilterStrings {
		if strings.Contains(resp.Request.URL.String(), i) {
			fmt.Println("Adserver found... blocking: ", resp.Request.URL.String())
			bb := modules.ClosingBuffer{bytes.NewBufferString("0")}
			resp.Body = bb
		}
	}
	return resp
}

func NewTestModule() *TestModule {
	tm := TestModule{}
	tm.Init()
	return &tm
}

func (tm TestModule) FilterRequest(req *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {
	// Block analytics requests
	if strings.Contains(req.URL.String(), "analytics.xcal.tv") {
		fmt.Println(req.URL.String(), "Analytics Request Intercepted...")
		return req, goproxy.NewResponse(req, goproxy.ContentTypeText, http.StatusOK, "0")
	}
	return req, nil
}

func ListenAndServeStop(addr string, handler http.Handler, wg *sync.WaitGroup) *stoppableListener.StoppableListener {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	sl, err := stoppableListener.New(listener)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	server := &http.Server{Addr: addr, Handler: handler}
	select {
	case <-stoppableListener.stop:
		fmt.Println("Stopped")
	}

	return sl
}

// Test Harness for testing the filters
func setup(wg *sync.WaitGroup) *stoppableListener.StoppableListener {
	os.Setenv("HTTP_PROXY", "http://127.0.0.1:9999")
	RegisterModule(NewTestModule())
	proxy := goproxy.NewProxyHttpServer()
	proxy.NonproxyHandler = http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if req.Host == "" {
			fmt.Fprintln(w, "Cannot handle request without Host header, e.g., HTTP 1.0")
			return
		}
		req.URL.Scheme = "http"
		req.URL.Host = req.Host
		proxy.ServeHTTP(w, req)
	})
	proxy.OnRequest().DoFunc(filterRequest)
	proxy.OnResponse().DoFunc(filterResponse)
	//	listener, err := net.Listen("tcp", ":9999")
	//	if err != nil {
	//		fmt.Println("Error:", err)
	//		os.Exit(1)
	//	}
	//	sl, err := stoppableListener.New(listener)
	//	if err != nil {
	//		fmt.Println("Error:", err)
	//		os.Exit(1)
	//	}
	//	server := http.Server{Addr: ":9999", Handler: proxy}

	fmt.Println("Serving HTTP for testing...")
	//	go func() {
	//		server.Serve(sl)
	//	}()
	sl := ListenAndServeStop(":9999", proxy, wg)
	//if err != nil {
	//	fmt.Println("Error: ", err)
	//	os.Exit(1)
	//}
	return sl
}

func tearDown(sl *stoppableListener.StoppableListener) {
	sl.Stop()
}

func TestFilterResponse(t *testing.T) {
	var wg sync.WaitGroup
	setup(&wg)
	fmt.Println("Getting")
	//r, _ := http.Get("http://filmgarb.com")
	fmt.Println("Got")
	wg.Wait()
	//defer r.Body.Close()
	//fmt.Println(r)
	//body, _ := ioutil.ReadAll(r.Body)
	//fmt.Println(body)
	//tearDown(sl)
}

/*
func TestNewResponse(t *testing.T) {
	req, _ := http.NewRequest("GET", "http://127.0.0.1", nil)
	nresp := CreateResponse(req)
	fmt.Println(nresp)
}
*/

func TestFormatTime(t *testing.T) {
	tm := time.Now()
	formatTime(tm)
}
