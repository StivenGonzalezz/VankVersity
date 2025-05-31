package http

import (
	"video-service/internal/service"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(router fiber.Router, videoService *service.VideoService) {
	//Endpoint para generar un link de subida de video
	router.Post("/upload", videoService.GenerateUploadLink)

	//Endpoint para listar todos los videos subidos
	router.Get("/videos", videoService.ListVideos)

	//Endpoint para obtener un video por su ID
	router.Get("/videos/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		if id == "" {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid ID"})
		}

		video, err := videoService.GetVideoByID(id)
		if err != nil {
			return c.Status(404).JSON(fiber.Map{"error": "Video not found"})
		}
		return c.JSON(video)
	})

	//Endpoint para eliminar un video por su ID
	router.Delete("/videos/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		if id == "" {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid ID"})
		}

		err := videoService.DeleteVideo(id)

		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Failed to delete video"})
		}

		return c.Status(204).JSON(fiber.Map{"message": "Video deleted successfully"})
	})
}
