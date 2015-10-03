package hulu

var FILTER = true

var FILTER_STRINGS = LoadFilterList()

//var FILTER_STRINGS = [...]string{"adserver", "insightexpressai.com", "imrworldwide.com", "doubleverify.com", "scorecardresearch.com", "ads", "rewardtv.com", "flurry.com", "doubleclick"}

var FILTER_STRINGS_XFINITY = [...]string{"adserver"}
var ADSTATE_FAKE = "adstate=RAYMONDW3NzdCxdW2JoLFRoZSBXcm9uZyBNYW5zO0h1bHVdW2l4aCxUaGUgV3JvbmcgTWFuc3xIdWx1XVtmY2gsQVFBQVFCQUFBMjhEQVFBQUF6ZWpBUUVBQTROQUFRRUFBajBwQkFBQkFRRUFBNE5QQXdFQkFRQURpbTBCQUFBRFFVd0NBUUVBQTZXU0FnRUJBQU5VOFFJQUFBQURWL1VCQVFBRE43b0NBUUVBQTV6YUFRRUFBNVBhQWdFQkFBT01Yd0VCQUFFQlhRSUJBUT09XVtmd3VpZCwzYmZhYTI2My0xNjkxLTQ2ZmUtYjk3OS04MTNhZWU1N2ZhMTRdW2ZpeGgsfHx8fHx8fHx8fHx8fHx8fHx8fHx8XVtmd2VjLGVKd0RBQUFBQUFFPV0&force_should_resume=0&guuid=000003a79fb9804111424ed13b1780e20b33&kids_only=0&kv=399578&version=441000"
var FAKERESPONSE = LoadFakeResponse()
var ContentTypeJSON = "application/json; charset=utf-8"
