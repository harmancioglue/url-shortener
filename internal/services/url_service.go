package services

import (
	"errors"
	"fmt"
	"time"

	"harmancioglue/url-shortener/internal/common/utils"
	"harmancioglue/url-shortener/internal/config"
	"harmancioglue/url-shortener/internal/domain/entity"
	"harmancioglue/url-shortener/internal/domain/repository"
	"harmancioglue/url-shortener/internal/domain/service"
	"harmancioglue/url-shortener/internal/dto/request"
	"harmancioglue/url-shortener/internal/dto/response"
)

type UrlService struct {
	urlRepository repository.UrlRepository
	idGenerator   service.IDGenerator
	config        *config.Config
}

func (u UrlService) ShortenUrl(request request.ShortenURLRequest) (*response.ShortenURLResponse, error) {
	existingURL, err := u.urlRepository.FindByOriginalURL(request.URL)
	if err != nil {
		return nil, err
	}
	if existingURL != nil {
		fullShortURL := fmt.Sprintf("http://%s:%d/%s", u.config.Server.Host, u.config.Server.Port, existingURL.ShortCode)
		return &response.ShortenURLResponse{
			ShortURL:    fullShortURL,
			OriginalURL: existingURL.OriginalURL,
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

	fullShortURL := fmt.Sprintf("http://%s:%d/%s", u.config.Server.Host, u.config.Server.Port, shortCode)
	return &response.ShortenURLResponse{
		ShortURL:    fullShortURL,
		OriginalURL: request.URL,
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

func NewUrlService(urlRepository repository.UrlRepository, idGenerator service.IDGenerator, config *config.Config) UrlService {
	return UrlService{
		urlRepository: urlRepository,
		idGenerator:   idGenerator,
		config:        config,
	}
}
