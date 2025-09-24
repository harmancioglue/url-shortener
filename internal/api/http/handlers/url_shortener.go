package handlers

import (
	"github.com/gofiber/fiber/v2"
	"harmancioglue/url-shortener/internal/app"
	"harmancioglue/url-shortener/internal/dto/request"
	"harmancioglue/url-shortener/internal/dto/response"
)

type UrlController struct {
	Application app.Application
}

func (u *UrlController) Shorten(c *fiber.Ctx) error {
	var shortenUrlRequest request.ShortenURLRequest

	if err := c.BodyParser(&shortenUrlRequest); err != nil {
		return response.BadRequestResponse(c, "Invalid request", map[string]interface{}{})
	}

	return nil
}

func (u *UrlController) GetUrl(c *fiber.Ctx) error {
	return nil
}

func NewUrlController(Application app.Application) *UrlController {
	return &UrlController{Application: Application}
}
