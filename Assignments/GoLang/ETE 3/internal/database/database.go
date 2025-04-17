package database

import (
	"fmt"
	"log"
	"os"

	"cinema-booking-system/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// InitDB initializes the database connection
func InitDB() (*gorm.DB, error) {
	// Load database configuration from environment variables
	host := getEnv("DB_HOST", "localhost")
	port := getEnv("DB_PORT", "5432")
	user := getEnv("DB_USER", "postgres")
	password := getEnv("DB_PASSWORD", "postgres")
	dbname := getEnv("DB_NAME", "cinema_booking")

	// Create database connection string
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// Open connection to database
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Get generic database object sql.DB to use its functions
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	// Set connection pool settings
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)

	log.Println("Database connection established")
	return db, nil
}

// MigrateDB performs database migrations
func MigrateDB(db *gorm.DB) {
	// Auto migrate all models
	db.AutoMigrate(
		&models.User{},
		&models.Movie{},
		&models.Show{},
		&models.Seat{},
		&models.Booking{},
	)
	log.Println("Database migration completed")

	// Create foreign key relationships
	db.Model(&models.Show{}).AddForeignKey("movie_id", "movies(id)", "CASCADE", "CASCADE")
	db.Model(&models.Seat{}).AddForeignKey("show_id", "shows(id)", "CASCADE", "CASCADE")
	db.Model(&models.Booking{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")
	db.Model(&models.Booking{}).AddForeignKey("seat_id", "seats(id)", "CASCADE", "CASCADE")
	db.Model(&models.Booking{}).AddForeignKey("show_id", "shows(id)", "CASCADE", "CASCADE")
}

// getEnv gets the environment variable or returns a default value
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
