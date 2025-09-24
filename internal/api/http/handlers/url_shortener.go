package handlers

import (
	"github.com/gofiber/fiber/v2"
	"harmancioglue/url-shortener/internal/app"
)

type UrlController struct {
	Application app.Application
}

func (u *UrlController) Shorten(c *fiber.Ctx) error {
	return nil
}

func (u *UrlController) GetUrl(c *fiber.Ctx) error {
	return nil
}

func NewUrlController(Application app.Application) *UrlController {
	return &UrlController{Application: Application}
}
