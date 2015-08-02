package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/elazarl/goproxy"
)

func main() {
	fmt.Println("Starting ad proxy on port " + LISTENPORT)
	proxy := goproxy.NewProxyHttpServer()
	proxy.Verbose = false

	f, err := os.OpenFile("log/proxylog.log", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0660)
	if err != nil {
		fmt.Println("File error: ", err)
		os.Exit(1)
	}
	defer f.Close()

	log.SetOutput(f)

	if FILTER {
		proxy.OnRequest().DoFunc(filterRequest)
		proxy.OnResponse().DoFunc(filterResponse)
	}
	log.Fatalln(http.ListenAndServe(":"+LISTENPORT, proxy))
	fmt.Println("Closing ad proxy")
}
