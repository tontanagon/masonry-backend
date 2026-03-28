package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/krittawatcode/go-soldier-mvc/controllers"
	"github.com/krittawatcode/go-soldier-mvc/repositories"
	"github.com/krittawatcode/go-soldier-mvc/services"
)

// SetupRouter initializes the Gin router with all routes
func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	// Health check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	// Initialize layers (dependency injection)
	galleryRepo := repositories.NewGalleryRepository(db)
	hashtagRepo := repositories.NewHashtagRepository(db)

	galleryService := services.NewGalleryService(galleryRepo, hashtagRepo)
	hashtagService := services.NewHashtagService(hashtagRepo)

	galleryController := controllers.NewGalleryController(galleryService)
	hashtagController := controllers.NewHashtagController(hashtagService)

	// API routes
	api := r.Group("/api")
	{
		// Gallery routes
		galleries := api.Group("/galleries")
		{
			galleries.GET("", galleryController.GetAll)
			galleries.GET("/:id", galleryController.GetByID)
			galleries.POST("", galleryController.Create)
			galleries.PUT("/:id", galleryController.Update)
			galleries.DELETE("/:id", galleryController.Delete)
			galleries.POST("/:id/hashtags", galleryController.AttachHashtags)
		}

		// Hashtag routes
		hashtags := api.Group("/hashtags")
		{
			hashtags.GET("", hashtagController.GetAll)
			hashtags.GET("/:id", hashtagController.GetByID)
			hashtags.POST("", hashtagController.Create)
			hashtags.PUT("/:id", hashtagController.Update)
			hashtags.DELETE("/:id", hashtagController.Delete)
		}
	}

	return r
}
