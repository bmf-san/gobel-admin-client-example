package controller

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"

	"github.com/bmf-san/gobel-admin-client-example/app/api"
	"github.com/bmf-san/gobel-admin-client-example/app/cookie"
	"github.com/bmf-san/gobel-admin-client-example/app/logger"
	"github.com/bmf-san/gobel-admin-client-example/app/model"
	"github.com/bmf-san/gobel-admin-client-example/app/response"
	"github.com/bmf-san/gobel-admin-client-example/app/session"
)

// A PostController is a controller for a post.
type PostController struct {
	Logger    *logger.Logger
	Client    *api.Client
	ConnRedis *session.RedisHandler
	Response  *response.Response
	Cookie    *cookie.Cookie
}

// NewPostController creates a NewPostController.
func NewPostController(logger *logger.Logger, apiClient *api.Client, connRedis *session.RedisHandler, response *response.Response, cookie *cookie.Cookie) *PostController {
	return &PostController{
		Logger:    logger,
		Client:    apiClient,
		ConnRedis: connRedis,
		Response:  response,
		Cookie:    cookie,
	}
}

// Index displays a listing of the resource.
func (pc *PostController) Index(w http.ResponseWriter, r *http.Request) {
	const defaultPage = 1
	const defaultLimit = 10

	var err error

	paramPage := r.URL.Query().Get("page")
	var page int
	if paramPage == "" {
		page = defaultPage
	} else {
		page, err = strconv.Atoi(paramPage)
		if err != nil {
			pc.Logger.Error(err.Error())
			pc.Response.Error(w, http.StatusInternalServerError)
			return
		}
	}

	paramLimit := r.URL.Query().Get("limit")
	var limit int
	if paramLimit == "" {
		limit = defaultLimit
	} else {
		limit, err = strconv.Atoi(paramLimit)
		if err != nil {
			pc.Logger.Error(err.Error())
			pc.Response.Error(w, http.StatusInternalServerError)
			return
		}
	}

	values := url.Values{}
	values.Add("page", strconv.Itoa(page))
	values.Add("limit", strconv.Itoa(limit))

	token, err := pc.Cookie.GetLoginSession(w, r)
	if err != nil {
		pc.Logger.Error(err.Error())
		pc.Response.Error(w, http.StatusInternalServerError)
		return
	}

	resp, body, err := pc.Client.GetPosts(values, token)
	if err != nil {
		pc.Logger.Error(err.Error())
		pc.Response.Error(w, http.StatusInternalServerError)
		return
	}

	var posts model.Posts

	if err := json.Unmarshal(body, &posts); err != nil {
		pc.Logger.Error(err.Error())
		pc.Response.Error(w, http.StatusInternalServerError)
		return
	}

	if resp.StatusCode == http.StatusOK {
		if err := pc.Response.ExecutePostIndex(w, &response.PostIndex{
			Posts: &posts,
		}); err != nil {
			pc.Logger.Error(err.Error())
			pc.Response.Error(w, http.StatusInternalServerError)
			return
		}
	}
	pc.Response.Error(w, http.StatusInternalServerError)
}

// Create displays a form for storing a newly resource in storage.
func (pc *PostController) Create(w http.ResponseWriter, r *http.Request) {
	flash, err := pc.Cookie.GetFlash(w, r)
	if err != nil {
		pc.Logger.Error(err.Error())
		pc.Response.Error(w, http.StatusInternalServerError)
		return
	}

	sessID, ok := r.Context().Value("session-id").(string)
	if !ok {
		pc.Logger.Error(err.Error())
		pc.Response.Error(w, http.StatusInternalServerError)
		return
	}

	old, err := pc.ConnRedis.GetInputOld(sessID)
	if err != nil {
		pc.Logger.Error(err.Error())
		pc.Response.Error(w, http.StatusInternalServerError)
		return
	}

	token, err := pc.Cookie.GetLoginSession(w, r)
	if err != nil {
		pc.Logger.Error(err.Error())
		pc.Response.Error(w, http.StatusInternalServerError)
		return
	}

	resp, body, err := pc.Client.GetCategories(token)
	if err != nil {
		pc.Logger.Error(err.Error())
		pc.Response.Error(w, http.StatusInternalServerError)
	}
	var categories *model.Categories
	if err := json.Unmarshal(body, &categories); err != nil {
		pc.Logger.Error(err.Error())
		pc.Response.Error(w, http.StatusInternalServerError)
		return
	}

	resp, body, err = pc.Client.GetTags(token)
	if err != nil {
		pc.Logger.Error(err.Error())
		pc.Response.Error(w, http.StatusInternalServerError)
	}
	var tags *model.Tags
	if err := json.Unmarshal(body, &tags); err != nil {
		pc.Logger.Error(err.Error())
		pc.Response.Error(w, http.StatusInternalServerError)
		return
	}

	var p *model.Post

	if resp.StatusCode == http.StatusOK {
		if err := pc.Response.ExecutePostCreate(w, &response.PostCreate{
			Flash:      flash,
			Old:        old,
			Categories: categories,
			Tags:       tags,
			Status:     p.GetStatusList(),
		}); err != nil {
			pc.Logger.Error(err.Error())
			pc.Response.Error(w, http.StatusInternalServerError)
			return
		}
	}
	pc.Response.Error(w, http.StatusInternalServerError)
}

// Store stores a newly created resource in storage.
func (pc *PostController) Store(w http.ResponseWriter, r *http.Request) {
}

// Edit displays a form for updating the specified resource.
func (pc *PostController) Edit(w http.ResponseWriter, r *http.Request) {
	// // TODO: createの実装してからやる

	// id := goblin.GetParam(r.Context(), "id")
}

// Update updates the specified resource in storage.
func (pc *PostController) Update(w http.ResponseWriter, r *http.Request) {
}
