package routes

import (
	"github.com/gofiber/fiber/v2"
)

func BindStats(app *fiber.App) {
	stats := app.Group("/stats")

	stats.Get("/healthcheck", healthcheck)
}

func healthcheck(c *fiber.Ctx) error {
	return c.Status(200).SendString("Server is healthy")
}
