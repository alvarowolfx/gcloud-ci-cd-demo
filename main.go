package main

import (
	"os"

	"github.com/gofiber/fiber"
)

func index(c *fiber.Ctx) {
	c.JSON(fiber.Map{"message": "Cloud Summit Cerrado"})
}

func about(c *fiber.Ctx) {
	c.JSON(fiber.Map{"message": "Com a estratégia de nuvem certa, você pode ficar pronto para qualquer coisa"})
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
