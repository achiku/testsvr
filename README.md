# testsvr

[![Build Status](https://travis-ci.org/achiku/testsvr.svg?branch=master)](https://travis-ci.org/achiku/testsvr)
[![GitHub license](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/achiku/testsvr/master/LICENSE)
[![Doc](https://img.shields.io/badge/godoc-reference-blue.svg)](https://godoc.org/github.com/achiku/testsvr)
[![Go Report Card](https://goreportcard.com/badge/github.com/achiku/testsvr)](https://goreportcard.com/report/github.com/achiku/testsvr)

## Description

Make `httptest` generated test server aware of `-v` of `go test`, and reuse that code for stand alone mock server


## Why created

Interacting with external services becomes almost an requirement for modern Web service development. Golang has an awesome `httptest` standard library making it really easy to start and close mock server in test code. However,  this mock server, which is supposed to be a part of test code, is not aware of `-v` of `go test`, and this leads to verbose output if you put logging in mock server handlers. This tiny little library makes `httptest` generated mock server aware of `-v`, keeping your test output sane and clean by default, detailed and comprehensive if needed.


## Installation

```
go get -u github.com/achiku/testsvr
```


## Example

Full example is in https://github.com/achiku/testsvr/tree/master/example

###### mock server code

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
```

###### client test code

```go
func TestClientHello(t *testing.T) {
	s := httptest.NewServer(testsvr.NewMux(DefaultHandlerMap, t))
	defer s.Close()

	testData := []struct {
		name   string
		status int
	}{
		{"moqada", http.StatusOK},
		{"8maki", http.StatusOK},
		{"achiku", http.StatusOK},
	}
	c := NewClient(s.URL)
	for _, d := range testData {
		status, resp, err := c.Hello(d.name)
		if err != nil {
			t.Fatal(err)
		}
		if status != d.status {
			t.Errorf("want %d got %d", http.StatusOK, status)
		}
		expResp := fmt.Sprintf("hello! %s.", d.name)
		if resp != expResp {
			t.Errorf("want %s got %s", expResp, resp)
		}
	}
}
```

When you don't want verbose output.

```
$ go test
PASS
ok      github.com/achiku/testsvr/example       0.016s
```


When you want detailed output.

```
$ go test -v
=== RUN   TestNewClient
--- PASS: TestNewClient (0.00s)
=== RUN   TestClientHello
--- PASS: TestClientHello (0.00s)
        mock_server.go:18: RawURL: /hello?name=moqada
        mock_server.go:19: Header: map[User-Agent:[Go-http-client/1.1] Accept-Encoding:[gzip]]
        mock_server.go:18: RawURL: /hello?name=8maki
        mock_server.go:19: Header: map[User-Agent:[Go-http-client/1.1] Accept-Encoding:[gzip]]
        mock_server.go:18: RawURL: /hello?name=achiku
        mock_server.go:19: Header: map[User-Agent:[Go-http-client/1.1] Accept-Encoding:[gzip]]
=== RUN   TestClientHelloError
--- PASS: TestClientHelloError (0.00s)
        client_test.go:51: something went wrong
        client_test.go:58: something went wrong
PASS
ok      github.com/achiku/testsvr/example       0.018s
```

