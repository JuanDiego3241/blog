package repository

import "github.com/JuanDiego3241/blog/src/models"

var posts = []models.Post{}

func GetAll() []models.Post { return posts }

func GetByID(id uint) *models.Post {
	for _, p := range posts {
		if p.ID == id {
			return &p
		}
	}
	return nil
}

func Create(p models.Post) models.Post {
	p.ID = uint(len(posts) + 1)
	posts = append(posts, p)
	return p
}
