package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model        // campos ID, CreatedAt, UpdatedAt, DeletedAt
	Title      string `gorm:"not null"`
	Image      string
	Content    string `gorm:"type:text"`
}
