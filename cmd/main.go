package main

import (
	"library/internal/db"
	"library/pkg/handlers"
	"library/pkg/repositories"
	"library/pkg/services"
	"log"

	"github.com/gofiber/fiber"
)

func main() {
	db, err := db.NewDB()
	if err != nil {
		log.Fatalln(err)
	}

	bookRepo := repositories.NewPsqlBookRepository(db)
	bookService := services.NewBookService(bookRepo)
	bookHandler := handlers.NewBookHandler(&bookService)

	app := fiber.New()
	app.Get("/get", bookHandler.GetBooks)
	app.Post("/add", bookHandler.AddBook)
	app.Delete("/delete", bookHandler.DeleteBook)

	app.Listen(":8080")
}
