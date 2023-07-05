package middleware

import "github.com/gofiber/fiber/v2"

func ReturnsJSON(c *fiber.Ctx) error {
	c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
	return c.Next()
}
