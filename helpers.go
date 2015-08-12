package main

import (
	"fmt"
	"net/http"
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
