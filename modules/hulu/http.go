package hulu

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/elazarl/goproxy"
)

func PrintResponse(r http.Response) {
	fmt.Println("Status:", r.Status)
	fmt.Println("Status Code:", r.StatusCode)
	fmt.Println("Proto:", r.Proto)
	fmt.Println("Proto Major:", r.ProtoMajor)
	fmt.Println("Proto Minor:", r.ProtoMinor)
	fmt.Println("Header:", r.Header)
	fmt.Println("Body:", checkBody(&r.Body))
	fmt.Println("Content Length:", r.ContentLength)
	fmt.Println("Transfer Encoding:", r.TransferEncoding)
	fmt.Println("Close:", r.Close)
	fmt.Println("Trailer:", r.Trailer)
}

/* 	This function inspects the body of a ReadCloser and
returns a string with its contents, then restores
the ReadCloser to its original state and returns its
contents as a string. Passes ReadCloser by reference
*/
func checkBody(rc *io.ReadCloser) string {
	var buf []byte
	buf, _ = ioutil.ReadAll(*rc)
	b := string(buf)
	*rc = ioutil.NopCloser(bytes.NewBuffer(buf))
	return b
}

func CreateResponse(req *http.Request) http.Response {
	n := time.Now()
	layout := "Sun, 09 Aug 1999 18:20:22 GMT"
	finalDate := n.Format(layout)
	respBody := LoadFakeResponse()
	nresp := goproxy.NewResponse(req, ContentTypeJSON, http.StatusOK, respBody)
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
