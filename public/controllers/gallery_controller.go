package controllers

import (
	"net/http"
	"strconv"

	"strings"

	"github.com/gin-gonic/gin"
	"github.com/krittawatcode/go-soldier-mvc/models"
	"github.com/krittawatcode/go-soldier-mvc/services"
)

// GalleryController handles HTTP requests for galleries
type GalleryController struct {
	Service *services.GalleryService
}

// NewGalleryController creates a new GalleryController
func NewGalleryController(service *services.GalleryService) *GalleryController {
	return &GalleryController{Service: service}
}

// Search handles POST /galleries/search with body filters
func (c *GalleryController) Search(ctx *gin.Context) {
	var req struct {
		Filters string `json:"filters"`
		Page    int    `json:"page"`
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		// Fallback to defaults if body is missing or malformed
		req.Filters = "all"
		req.Page = 1
	}

	var tags []string
	if req.Filters != "" && req.Filters != "all" {
		rawTags := strings.Split(req.Filters, ",")
		for _, tag := range rawTags {
			cleaned := strings.TrimSpace(strings.TrimPrefix(tag, "#"))
			if cleaned != "" {
				tags = append(tags, cleaned)
			}
		}
	}

	page := 1
	if req.Page > 0 {
		page = req.Page
	}

	limit := 8

	galleries, err := c.Service.GetAll(tags, page, limit)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": galleries})
}

// GetByID handles GET /galleries/:id
func (c *GalleryController) GetByID(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	gallery, err := c.Service.GetByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": gallery})
}

// Create handles POST /galleries
func (c *GalleryController) Create(ctx *gin.Context) {
	var gallery models.Gallery
	if err := ctx.ShouldBindJSON(&gallery); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.Service.Create(&gallery); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"data": gallery})
}

// Update handles PUT /galleries/:id
func (c *GalleryController) Update(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var gallery models.Gallery
	if err := ctx.ShouldBindJSON(&gallery); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.Service.Update(uint(id), &gallery); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": gallery})
}

// Delete handles DELETE /galleries/:id
func (c *GalleryController) Delete(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := c.Service.Delete(uint(id)); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Gallery deleted successfully"})
}

// AttachHashtags handles POST /galleries/:id/hashtags
func (c *GalleryController) AttachHashtags(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var body struct {
		HashtagIDs []uint `json:"hashtag_ids" binding:"required"`
	}
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.Service.AttachHashtags(uint(id), body.HashtagIDs); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Hashtags attached successfully"})
}
