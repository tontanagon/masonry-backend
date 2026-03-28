package models

import "time"

// Gallery represents the gallery model
type Gallery struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Name      *string   `json:"name"`
	Image     string    `json:"image" gorm:"not null" binding:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Hashtags  []Hashtag `json:"hashtags,omitempty" gorm:"many2many:gallery_hashtags;"`
}
