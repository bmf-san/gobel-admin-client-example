package api

import (
	"net/http"
	"net/url"
	"os"
)

// GetPosts requests posts
func (a *Client) GetPosts(values url.Values, token string) (*http.Response, []byte, error) {
	req, err := http.NewRequest(http.MethodGet, os.Getenv("API_URL")+"/private/posts?", nil)
	if err != nil {
		return nil, nil, err
	}
	req.URL.RawQuery = values.Encode()
	req.Header.Set("Authorization", token)

	resp, body, err := a.HandleResponseBody(req)
	if err != nil {
		return nil, nil, err
	}

	return resp, body, err
}
