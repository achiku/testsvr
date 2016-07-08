package client

import (
	"fmt"
	"net/http"

	"github.com/achiku/testserver"
)

func hello(logger testserver.Logfer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger.Logf("yeaaaaaaaahhh!!")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "hello!")
	}
}

func goodbye(logger testserver.Logfer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger.Logf("goodby!!!!!")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "goodbye!")
	}
}

// NewMockServerMux creates new mock server mux
func NewMockServerMux(hm testserver.URLHandlerMap, logger testserver.Logfer) *http.ServeMux {
	mux := http.NewServeMux()
	if hm == nil {
		mux.HandleFunc("/hello", hello(logger))
		mux.HandleFunc("/goodbye", goodbye(logger))
		return mux
	}

	for url, handler := range hm {
		mux.HandleFunc(url, handler)
	}
	return mux
}

// NewMockServer creates new mock server
func NewMockServer(port string) *http.Server {
	logger := testserver.Logger{}
	mux := NewMockServerMux(nil, logger)
	server := &http.Server{
		Handler: mux,
		Addr:    "localhost:" + port,
	}
	return server
}
