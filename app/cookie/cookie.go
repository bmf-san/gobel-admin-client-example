package cookie

import (
	"encoding/base64"
	"net/http"
	"os"
	"time"
)

// Cookie represents the singular of cookie.
type Cookie struct {
	Flash *Flash
}

// NewCookie creates a Cookie.
func NewCookie() *Cookie {
	return &Cookie{
		Flash: &Flash{
			Messages: make(map[string]string),
		},
	}
}

// A Flash is a flash messages.
type Flash struct {
	Messages map[string]string
}

// SetFlash sets a flash message.
func (c *Cookie) SetFlash(w http.ResponseWriter, key string, value []byte) {
	c.Flash.Messages[key] = string(value)

	http.SetCookie(w, &http.Cookie{
		Name:  key,
		Value: base64.RawURLEncoding.EncodeToString(value),
	})
}

// GetFlash gets a flash message from cookie.
func (c *Cookie) GetFlash(w http.ResponseWriter, r *http.Request) (*Flash, error) {
	for key := range c.Flash.Messages {
		_, err := r.Cookie(key)
		if err != nil {
			switch err {
			case http.ErrNoCookie:
				return nil, nil
			default:
				return nil, err
			}
		}

		http.SetCookie(w, &http.Cookie{
			Name:    key,
			MaxAge:  -1,
			Expires: time.Unix(1, 0),
		})
	}

	return c.Flash, nil
}

// SetSession sets a login session.
func (c *Cookie) SetSession(w http.ResponseWriter, r *http.Request, body []byte) {
	http.SetCookie(w, &http.Cookie{
		Name:    os.Getenv("SESSION_NAME"),
		Value:   string(body),
		Expires: time.Now().Add(3600 * 24 * 7 * time.Second),
	})
}

// GetSession gets a session.
func (c *Cookie) GetSession(w http.ResponseWriter, r *http.Request) (string, error) {
	token, err := r.Cookie(os.Getenv("SESSION_NAME"))
	if err != nil {
		return "", err
	}

	return token.Value, nil
}

// SetLoginSession sets a login session.
func (c *Cookie) SetLoginSession(w http.ResponseWriter, r *http.Request, body []byte) {
	http.SetCookie(w, &http.Cookie{
		Name:    os.Getenv("LOGIN_SESSION_NAME"),
		Value:   string(body),
		Expires: time.Now().Add(3600 * 24 * 7 * time.Second),
	})
}

// GetLoginSession gets a login session.
func (c *Cookie) GetLoginSession(w http.ResponseWriter, r *http.Request) (string, error) {
	token, err := r.Cookie(os.Getenv("LOGIN_SESSION_NAME"))
	if err != nil {
		return "", err
	}

	return token.Value, nil
}

// DelLoginSession deletes a login session.
func (c *Cookie) DelLoginSession(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:   os.Getenv("LOGIN_SESSION_NAME"),
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	})
}
