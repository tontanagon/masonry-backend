package repositories

import (
	"github.com/krittawatcode/go-soldier-mvc/models"
	"gorm.io/gorm"
)

// GalleryRepository handles database operations for galleries
type GalleryRepository struct {
	DB *gorm.DB
}

// NewGalleryRepository creates a new GalleryRepository
func NewGalleryRepository(db *gorm.DB) *GalleryRepository {
	return &GalleryRepository{DB: db}
}

// FindAll returns all galleries with their hashtags
func (r *GalleryRepository) FindAll() ([]models.Gallery, error) {
	var galleries []models.Gallery
	err := r.DB.Preload("Hashtags").Order("created_at DESC").Find(&galleries).Error
	return galleries, err
}

// FindByID returns a gallery by ID with its hashtags
func (r *GalleryRepository) FindByID(id uint) (*models.Gallery, error) {
	var gallery models.Gallery
	err := r.DB.Preload("Hashtags").First(&gallery, id).Error
	if err != nil {
		return nil, err
	}
	return &gallery, nil
}

// Create inserts a new gallery
func (r *GalleryRepository) Create(gallery *models.Gallery) error {
	return r.DB.Create(gallery).Error
}

// Update updates an existing gallery
func (r *GalleryRepository) Update(gallery *models.Gallery) error {
	return r.DB.Save(gallery).Error
}

// Delete removes a gallery by ID
func (r *GalleryRepository) Delete(id uint) error {
	return r.DB.Delete(&models.Gallery{}, id).Error
}

// AttachHashtags replaces the hashtags for a gallery
func (r *GalleryRepository) AttachHashtags(gallery *models.Gallery, hashtags []models.Hashtag) error {
	return r.DB.Model(gallery).Association("Hashtags").Replace(hashtags)
}
