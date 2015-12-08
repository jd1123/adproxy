package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/elazarl/goproxy"
	"github.com/jd1123/adproxy/config"
	"github.com/jd1123/adproxy/modules"
	"github.com/jd1123/adproxy/modules/hulu"
	"github.com/jd1123/adproxy/modules/xfinity"
)

var mods = make([]modules.Module, 0)

func main() {
	// Load Configuration
	configSettings := config.DefaultConfig()

	// Load the modules you want to use
	RegisterModule(xfinity.NewXfinity())
	RegisterModule(hulu.NewHulu(configSettings.Quiet))

	// Begin
	fmt.Println("Starting ad proxy on port " + configSettings.ListenPort)
	proxy := goproxy.NewProxyHttpServer()
	proxy.Verbose = configSettings.ProxyVerbose
	f, err := os.OpenFile(configSettings.LogFile, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0660)
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

	// Set the output of the log function
	// ad the logfile
	log.SetOutput(f)

	// Set it up so on requests and responses, it
	// uses the filters
	proxy.OnRequest().DoFunc(filterRequest)
	proxy.OnResponse().DoFunc(filterResponse)

	// Start her up
	log.Fatalln(http.ListenAndServe(":"+LISTENPORT, proxy))
	fmt.Println("Closing ad proxy")

}
