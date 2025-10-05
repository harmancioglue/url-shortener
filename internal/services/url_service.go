package services

import (
	"errors"
	"time"

	"harmancioglue/url-shortener/internal/common/utils"
	"harmancioglue/url-shortener/internal/domain/entity"
	"harmancioglue/url-shortener/internal/domain/repository"
	"harmancioglue/url-shortener/internal/domain/service"
	"harmancioglue/url-shortener/internal/dto/request"
	"harmancioglue/url-shortener/internal/dto/response"
)

type UrlService struct {
	urlRepository repository.UrlRepository
	idGenerator   service.IDGenerator
}

func (u UrlService) ShortenUrl(request request.ShortenURLRequest) (*response.ShortenURLResponse, error) {
	existingURL, err := u.urlRepository.FindByOriginalURL(request.URL)
	if err != nil {
		return nil, err
	}
	if existingURL != nil {
		return &response.ShortenURLResponse{
			ShortURL: existingURL.ShortCode,
		}, nil
	}

	id, err := u.idGenerator.GenerateID()
	if err != nil {
		return nil, err
	}

	shortCode := utils.Encode(id)

	urlEntity := &entity.URL{
		ID:          id,
		ShortCode:   shortCode,
		OriginalURL: request.URL,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		IsActive:    true,
	}

	err = u.urlRepository.Save(urlEntity)
	if err != nil {
		return nil, err
	}

	return &response.ShortenURLResponse{
		ShortURL: shortCode,
	}, nil
}

func (u UrlService) GetOriginalURL(shortCode string) (*response.GetURLResponse, error) {
	url, err := u.urlRepository.FindByShortCode(shortCode)
	if err != nil {
		return nil, err
	}
	if url == nil {
		return nil, errors.New("URL not found")
	}

	err = u.urlRepository.UpdateClickCount(shortCode)
	if err != nil {
		// Log error but don't fail the request
		// TODO: Add proper logging
	}

	return &response.GetURLResponse{
		OriginalURL: url.OriginalURL,
		ClickCount:  url.ClickCount,
	}, nil
}

func NewUrlService(urlRepository repository.UrlRepository, idGenerator service.IDGenerator) UrlService {
	return UrlService{
		urlRepository: urlRepository,
		idGenerator:   idGenerator,
	}
}
