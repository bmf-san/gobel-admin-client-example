package main

import (
	"net/http"
	"os"
	"time"

	"github.com/bmf-san/gobel-admin-client-example/app/api"
	"github.com/bmf-san/gobel-admin-client-example/app/controller"
	"github.com/bmf-san/gobel-admin-client-example/app/cookie"
	"github.com/bmf-san/gobel-admin-client-example/app/logger"
	"github.com/bmf-san/gobel-admin-client-example/app/middleware"
	"github.com/bmf-san/gobel-admin-client-example/app/response"
	"github.com/bmf-san/gobel-admin-client-example/app/session"
	"github.com/bmf-san/goblin"
)

func init() {
	location := os.Getenv("TIME_ZONE")
	loc, err := time.LoadLocation(location)
	if err != nil {
		loc = time.FixedZone(location, 9*60*60)
	}
	time.Local = loc
}

func main() {
	logger := logger.NewLogger()
	connRedis, err := session.NewRedisHandler()
	if err != nil {
		logger.Error(err.Error())
	}
	cookie := cookie.NewCookie()
	client := api.NewClient()
	response := response.NewResponse()

	asset := middleware.NewAsset(connRedis, logger, cookie, response)

	homeController := controller.NewHomeController(logger, client, connRedis, response, cookie)
	authController := controller.NewAuthController(logger, client, connRedis, response, cookie)
	postController := controller.NewPostController(logger, client, connRedis, response, cookie)

	r := goblin.NewRouter()

	r.GET("/", middleware.NewMiddlewares(
		asset.CSRF,
		asset.Auth,
		asset.Session,
	).Then(homeController.Index))

	r.GET("/signin", middleware.NewMiddlewares(
		asset.Session,
		asset.CSRF,
	).Then(authController.Index))

	r.POST("/signin", middleware.NewMiddlewares(
		asset.Session,
		asset.CSRF,
	).Then(authController.SignIn))

	r.POST("/signout", middleware.NewMiddlewares(
		asset.Session,
		asset.CSRF,
	).Then(authController.SignOut))

	r.GET("/posts", middleware.NewMiddlewares(
		asset.Session,
		asset.Auth,
	).Then(postController.Index))

	r.GET("/posts/create", middleware.NewMiddlewares(
		asset.Session,
		asset.Auth,
	).Then(postController.Create))

	r.POST("/posts", middleware.NewMiddlewares(
		asset.Session,
		asset.Auth,
	).Then(postController.Store))

	r.GET("/posts/:id/edit", middleware.NewMiddlewares(
		asset.Session,
		asset.Auth,
	).Then(postController.Edit))

	r.POST("/posts/:id", middleware.NewMiddlewares(
		asset.Session,
		asset.Auth,
	).Then(postController.Update))

	// r.GET("/comments", middleware.NewMiddlewares(
	// 	asset.Session,
	// 	asset.Auth,
	// ).Then(commentController.Index))

	// r.GET("/categories", middleware.NewMiddlewares(
	// 	asset.Session,
	// 	asset.Auth,
	// ).Then(categoryController.Index))

	// r.GET("/categories/create", middleware.NewMiddlewares(
	// 	asset.Session,
	// 	asset.Auth,
	// ).Then(categoryController.Create))

	// r.POST("/categories", middleware.NewMiddlewares(
	// 	asset.Session,
	// 	asset.Auth,
	// ).Then(categoryController.Store))

	// r.GET("/categories/:id/edit", middleware.NewMiddlewares(
	// 	asset.Session,
	// 	asset.Auth,
	// ).Then(categoryController.Edit))

	// r.POST("/categories/:id", middleware.NewMiddlewares(
	// 	asset.Session,
	// 	asset.Auth,
	// ).Then(categoryController.Update))

	// r.GET("/tags", middleware.NewMiddlewares(
	// 	asset.Session,
	// 	asset.Auth,
	// ).Then(TagController.Index))

	// r.GET("/tags/create", middleware.NewMiddlewares(
	// 	asset.Session,
	// 	asset.Auth,
	// ).Then(TagController.Create))

	// r.POST("/tags", middleware.NewMiddlewares(
	// 	asset.Session,
	// 	asset.Auth,
	// ).Then(TagController.Store))

	// r.GET("/tags/:id/edit", middleware.NewMiddlewares(
	// 	asset.Session,
	// 	asset.Auth,
	// ).Then(TagController.Edit))

	// r.POST("/tags/:id", middleware.NewMiddlewares(
	// 	asset.Session,
	// 	asset.Auth,
	// ).Then(TagController.Update))

	if err := http.ListenAndServe(":"+os.Getenv("SERVER_PORT"), r); err != nil {
		logger.Error(err.Error())
	}
}
