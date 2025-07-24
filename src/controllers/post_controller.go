package controllers

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"strconv"
	"time"

	"github.com/JuanDiego3241/blog/src/models"
	"github.com/JuanDiego3241/blog/src/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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

func UploadPostImage(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Imagen requerida"})
		return
	}

	title := c.PostForm("title")
	content := c.PostForm("content")

	ext := filepath.Ext(file.Filename)
	name := fmt.Sprintf("%d%s", time.Now().UnixNano(), ext)
	dst := filepath.Join("uploads", name)

	if err := c.SaveUploadedFile(file, dst); err != nil {
		log.Printf("Error al guardar archivo: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo guardar imagen"})
		return
	}

	// Crear registro en DB
	post := models.Post{
		Title:   title,
		Content: content,
		Image:   dst,
	}
	if err := db.Create(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error guardando datos"})
		return
	}

	c.JSON(http.StatusCreated, post)
}
