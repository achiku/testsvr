# testsvr

[![GitHub license](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/achiku/testsvr/master/LICENSE)
[![Doc](https://img.shields.io/badge/godoc-reference-blue.svg)](https://godoc.org/github.com/achiku/testsvr)

## Description

Make test server aware of `-v` of `go test`, and reuse that code for stand alone mock server


## Why created

Interacting with external services becomes almost an requirement for modern Web service development. This simple set of functions makes it easier to share codes between mock server for development and tests. Also, using this library makes your test server created by `httptest` aware of `-v` of `go test` so that you can switch on and off request/response logging easily.


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
