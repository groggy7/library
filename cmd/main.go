package main

import (
	"library/pkg/handlers"

	"github.com/gofiber/fiber"
)

func main() {
	app := fiber.New()

	app.Get("/get", handlers.GetBooks)

	app.Listen(":4444")
}
