package services

import (
	"errors"

	"github.com/krittawatcode/go-soldier-mvc/models"
	"github.com/krittawatcode/go-soldier-mvc/repositories"
)

// HashtagService handles business logic for hashtags
type HashtagService struct {
	Repo *repositories.HashtagRepository
}

// NewHashtagService creates a new HashtagService
func NewHashtagService(repo *repositories.HashtagRepository) *HashtagService {
	return &HashtagService{Repo: repo}
}

// GetAll returns all hashtags
func (s *HashtagService) GetAll() ([]models.Hashtag, error) {
	return s.Repo.FindAll()
}

// GetByID returns a hashtag by ID
func (s *HashtagService) GetByID(id uint) (*models.Hashtag, error) {
	hashtag, err := s.Repo.FindByID(id)
	if err != nil {
		return nil, errors.New("hashtag not found")
	}
	return hashtag, nil
}

// Create creates a new hashtag
func (s *HashtagService) Create(hashtag *models.Hashtag) error {
	return s.Repo.Create(hashtag)
}

// Update updates an existing hashtag
func (s *HashtagService) Update(id uint, hashtag *models.Hashtag) error {
	existing, err := s.Repo.FindByID(id)
	if err != nil {
		return errors.New("hashtag not found")
	}

	existing.Name = hashtag.Name
	return s.Repo.Update(existing)
}

// Delete deletes a hashtag by ID
func (s *HashtagService) Delete(id uint) error {
	_, err := s.Repo.FindByID(id)
	if err != nil {
		return errors.New("hashtag not found")
	}
	return s.Repo.Delete(id)
}
