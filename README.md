# testsvr

[![Build Status](https://travis-ci.org/achiku/testsvr.svg?branch=master)](https://travis-ci.org/achiku/testsvr)
[![GitHub license](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/achiku/testsvr/master/LICENSE)
[![Doc](https://img.shields.io/badge/godoc-reference-blue.svg)](https://godoc.org/github.com/achiku/testsvr)
[![Go Report Card](https://goreportcard.com/badge/github.com/achiku/testsvr)](https://goreportcard.com/report/github.com/achiku/testsvr)

## Description

Make `httptest` generated test server aware of `-v` of `go test`, and reuse that code for stand alone mock server


## Why created

Interacting with external services becomes almost an requirement for modern Web service development. Golang has an awesome `httptest` standard library making it really easy to start and close mock server in test code. However,  this mock server, which is supposed to be a part of test code, is not aware of `-v` of `go test`, and this leads to verbose output if you put logging in mock server handlers. This tiny little library makes `httptest` generated mock server aware of `-v`, and keep your test output sane and clean.


## Installation

```
go get -u github.com/achiku/testsvr
```


## Example

```go
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
```

```go

func TestClientHello(t *testing.T) {
	c := NewClient()
	s := httptest.NewServer(testsvr.NewMux(DefaultHandlerMap, t))
	defer s.Close()

	name := "achiku"
	status, resp, err := c.Hello(s.URL, name)
	if err != nil {
		t.Fatal(err)
	}
	if status != http.StatusOK {
		t.Errorf("want %d got %d", http.StatusOK, status)
	}
	if resp != "hello!" {
		t.Errorf("want hello! got %s", resp)
	}
}
```
