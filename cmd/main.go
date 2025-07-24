package main

import (
	"fmt"
	"log"

	"github.com/JuanDiego3241/blog/config"
	"github.com/JuanDiego3241/blog/router"
	"github.com/JuanDiego3241/blog/src/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("No se cargó config: %v", err)
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		cfg.DBHost, cfg.DBUser, cfg.DBPassword,
		cfg.DBName, cfg.DBPort, cfg.DBSSLMode, cfg.DBTimeZone,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error al conectar BD: %v", err)
	}
	sqlDB, _ := db.DB()
	if err := sqlDB.Ping(); err != nil {
		log.Fatalf("Ping BD falló: %v", err)
	}
	log.Println("Conexión a DB exitosa")

	if err := db.AutoMigrate(&models.Post{}); err != nil {
		log.Fatalf("Migración automática falló: %v", err)
	}
	log.Println("Migración automática exitosa")

	r := router.SetupRouter()

	addr := fmt.Sprintf(":%s", cfg.ServerPort)
	log.Printf("Servidor escuchando en %s", addr)

	r.Static("/uploads", "./uploads")
	if err := r.Run(addr); err != nil {
		log.Fatalf("Error al iniciar servidor: %v", err)
	}
}
