package hulu

import (
	"net/http"

	"github.com/elazarl/goproxy"
)

var filterStrings = []string{"insightexpressai.com", "imrworldwide.com", "doubleverify.com", "scorecardresearch.com",
	"ads", "rewardtv.com", "flurry.com", "doubleclick"}

type Hulu struct {
	Metadata MetaStruct
}

func (u Hulu) FilterResponse(resp *http.Response, ctx *goproxy.ProxyCtx) *http.Response {
}

func (x Xfinity) FilterRequest(req *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {
}
