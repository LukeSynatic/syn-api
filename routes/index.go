package routes

import "github.com/gofiber/fiber/v2"

func BindRoutes(app *fiber.App) {
	BindStats(app)
	BindV1(app)
}
