package router

import (
	"github.com/JuanDiego3241/blog/src/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	posts := r.Group("/posts")
	{
		posts.GET("", controllers.GetPosts)
		posts.GET("/:id", controllers.GetPost)
		posts.POST("", controllers.CreatePost)
	}
	return r
}
