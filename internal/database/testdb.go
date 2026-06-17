package database

import (
	"github.com/coderflexx/blog-api/internal/models"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
    "gorm.io/gorm/logger"
)

// SetupTestDB creates an in-memory SQLite DB for tests
func SetupTestDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		panic("failed to connect to test database")
	}

	db.AutoMigrate(
		&models.Category{},
		&models.Post{},
	)

	// override the global DB with the test DB
	DB = db

	return db
}

// CleanupTestDB wipes all tables between tests
func CleanupTestDB() {
	DB.Exec("DELETE FROM posts")
	DB.Exec("DELETE FROM categories")
}
