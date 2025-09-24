package http

import (
	"github.com/gofiber/fiber/v2"
	"harmancioglue/url-shortener/internal/api/http/handlers"
	"harmancioglue/url-shortener/internal/app"
)

type API struct {
	Server      *fiber.App
	Application app.Application

	UrlController *handlers.UrlController
}

func NewApi(application *app.Application) *API {
	api := &API{}

	api.Server = fiber.New()

	api.setupControllers()
	api.setupRoutes()

	return api
}

func (a *API) setupRoutes() {
	a.Server.Post("/shorten", a.UrlController.Shorten)
	a.Server.Get("/:url", a.UrlController.GetUrl)
}

func (a *API) setupControllers() {
	a.UrlController = handlers.NewUrlController(a.Application)
}
