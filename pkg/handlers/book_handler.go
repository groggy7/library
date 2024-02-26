package handlers

import (
	"library/pkg/dto"
	"library/pkg/services"
	"log"

	"github.com/gofiber/fiber"
)

type BookHandler interface {
	GetBooks(ctx *fiber.Ctx)
	AddBook(ctx *fiber.Ctx)
	DeleteBook(ctx *fiber.Ctx)
}

type bookHandler struct {
	bookService services.BookService
}

func NewBookHandler(svc *services.BookService) BookHandler {
	return &bookHandler{
		bookService: *svc,
	}
}

func (h *bookHandler) AddBook(ctx *fiber.Ctx) {
	book := new(dto.AddBookRequest)

	err := ctx.BodyParser(book)
	if err != nil {
		log.Println(err)
		ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err})
		return
	}

	if err := h.bookService.AddBook(book); err != nil {
		log.Println(err)
		ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err})
	}
}

func (h *bookHandler) GetBooks(ctx *fiber.Ctx) {
	books, err := h.bookService.GetBooks()
	if err != nil {
		log.Println(err)
		ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err})
		return
	}

	if err := ctx.JSON(books); err != nil {
		ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
}

func (h *bookHandler) DeleteBook(ctx *fiber.Ctx) {
	req := new(dto.DeleteBookRequest)

	err := ctx.BodyParser(req)
	if err != nil {
		log.Println(err)
		ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err})
		return
	}

	if err := h.bookService.DeleteBook(req); err != nil {
		log.Println(err)
		ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err})
	}
}
