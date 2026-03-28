package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/krittawatcode/go-soldier-mvc/models"
	"github.com/krittawatcode/go-soldier-mvc/services"
)

// HashtagController handles HTTP requests for hashtags
type HashtagController struct {
	Service *services.HashtagService
}

// NewHashtagController creates a new HashtagController
func NewHashtagController(service *services.HashtagService) *HashtagController {
	return &HashtagController{Service: service}
}

// GetAll handles GET /hashtags
func (c *HashtagController) GetAll(ctx *gin.Context) {
	hashtags, err := c.Service.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": hashtags})
}

// GetByID handles GET /hashtags/:id
func (c *HashtagController) GetByID(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	hashtag, err := c.Service.GetByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": hashtag})
}

// Create handles POST /hashtags
func (c *HashtagController) Create(ctx *gin.Context) {
	var hashtag models.Hashtag
	if err := ctx.ShouldBindJSON(&hashtag); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.Service.Create(&hashtag); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"data": hashtag})
}

// Update handles PUT /hashtags/:id
func (c *HashtagController) Update(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var hashtag models.Hashtag
	if err := ctx.ShouldBindJSON(&hashtag); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.Service.Update(uint(id), &hashtag); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": hashtag})
}

// Delete handles DELETE /hashtags/:id
func (c *HashtagController) Delete(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := c.Service.Delete(uint(id)); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Hashtag deleted successfully"})
}
