package services

import (
	"errors"

	"github.com/krittawatcode/go-soldier-mvc/models"
	"github.com/krittawatcode/go-soldier-mvc/repositories"
)

// GalleryService handles business logic for galleries
type GalleryService struct {
	GalleryRepo *repositories.GalleryRepository
	HashtagRepo *repositories.HashtagRepository
}

// NewGalleryService creates a new GalleryService
func NewGalleryService(galleryRepo *repositories.GalleryRepository, hashtagRepo *repositories.HashtagRepository) *GalleryService {
	return &GalleryService{
		GalleryRepo: galleryRepo,
		HashtagRepo: hashtagRepo,
	}
}

// GetAll returns all galleries
func (s *GalleryService) GetAll() ([]models.Gallery, error) {
	return s.GalleryRepo.FindAll()
}

// GetByID returns a gallery by ID
func (s *GalleryService) GetByID(id uint) (*models.Gallery, error) {
	gallery, err := s.GalleryRepo.FindByID(id)
	if err != nil {
		return nil, errors.New("gallery not found")
	}
	return gallery, nil
}

// Create creates a new gallery
func (s *GalleryService) Create(gallery *models.Gallery) error {
	return s.GalleryRepo.Create(gallery)
}

// Update updates an existing gallery
func (s *GalleryService) Update(id uint, gallery *models.Gallery) error {
	existing, err := s.GalleryRepo.FindByID(id)
	if err != nil {
		return errors.New("gallery not found")
	}

	existing.Name = gallery.Name
	existing.Image = gallery.Image
	return s.GalleryRepo.Update(existing)
}

// Delete deletes a gallery by ID
func (s *GalleryService) Delete(id uint) error {
	_, err := s.GalleryRepo.FindByID(id)
	if err != nil {
		return errors.New("gallery not found")
	}
	return s.GalleryRepo.Delete(id)
}

// AttachHashtags attaches hashtags to a gallery
func (s *GalleryService) AttachHashtags(galleryID uint, hashtagIDs []uint) error {
	gallery, err := s.GalleryRepo.FindByID(galleryID)
	if err != nil {
		return errors.New("gallery not found")
	}

	hashtags, err := s.HashtagRepo.FindByIDs(hashtagIDs)
	if err != nil {
		return errors.New("failed to find hashtags")
	}

	return s.GalleryRepo.AttachHashtags(gallery, hashtags)
}
