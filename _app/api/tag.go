package api

import (
	"net/http"
	"os"
)

// GetTags requests tags.
func (a *Client) GetTags(token string) (*http.Response, []byte, error) {
	req, err := http.NewRequest(http.MethodGet, os.Getenv("API_URL")+"/private/tags?", nil)
	if err != nil {
		return nil, nil, err
	}
	req.Header.Set("Authorization", token)

	resp, body, err := a.HandleResponseBody(req)
	if err != nil {
		return nil, nil, err
	}

	return resp, body, err
}
