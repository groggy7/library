package handlers

import (
	"library/pkg/models"

	"github.com/gofiber/fiber"
)

func GetBooks(ctx *fiber.Ctx) {
	book := models.Book{
		Name:   "Martin Eden",
		Author: "Jack London",
		Year:   2001,
	}

	if err := ctx.JSON(book); err != nil {
		ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
}
