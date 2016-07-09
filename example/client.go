package example

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// Client struct
type Client struct {
	c    *http.Client
	host string
}

// NewClient returns Client
func NewClient(host string) *Client {
	return &Client{c: &http.Client{}, host: host}
}

// Hello accesses /hello and return the response body in the string type
func (c *Client) Hello(name string) (int, string, error) {
	v := url.Values{}
	v.Add("name", name)
	req, err := http.NewRequest("GET", c.host+"/hello?"+v.Encode(), nil)
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
func (c *Client) Goodbye(name string) (int, string, error) {
	v := url.Values{}
	v.Add("name", name)
	req, err := http.NewRequest("GET", c.host+"/goodbye?"+v.Encode(), nil)
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
