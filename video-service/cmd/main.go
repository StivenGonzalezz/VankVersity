package main

import (
	"log"
	"video-service/internal/adapter/http"
	"video-service/internal/service"
	"video-service/internal/adapter/repository"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("No se pudo cargar el archivo .env, usando variables del sistema si existen.")
	}
}

func main() {
	repo := repository.NewPostgresRepo()
	video := service.NewVideoService(repo)

	router := fiber.New()

	router.Use(cors.New(cors.Config{
		AllowOrigins: "http://127.0.0.1:5500",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	http.SetupRoutes(router, video)
	router.Listen(":8080")
}
