package main

import (
	"os"

	"github.com/gofiber/fiber"
)

func index(c *fiber.Ctx) {
	c.JSON(fiber.Map{"message": "Olá TDC 2020"})
}

func about(c *fiber.Ctx) {
	c.JSON(fiber.Map{"message": "Um grande evento online de desenvolvimento"})
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	app := fiber.New()
	app.Get("/", index)
	app.Get("/about", about)

	app.Listen(port)
}
