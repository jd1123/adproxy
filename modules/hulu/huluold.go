package modules

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/elazarl/goproxy"
)

var FILTER = true

var FILTER_STRINGS = LoadFilterList()

//var FILTER_STRINGS = [...]string{"adserver", "insightexpressai.com", "imrworldwide.com", "doubleverify.com", "scorecardresearch.com", "ads", "rewardtv.com", "flurry.com", "doubleclick"}

var FILTER_STRINGS_XFINITY = [...]string{"adserver"}
var ADSTATE_FAKE = "adstate=RAYMONDW3NzdCxdW2JoLFRoZSBXcm9uZyBNYW5zO0h1bHVdW2l4aCxUaGUgV3JvbmcgTWFuc3xIdWx1XVtmY2gsQVFBQVFCQUFBMjhEQVFBQUF6ZWpBUUVBQTROQUFRRUFBajBwQkFBQkFRRUFBNE5QQXdFQkFRQURpbTBCQUFBRFFVd0NBUUVBQTZXU0FnRUJBQU5VOFFJQUFBQURWL1VCQVFBRE43b0NBUUVBQTV6YUFRRUFBNVBhQWdFQkFBT01Yd0VCQUFFQlhRSUJBUT09XVtmd3VpZCwzYmZhYTI2My0xNjkxLTQ2ZmUtYjk3OS04MTNhZWU1N2ZhMTRdW2ZpeGgsfHx8fHx8fHx8fHx8fHx8fHx8fHx8XVtmd2VjLGVKd0RBQUFBQUFFPV0&force_should_resume=0&guuid=000003a79fb9804111424ed13b1780e20b33&kids_only=0&kv=399578&version=441000"
var FAKERESPONSE = LoadFakeResponse()
var ContentTypeJSON = "application/json; charset=utf-8"

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
