package services

import (
	"harmancioglue/url-shortener/internal/domain/repository"
	"harmancioglue/url-shortener/internal/dto/request"
	"harmancioglue/url-shortener/internal/dto/response"
)

type UrlService struct {
	urlRepository repository.UrlRepository
}

func (u UrlService) ShortenUrl(request request.ShortenURLRequest) (*response.ShortenURLResponse, error) {
	return nil, nil
}

func NewUrlService(urlRepository repository.UrlRepository) UrlService {
	return UrlService{urlRepository: urlRepository}
}
