package database

import (
	"log"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/coderflexx/blog-api/internal/models"
)

var DB *gorm.DB

func Connect() {
	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		dbPath = "./database.db" // fallback
	}

	var err error

	DB, err = gorm.Open(sqlite.Open(dbPath), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // show SQL queriries
	})

	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	migrate()

	log.Println("Database connected")
}

func migrate() {
	err := DB.AutoMigrate(
		&models.Category{},
		&models.Post{},
	)

	if err != nil {
		log.Fatal("Migration failed:", err)
	}

	log.Println("Database migrated")
}

func Seed() {
	category := models.Category{Name: "Laravel", Slug: "laravel"}
	DB.FirstOrCreate(&category, models.Category{Slug: "laravel"})

	posts := []models.Post{
		{
			Title:      "Getting Started with Go",
			Slug:       "getting-started-with-go",
			Content:    "Go is a statically typed language...",
			Excerpt:    "An intro to Go",
			CategoryID: category.ID,
		},
		{
			Title:      "GORM vs Eloquent",
			Slug:       "gorm-vs-eloquent",
			Content:    "Both are great ORMs but...",
			Excerpt:    "Comparing Go and Laravel ORMs",
			CategoryID: category.ID,
		},
	}

	for _, post := range posts {
		DB.FirstOrCreate(&post, models.Post{Slug: post.Slug})
	}

	log.Println("Database seeded")
}
