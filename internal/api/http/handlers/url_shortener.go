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
		return response.BadRequestResponse(c, "Invalid request body", map[string]interface{}{
			"error": err.Error(),
		})
	}

	if shortenUrlRequest.URL == "" {
		return response.BadRequestResponse(c, "URL is required", map[string]interface{}{
			"field": "url",
		})
	}

	result, err := u.Application.UrlService.ShortenUrl(shortenUrlRequest)
	if err != nil {
		return response.BadRequestResponse(c, "Failed to shorten URL", map[string]interface{}{
			"error": err.Error(),
		})
	}

	return response.CreatedResponse(c, "URL shortened successfully", result)
}

func (u *UrlController) GetUrl(c *fiber.Ctx) error {
	shortCode := c.Params("url")
	if shortCode == "" {
		return response.BadRequestResponse(c, "Short code is required", map[string]interface{}{
			"field": "url",
		})
	}

	result, err := u.Application.UrlService.GetOriginalURL(shortCode)
	if err != nil {
		return response.BadRequestResponse(c, "Failed to get original URL", map[string]interface{}{
			"error": err.Error(),
		})
	}

	if result == nil {
		return response.NotFoundResponse(c, "URL not found", map[string]interface{}{
			"short_code": shortCode,
		})
	}

	// For redirect, we could use c.Redirect(result.OriginalURL)
	// But for now, return JSON response
	return response.SuccessResponse(c, "URL retrieved successfully", result)
}

func NewUrlController(Application app.Application) *UrlController {
	return &UrlController{Application: Application}
}
