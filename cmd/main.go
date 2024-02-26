package main

import (
	"library/pkg/handlers"
	"library/pkg/repositories"
	"library/pkg/services"

	"github.com/gofiber/fiber"
)

func main() {
	bookRepo := repositories.NewMockBookRepository()
	bookService := services.NewBookService(&bookRepo)
	bookHandler := handlers.NewBookHandler(&bookService)

	app := fiber.New()
	app.Get("/get", bookHandler.GetBooks)
	app.Post("/add", bookHandler.AddBook)
	app.Delete("/delete", bookHandler.DeleteBook)

	app.Listen(":8080")
}
