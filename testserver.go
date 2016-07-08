package testserver

import (
	"log"
	"net/http"
)

// URLHandlerMap url and handler
type URLHandlerMap map[string]http.HandlerFunc

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
type NewMux func(URLHandlerMap, Logfer) *http.ServeMux

// NewServer creates dev server
type NewServer func(port string) *http.Server

// CreateHandler creates handlers for test/dev server
type CreateHandler func(Logfer) http.HandlerFunc
