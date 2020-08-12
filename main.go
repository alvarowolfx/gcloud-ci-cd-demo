package main

import (
	"os"

	"github.com/gofiber/fiber"
)

func index(c *fiber.Ctx) {
	c.JSON(fiber.Map{"message": "Hello World"})
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	app := fiber.New()
	app.Get("/", index)

	app.Listen(port)
}
