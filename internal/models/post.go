package models

import (
	"time"

	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Title       string     `gorm:"not null" json:"title"`
	Slug        string     `gorm:"uniqueIndex;not null" json:"slug"`
	Content     string     `gorm:"not null" json:"content"`
	Excerpt     string     `json:"excerpt"`
	CategoryID  uint       `json:"category_id"`
	Category    Category   `json:"category"`     // belongs to Category
	PublishedAt *time.Time `json:"published_at"` // nil = draft
}
