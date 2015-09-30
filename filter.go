package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/elazarl/goproxy"
)

/* 	This function inspects the body of a ReadCloser and
returns a string with its contents, then restores
the ReadCloser to its original state and returns its
contents as a string. Passes ReadCloser by reference
*/
// Request filter function
func filterRequest(req *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {
	var resp *http.Response
	for i := range mods {
		req, resp = mods[i].FilterRequest(req, ctx)
	}
	/*
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
			if VERB {
				fmt.Println(req.URL.String(), "Analytics Request Intercepted...")
			}
			return req, goproxy.NewResponse(req, goproxy.ContentTypeText, http.StatusOK, "0")
		}

		b := checkBody(&req.Body)

		// Flag for response spoofing
		g := false
		if strings.Contains(b, "adstate") && g {
			//return req, goproxy.NewResponse(req, ContentTypeJSON, http.StatusOK, FAKERESPONSE)
			nresp := CreateResponse(req)
			return req, &nresp

			//fmt.Println(b)
		}
	*/
	return req, resp
}

// Response filter function
func filterResponse(resp *http.Response, ctx *goproxy.ProxyCtx) *http.Response {
	for i := range mods {
		resp = mods[i].FilterResponse(resp, ctx)
	}
	return resp
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

func formatTime(t time.Time) string {
	h := strconv.Itoa(t.Hour())
	m := strconv.Itoa(t.Minute())
	s := strconv.Itoa(t.Second())
	d := strconv.Itoa(t.Day())
	y := strconv.Itoa(t.Year())
	mnth := t.Month().String()[:3]
	wd := t.Weekday().String()[:3]
	if len(s) == 1 {
		s = "0" + s
	}
	if len(d) == 1 {
		d = "0" + d
	}

	fmt.Println(h, m, s, d, y, wd)
	str := wd + ", " + d + " " + mnth + " " + y + " " + h + ":" + m + ":" + s
	fmt.Println(str)
	return ""
}
