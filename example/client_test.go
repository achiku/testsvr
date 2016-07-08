package client

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/achiku/testserver"
)

func TestNewClient(t *testing.T) {
	c := NewClient()
	if c.c == nil {
		t.Error("failed to create client")
	}
}

func TestClientHello(t *testing.T) {
	c := NewClient()
	s := httptest.NewServer(NewMockServerMux(nil, t))
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
	hm := testserver.URLHandlerMap{
		"/hello": func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "failed")
		},
	}
	s := httptest.NewServer(NewMockServerMux(hm, t))
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
