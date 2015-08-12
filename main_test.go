package main

import (
	"fmt"
	"net/http"
	"testing"
	"time"
)

func TestNewResponse(t *testing.T) {
	req, _ := http.NewRequest("GET", "http://127.0.0.1", nil)
	nresp := CreateResponse(req)
	fmt.Println(nresp)
}

func TestFormatTime(t *testing.T) {
	tm := time.Now()
	formatTime(tm)
}
