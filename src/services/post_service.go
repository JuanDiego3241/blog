package services

import (
	"github.com/JuanDiego3241/blog/src/models"
	"github.com/JuanDiego3241/blog/src/repository"
)

func GetAllPosts() []models.Post {
	return repository.GetAll()
}

func GetPostByID(id uint) (*models.Post, bool) {
	p := repository.GetByID(id)
	return p, p != nil
}

func CreatePost(p models.Post) models.Post {
	return repository.Create(p)
}
