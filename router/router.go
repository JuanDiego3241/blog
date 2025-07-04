package router

import (
	"github.com/JuanDiego3241/blog/src/controllers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:4200"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		AllowCredentials: true,
	}))
	r.HandleMethodNotAllowed = true

	r.POST("/posts", controllers.CreatePost)
	r.GET("/posts", controllers.GetPosts)

	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"error": "ruta no existe", "path": c.FullPath()})
	})

	return r
}
