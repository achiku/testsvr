package testsvr

import (
	"log"
	"net/http"
)

// CreateHandler creates handlers for test/dev server
type CreateHandler func(Logfer) http.HandlerFunc

// URLHandlerMap url and handler
type URLHandlerMap map[string]CreateHandler

// Logfer inerface Logf
type Logfer interface {
	Logf(format string, args ...interface{})
}

// Logger is used in NewServer for logging
type Logger struct{}

// Logf output log
func (l Logger) Logf(format string, args ...interface{}) {
	log.Println(format, args)
}

// NewMux creates mux for test/dev server
func NewMux(hm URLHandlerMap, l Logfer) *http.ServeMux {
	mux := http.NewServeMux()
	for url, handler := range hm {
		mux.HandleFunc(url, handler(l))
	}
	return mux
}

// NewServer creates dev server
func NewServer(hm URLHandlerMap, port string) *http.Server {
	mux := NewMux(hm, Logger{})
	server := &http.Server{
		Handler: mux,
		Addr:    "localhost:" + port,
	}
	return server
}
