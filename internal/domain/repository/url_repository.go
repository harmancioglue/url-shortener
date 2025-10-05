package repository

import "harmancioglue/url-shortener/internal/domain/entity"

type UrlRepository interface {
	Save(url *entity.URL) error
	FindByShortCode(shortCode string) (*entity.URL, error)
	FindByOriginalURL(originalURL string) (*entity.URL, error)
	UpdateClickCount(shortCode string) error
}
