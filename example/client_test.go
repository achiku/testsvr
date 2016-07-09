package example

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/achiku/testsvr"
)

func TestNewClient(t *testing.T) {
	c := NewClient("http://localhost:8080")
	if c.c == nil {
		t.Error("failed to create client")
	}
}

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

func TestClientHelloError(t *testing.T) {
	hl := func(logger testsvr.Loger) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
			logger.Logf("something went wrong")
			fmt.Fprintf(w, "failed")
		}
	}
	by := func(logger testsvr.Loger) http.HandlerFunc {
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

	c := NewClient(s.URL)
	name := "achiku"
	status, resp, err := c.Hello(name)
	if err != nil {
		t.Fatal(err)
	}
	if status != http.StatusInternalServerError {
		t.Errorf("want %d got %d", http.StatusOK, status)
	}
	if resp != "failed" {
		t.Errorf("want failed got %s", resp)
	}

	status, resp, err = c.Goodbye(name)
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
