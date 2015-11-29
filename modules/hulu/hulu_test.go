package hulu

import (
	"fmt"
	"net/http"
	"testing"
)

func TestLoadFakeResponse(t *testing.T) {
	fmt.Println(LoadFakeResponse())
}

func TestLoadFakeResponseNoCommercial(t *testing.T) {
	fmt.Println(LoadCommericalFreeResponse())
}

func TestCreateResponse(t *testing.T) {
	r := CreateResponse(&http.Request{})
	fmt.Println(r)
}
