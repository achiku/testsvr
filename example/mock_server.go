package example

import (
	"fmt"
	"net/http"

	"github.com/achiku/testsvr"
)

// DefaultHandlerMap default url and handler map
var DefaultHandlerMap = map[string]testsvr.CreateHandler{
	"/hello":   hello,
	"/goodbye": goodbye,
}

func hello(logger testsvr.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger.Logf("RawURL: %s", r.URL)
		logger.Logf("Header: %s", r.Header)
		name := r.URL.Query().Get("name")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "hello! %s.", name)
	}
}

func goodbye(logger testsvr.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger.Logf("RawURL: %s", r.URL)
		logger.Logf("Header: %s", r.Header)
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "goodbye!")
	}
}
