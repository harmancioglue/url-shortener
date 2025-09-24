package services

import "harmancioglue/url-shortener/internal/domain/repository"

type UrlService struct {
	urlRepository repository.UrlRepository
}

func (u UrlService) Save() {
	//TODO implement me
	panic("implement me")
}

func NewUrlService(urlRepository repository.UrlRepository) UrlService {
	return UrlService{urlRepository: urlRepository}
}
