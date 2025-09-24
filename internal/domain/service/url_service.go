package service

import (
	"harmancioglue/url-shortener/internal/dto/request"
	"harmancioglue/url-shortener/internal/dto/response"
)

type UrlService interface {
	ShortenUrl(request request.ShortenURLRequest) (*response.ShortenURLResponse, error)
}
