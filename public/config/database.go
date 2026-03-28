package config

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/krittawatcode/go-soldier-mvc/models"
)

// ConnectDB establishes a connection to MySQL database using GORM
func ConnectDB() (*gorm.DB, error) {
	user := getEnv("MYSQL_USER", "root")
	password := getEnv("MYSQL_PASSWORD", "masonry1428")
	host := getEnv("MYSQL_HOST", "localhost")
	port := getEnv("MYSQL_PORT", "3306")
	database := getEnv("MYSQL_DATABASE", "masonry")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&charset=utf8mb4&loc=Local",
		user, password, host, port, database)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("error connecting to database: %w", err)
	}

	log.Println("Database connected successfully")
	return db, nil
}

// Migrate auto-migrates all models
func Migrate(db *gorm.DB) {
	err := db.AutoMigrate(
		&models.Gallery{},
		&models.Hashtag{},
	)
	if err != nil {
		log.Fatalf("Error migrating database: %v", err)
	}

	log.Println("Database migrated successfully")
	
	// Automatically execute Seeding after Migration
	SeedData(db)
}

func strPtr(s string) *string {
	return &s
}

// SeedData injects initial mock data to the database if it is empty
func SeedData(db *gorm.DB) {
	var count int64
	db.Model(&models.Gallery{}).Count(&count)
	
	if count == 0 {
		log.Println("Seeding initial mock data into database...")
		
		hashtags := []models.Hashtag{
			{Name: "Photography"}, {Name: "Landscape"}, {Name: "Monochrome"},
			{Name: "Urban"}, {Name: "Brutalism"}, {Name: "Portrait"}, {Name: "Night"}, {Name: "Street"},
		}
		for i := range hashtags {
			db.FirstOrCreate(&hashtags[i], models.Hashtag{Name: hashtags[i].Name})
		}

		galleries := []models.Gallery{
			{
				Name:  strPtr("Eternal Ridges"),
				Image: "https://placehold.co/600x800/e2e8f0/64748b?font=inter&text=Eternal+Ridges",
				Hashtags: []models.Hashtag{hashtags[0], hashtags[1]},
			},
			{
				Name:  strPtr("Neon Solitude"),
				Image: "https://placehold.co/600x450/e2e8f0/64748b?font=inter&text=Neon+Solitude",
				Hashtags: []models.Hashtag{hashtags[7], hashtags[6]},
			},
			{
				Name:  strPtr("Path of Spans"),
				Image: "https://placehold.co/600x1000/e2e8f0/64748b?font=inter&text=Path+of+Spans",
				Hashtags: []models.Hashtag{hashtags[1]},
			},
			{
				Name:  strPtr("Concrete Rhythm"),
				Image: "https://placehold.co/600x500/e2e8f0/64748b?font=inter&text=Concrete+Rhythm",
				Hashtags: []models.Hashtag{hashtags[4], hashtags[3]},
			},
			{
				Name:  strPtr("Silent Monochrome"),
				Image: "https://placehold.co/600x700/e2e8f0/64748b?font=inter&text=Silent+Monochrome",
				Hashtags: []models.Hashtag{hashtags[2], hashtags[5]},
			},
		}

		for _, gallery := range galleries {
			db.Create(&gallery)
		}
		
		log.Println("Database seeded successfully")
	} else {
		log.Println("Database already has data, skipping seeding")
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
