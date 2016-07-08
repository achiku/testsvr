package example

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// Client struct
type Client struct {
	c *http.Client
}

// NewClient returns Client
func NewClient() *Client {
	return &Client{c: &http.Client{}}
}

// Hello accesses /hello and return the response body in the string type
func (c *Client) Hello(host, name string) (int, string, error) {
	req, err := http.NewRequest("GET", host+"/hello", nil)
	if err != nil {
		return 0, "", fmt.Errorf("failed to create request: %s", err.Error())
	}
	resp, err := c.c.Do(req)
	if err != nil {
		return 0, "", fmt.Errorf("failed to send request: %s", err.Error())
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, "", fmt.Errorf("failed to read body: %s", err.Error())
	}
	return resp.StatusCode, string(body), nil
}

// Goodbye accesses /goodbye and return the response body in the string type
func (c *Client) Goodbye(host, name string) (int, string, error) {
	return 0, "", nil
}
