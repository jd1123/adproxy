package stopper

import (
	"io"
	"net/http"
	"testing"
)

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello World!")
}

func testStopper(t *testing.T) {
	http.HandleFunc("/", hello)
	StoppableListenAndServe(":9999", nil)

}
