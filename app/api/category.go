package api

import (
	"net/http"
	"os"
)

// GetCategories requests categories.
func (a *Client) GetCategories(token string) (*http.Response, []byte, error) {
	req, err := http.NewRequest(http.MethodGet, os.Getenv("API_URL")+"/private/categories?", nil)
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
