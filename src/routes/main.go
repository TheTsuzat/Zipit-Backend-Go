package routes

import (
	"github.com/Tsuzat/zipit/src/config"
	"github.com/gofiber/fiber/v2"
)

func InitRoutes() {
	// Health Check
	config.APP.Get("/api/v1/health", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"message": "status ok",
		})
	})

	// Register Auth Routes
	InitAuthRouter()
	// Register URL Routes
	InitUrlRouter()
}
