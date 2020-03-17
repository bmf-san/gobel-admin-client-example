package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"
)

// A SignIn represents the singular of credential for sign in.
type SignIn struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// PostSignIn requests signin.
func (a *Client) PostSignIn(s *SignIn) (*http.Response, []byte, error) {
	reqBody, err := json.Marshal(s)
	if err != nil {
		return nil, nil, err
	}

	req, err := http.NewRequest(http.MethodPost, os.Getenv("API_URL")+"/signin", bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, nil, err
	}

	resp, body, err := a.HandleResponseBody(req)
	if err != nil {
		return nil, nil, err
	}

	return resp, body, err
}

// GetSignOut requests signout.
func (a *Client) GetSignOut(token string) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodPost, os.Getenv("API_URL")+"/private/signout", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", token)

	resp, err := a.HandleResponse(req)
	if err != nil {
		return nil, err
	}

	return resp, err
}
