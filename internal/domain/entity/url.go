package entity

import "time"

type URL struct {
	ID          int64     `json:"id" gorm:"primaryKey;autoIncrement"`
	ShortCode   string    `json:"short_code" gorm:"uniqueIndex;not null"`
	OriginalURL string    `json:"original_url" gorm:"not null"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	ExpiresAt   *time.Time `json:"expires_at,omitempty"`
	ClickCount  int64     `json:"click_count" gorm:"default:0"`
	IsActive    bool      `json:"is_active" gorm:"default:true"`
}
