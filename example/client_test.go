package example

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/achiku/testsvr"
)

func TestNewClient(t *testing.T) {
	c := NewClient()
	if c.c == nil {
		t.Error("failed to create client")
	}
}

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

func TestClientHelloError(t *testing.T) {
	c := NewClient()
	hl := func(logger testsvr.Logfer) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
			logger.Logf("something went wrong")
			fmt.Fprintf(w, "failed")
		}
	}
	by := func(logger testsvr.Logfer) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
			logger.Logf("something went wrong")
			fmt.Fprintf(w, "failed")
		}
	}
	hm := testsvr.URLHandlerMap{
		"/hello":   hl,
		"/goodbye": by,
	}
	s := httptest.NewServer(testsvr.NewMux(hm, t))
	defer s.Close()

	name := "achiku"
	status, resp, err := c.Hello(s.URL, name)
	if err != nil {
		t.Fatal(err)
	}
	if status != http.StatusInternalServerError {
		t.Errorf("want %d got %d", http.StatusOK, status)
	}
	if resp != "failed" {
		t.Errorf("want failed got %s", resp)
	}
}
