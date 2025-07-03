package router

import (
	"fmt"
	"log"

	"github.com/JuanDiego3241/blog/config"

	"github.com/JuanDiego3241/blog/src/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	posts := r.Group("/posts")
	{
		posts.GET("", controllers.GetPosts)
		posts.GET("/:id", controllers.GetPost)
		cfg, err := config.LoadConfig()
		if err != nil {
			log.Fatalf("Error loading configuration: %v", err)
		}
		addr := fmt.Sprintf(":%s", cfg.ServerPort)
		log.Printf("🏃‍♂️ Servidor escuchando en %s", addr)
		if err := r.Run(addr); err != nil {
			log.Fatalf("Error al iniciar servidor: %v", err)
		}
		return r
	}
}
