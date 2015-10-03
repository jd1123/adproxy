package main

import (
	"fmt"
	"log"
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

	log.Println("Req: ", req.Method, ": ", req.URL.String())

	return req, nil
	return req, resp
}

// Response filter function
func filterResponse(resp *http.Response, ctx *goproxy.ProxyCtx) *http.Response {
	for i := range mods {
		resp = mods[i].FilterResponse(resp, ctx)
	}
	return resp
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
