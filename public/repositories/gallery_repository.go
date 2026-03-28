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

// FindAll returns all galleries with their hashtags, optionally filtered by tags and paginated
func (r *GalleryRepository) FindAll(tags []string, page int, limit int) ([]models.Gallery, error) {
	var galleries []models.Gallery
	query := r.DB.Preload("Hashtags").Order("created_at DESC")

	if len(tags) > 0 {
		query = query.Joins("JOIN gallery_hashtags ON gallery_hashtags.gallery_id = galleries.id").
			Joins("JOIN hashtags ON hashtags.id = gallery_hashtags.hashtag_id").
			Where("hashtags.name IN ?", tags).
			Group("galleries.id")
	}

	offset := (page - 1) * limit
	err := query.Offset(offset).Limit(limit).Find(&galleries).Error
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
