package mysql

import (
	"gorm.io/gorm"
	"harmancioglue/url-shortener/internal/domain/repository"
)

type UrlRepository struct {
	db *gorm.DB
}

func (u *UrlRepository) SaveUrl() {

}

func NewUrlRepository(db *gorm.DB) repository.UrlRepository {
	return &UrlRepository{
		db: db,
	}
}
