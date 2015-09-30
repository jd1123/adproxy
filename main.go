package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/elazarl/goproxy"
	"github.com/jd1123/adproxy/modules"
)

var mods = make([]modules.Module, 0)

func main() {
	RegisterModule(modules.Xfinity{})

	fmt.Println("Starting ad proxy on port " + LISTENPORT)
	proxy := goproxy.NewProxyHttpServer()
	proxy.Verbose = false
	f, err := os.OpenFile("/etc/adproxy/log/proxylog.log", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0660)
	if err != nil {
		fmt.Println("File error: ", err)
		os.Exit(1)
	}
	defer f.Close()

	proxy.NonproxyHandler = http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if req.Host == "" {
			fmt.Fprintln(w, "Cannot handle request without Host header, e.g., HTTP 1.0")
			return
		}
		req.URL.Scheme = "http"
		req.URL.Host = req.Host
		proxy.ServeHTTP(w, req)
	})

	log.SetOutput(f)
	// Setup modules

	if FILTER {
		proxy.OnRequest().DoFunc(filterRequest)
		proxy.OnResponse().DoFunc(filterResponse)
	}
	log.Fatalln(http.ListenAndServe(":"+LISTENPORT, proxy))
	fmt.Println("Closing ad proxy")

}
