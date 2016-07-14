package testsvr

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

// DefaultHandlerMap default url and handler map
var DefaultHandlerMap = map[string]CreateHandler{
	"/hello":   hello,
	"/goodbye": goodbye,
}

func hello(logger Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger.Logf("RawURL: %s", r.URL)
		logger.Logf("Header: %s", r.Header)
		name := r.URL.Query().Get("name")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "hello! %s.", name)
	}
}

func goodbye(logger Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger.Logf("RawURL: %s", r.URL)
		logger.Logf("Header: %s", r.Header)
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "goodbye!")
	}
}

func TestNewMux(t *testing.T) {
	s := httptest.NewServer(NewMux(DefaultHandlerMap, t))
	defer s.Close()

	testData := []struct {
		name   string
		status int
	}{
		{"moqada", http.StatusOK},
		{"8maki", http.StatusOK},
		{"achiku", http.StatusOK},
	}

	for _, d := range testData {
		v := url.Values{}
		v.Add("name", d.name)
		c := &http.Client{}
		res, err := c.Get(s.URL + "/hello?" + v.Encode())
		if err != nil {
			t.Fatal(err)
		}
		if res.StatusCode != d.status {
			t.Errorf("want %d got %d", d.status, res.StatusCode)
		}
	}
}
