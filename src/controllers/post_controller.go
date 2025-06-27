package controllers

import (
	"net/http"
	"strconv"

	"github.com/JuanDiego3241/blog/src/models"
	"github.com/JuanDiego3241/blog/src/services"
	"github.com/gin-gonic/gin"
)

func GetPosts(c *gin.Context) {
	posts := services.GetAllPosts()
	c.JSON(http.StatusOK, posts)
}

func GetPost(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if post, ok := services.GetPostByID(uint(id)); ok {
		c.JSON(http.StatusOK, post)
	} else {
		c.JSON(http.StatusNotFound, gin.H{"message": "no encontrado"})
	}
}

func CreatePost(c *gin.Context) {
	var p models.Post
	if err := c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newP := services.CreatePost(p)
	c.JSON(http.StatusCreated, newP)
}
