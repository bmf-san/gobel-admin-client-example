package controller

import (
	"net/http"

	"github.com/bmf-san/gobel-admin-client-example/app/api"
	"github.com/bmf-san/gobel-admin-client-example/app/cookie"
	"github.com/bmf-san/gobel-admin-client-example/app/logger"
	"github.com/bmf-san/gobel-admin-client-example/app/response"
	"github.com/bmf-san/gobel-admin-client-example/app/session"
)

// An AuthController is a controller for a tag.
type AuthController struct {
	Logger    *logger.Logger
	Client    *api.Client
	ConnRedis *session.RedisHandler
	Response  *response.Response
	Cookie    *cookie.Cookie
}

// NewAuthController creates a AuthController.
func NewAuthController(logger *logger.Logger, apiClient *api.Client, connRedis *session.RedisHandler, response *response.Response, cookie *cookie.Cookie) *AuthController {
	return &AuthController{
		Logger:    logger,
		Client:    apiClient,
		ConnRedis: connRedis,
		Response:  response,
		Cookie:    cookie,
	}
}

// Index displays a sign in view.
func (ac *AuthController) Index(w http.ResponseWriter, r *http.Request) {
	flash, err := ac.Cookie.GetFlash(w, r)
	if err != nil {
		ac.Logger.Error(err.Error())
		ac.Response.Error(w, http.StatusInternalServerError)
		return
	}

	sessID, ok := r.Context().Value("session-id").(string)
	if !ok {
		ac.Logger.Error(err.Error())
		ac.Response.Error(w, http.StatusInternalServerError)
		return
	}

	token, ok := r.Context().Value("csrf-token").(string)
	if !ok {
		ac.Logger.Error(err.Error())
		ac.Response.Error(w, http.StatusInternalServerError)
		return
	}

	old, err := ac.ConnRedis.GetInputOld(sessID)
	if err != nil {
		ac.Logger.Error(err.Error())
		ac.Response.Error(w, http.StatusInternalServerError)
		return
	}

	if err := ac.Response.ExecuteAuthIndex(w, &response.AuthIndex{
		Flash:     flash,
		CSRFToken: token,
		Old:       old,
	}); err != nil {
		ac.Logger.Error(err.Error())
		ac.Response.Error(w, http.StatusInternalServerError)
	}
}

// SignIn signs in with a credential.
func (ac *AuthController) SignIn(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	email := r.Form.Get("email")
	password := r.Form.Get("password")

	sessID, err := ac.Cookie.GetSession(w, r)
	if err != nil {
		ac.Logger.Error(err.Error())
		ac.Response.Error(w, http.StatusInternalServerError)
		return
	}

	ac.ConnRedis.SetInputOld(r.Form, sessID, "csrf_token", "password")

	resp, body, err := ac.Client.PostSignIn(&api.SignIn{
		Email:    email,
		Password: password,
	})
	if err != nil {
		ac.Logger.Error(err.Error())
		ac.Response.Error(w, http.StatusInternalServerError)
		return
	}

	if resp.StatusCode == http.StatusOK {
		ac.Cookie.SetLoginSession(w, r, body)

		if err := ac.ConnRedis.DelInputOld(sessID); err != nil {
			ac.Logger.Error(err.Error())
			ac.Response.Error(w, http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/", http.StatusSeeOther)
	}

	ac.Cookie.SetFlash(w, "error", []byte("Failed login attempt"))

	http.Redirect(w, r, "/signin", http.StatusSeeOther)
}

// SignOut signs out.
func (ac *AuthController) SignOut(w http.ResponseWriter, r *http.Request) {
	token, err := ac.Cookie.GetLoginSession(w, r)
	if err != nil {
		ac.Logger.Error(err.Error())
		ac.Response.Error(w, http.StatusInternalServerError)
		return
	}

	resp, err := ac.Client.GetSignOut(token)
	if err != nil {
		ac.Logger.Error(err.Error())
		ac.Response.Error(w, http.StatusInternalServerError)
	}

	if resp.StatusCode == http.StatusOK {
		ac.Cookie.DelLoginSession(w, r)

		http.Redirect(w, r, "/signin", http.StatusSeeOther)
	}

	ac.Response.Error(w, http.StatusInternalServerError)
	return
}
