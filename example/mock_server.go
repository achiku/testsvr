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

func hello(logger testsvr.Logfer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger.Logf("yeaaaaaaaahhh!!")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "hello!")
	}
}

func goodbye(logger testsvr.Logfer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger.Logf("goodby!!!!!")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "goodbye!")
	}
}
