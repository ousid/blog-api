package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model        // gives you ID,CreatedAt, UpdatedAt...
	Name       string `gorm:"uniqueIndex;not null" json:"name"`
	Slug       string `gorm:"uniqueIndex;not null" json:"slug"`
}
