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
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
