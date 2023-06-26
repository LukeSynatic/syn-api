package main

import (
	"fmt"
	"log"
	"os"
	"syn-api/middleware"
	"syn-api/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
}

func main() {
	app := fiber.New(fiber.Config{
		Prefork: true,
	})

	_, err := middleware.GetOrInitMongo()
	if err != nil {
		log.Fatal("Error: Connection to MongoDB failed")
	}
	defer middleware.CloseMongo()

	routes.BindRoutes(app)

	app.Listen(fmt.Sprintf(":%v", os.Getenv("PORT")))
}
