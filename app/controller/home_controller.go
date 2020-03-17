package controller

import (
	"net/http"

	"github.com/bmf-san/gobel-admin-client-example/app/api"
	"github.com/bmf-san/gobel-admin-client-example/app/cookie"
	"github.com/bmf-san/gobel-admin-client-example/app/logger"
	"github.com/bmf-san/gobel-admin-client-example/app/response"
	"github.com/bmf-san/gobel-admin-client-example/app/session"
)

// A HomeController is a controller for a home.
type HomeController struct {
	Logger    *logger.Logger
	Client    *api.Client
	ConnRedis *session.RedisHandler
	Response  *response.Response
	Cookie    *cookie.Cookie
}

// NewHomeController creates a HomeController.
func NewHomeController(logger *logger.Logger, apiClient *api.Client, connRedis *session.RedisHandler, response *response.Response, cookie *cookie.Cookie) *HomeController {
	return &HomeController{
		Logger:    logger,
		Client:    apiClient,
		ConnRedis: connRedis,
		Response:  response,
		Cookie:    cookie,
	}
}

// Index displays a listing of the resource.
func (hc *HomeController) Index(w http.ResponseWriter, r *http.Request) {
	token, ok := r.Context().Value("csrf-token").(string)
	if !ok {
		hc.Response.Error(w, http.StatusInternalServerError)
		return
	}

	loginSessID, err := hc.Cookie.GetLoginSession(w, r)
	userID, err := hc.ConnRedis.GetLoginSession(loginSessID)
	if err != nil {
		hc.Logger.Error(err.Error())
		hc.Response.Error(w, http.StatusInternalServerError)
		return
	}

	if err := hc.Response.ExectuteHomeIndex(w, &response.HomeIndex{
		UserID:    userID,
		CSRFToken: token,
	}); err != nil {
		hc.Logger.Error(err.Error())
		hc.Response.Error(w, http.StatusInternalServerError)
		return
	}
}
