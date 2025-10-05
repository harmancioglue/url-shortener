package mysql

import (
	"errors"
	"gorm.io/gorm"
	"harmancioglue/url-shortener/internal/domain/entity"
	"harmancioglue/url-shortener/internal/domain/repository"
)

type UrlRepository struct {
	db *gorm.DB
}

func (u *UrlRepository) Save(url *entity.URL) error {
	return u.db.Create(url).Error
}

func (u *UrlRepository) FindByShortCode(shortCode string) (*entity.URL, error) {
	var url entity.URL
	err := u.db.Where("short_code = ? AND is_active = ?", shortCode, true).First(&url).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &url, nil
}

func (u *UrlRepository) FindByOriginalURL(originalURL string) (*entity.URL, error) {
	var url entity.URL
	err := u.db.Where("original_url = ? AND is_active = ?", originalURL, true).First(&url).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &url, nil
}

func (u *UrlRepository) UpdateClickCount(shortCode string) error {
	return u.db.Model(&entity.URL{}).Where("short_code = ?", shortCode).
		UpdateColumn("click_count", gorm.Expr("click_count + ?", 1)).Error
}

func NewUrlRepository(db *gorm.DB) repository.UrlRepository {
	return &UrlRepository{
		db: db,
	}
}
