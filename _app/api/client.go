package api

import (
	"io/ioutil"
	"net/http"
	"time"
)

// Client represents the singular of api client.
type Client struct {
	Client *http.Client
}

// NewClient creates a Client.
func NewClient() *Client {
	return &Client{
		Client: &http.Client{
			Timeout: time.Duration(100 * time.Second),
		},
	}
}

// HandleResponseBody handles an api client for getting response and body.
func (a *Client) HandleResponseBody(req *http.Request) (*http.Response, []byte, error) {
	resp, err := a.Client.Do(req)
	if err != nil {
		return nil, nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, err
	}

	return resp, body, nil
}

// HandleResponse handles an api client for getting response.
func (a *Client) HandleResponse(req *http.Request) (*http.Response, error) {
	resp, err := a.Client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	return resp, nil
}
