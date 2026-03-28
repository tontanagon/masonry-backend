package repositories

import (
	"github.com/krittawatcode/go-soldier-mvc/models"
	"gorm.io/gorm"
)

// HashtagRepository handles database operations for hashtags
type HashtagRepository struct {
	DB *gorm.DB
}

// NewHashtagRepository creates a new HashtagRepository
func NewHashtagRepository(db *gorm.DB) *HashtagRepository {
	return &HashtagRepository{DB: db}
}

// FindAll returns all hashtags
func (r *HashtagRepository) FindAll() ([]models.Hashtag, error) {
	var hashtags []models.Hashtag
	err := r.DB.Order("name ASC").Find(&hashtags).Error
	return hashtags, err
}

// FindByID returns a hashtag by ID
func (r *HashtagRepository) FindByID(id uint) (*models.Hashtag, error) {
	var hashtag models.Hashtag
	err := r.DB.First(&hashtag, id).Error
	if err != nil {
		return nil, err
	}
	return &hashtag, nil
}

// FindByIDs returns hashtags by multiple IDs
func (r *HashtagRepository) FindByIDs(ids []uint) ([]models.Hashtag, error) {
	var hashtags []models.Hashtag
	err := r.DB.Where("id IN ?", ids).Find(&hashtags).Error
	return hashtags, err
}

// Create inserts a new hashtag
func (r *HashtagRepository) Create(hashtag *models.Hashtag) error {
	return r.DB.Create(hashtag).Error
}

// Update updates an existing hashtag
func (r *HashtagRepository) Update(hashtag *models.Hashtag) error {
	return r.DB.Save(hashtag).Error
}

// Delete removes a hashtag by ID
func (r *HashtagRepository) Delete(id uint) error {
	return r.DB.Delete(&models.Hashtag{}, id).Error
}
