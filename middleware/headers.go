package middleware

import "github.com/gofiber/fiber/v2"

func ReturnsJSON(c *fiber.Ctx) error {
	c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
	return c.Next()
}

func AcceptsEJSON(c *fiber.Ctx) error {
	c.Accepts("application/ejson")
	return c.Next()
}

func AcceptsJSON(c *fiber.Ctx) error {
	c.Accepts(fiber.MIMEApplicationJSON)
	return c.Next()
}

func EJSON(c *fiber.Ctx) error {
	if string(c.Request().Header.ContentType()) == "application/ejson" {
		c.Request().Header.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
	}

	return c.Next()
}
