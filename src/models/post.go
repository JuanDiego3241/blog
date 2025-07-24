package models

import (
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model        // campos ID, CreatedAt, UpdatedAt, DeletedAt
	Title      string `gorm:"not null"`
	Image      string
	Content    string `gorm:"type:text"`
}

type Playlist struct {
	ID          string
	Name        string
	Description string
	Tracks      []Track
}

type Track struct {
	ID     string
	Name   string
	Artist string
	URI    string
}
