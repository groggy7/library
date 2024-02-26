package services

import (
	"library/pkg/dto"
	"library/pkg/models"
	"library/pkg/repositories"
)

type BookService interface {
	AddBook(book *dto.AddBookRequest) error
	GetBooks() ([]models.Book, error)
	DeleteBook(req *dto.DeleteBookRequest) error
}

type bookService struct {
	bookRepo repositories.BookRepository
}

func NewBookService(bookRepo repositories.BookRepository) BookService {
	return &bookService{
		bookRepo: bookRepo,
	}
}

func (b *bookService) AddBook(book *dto.AddBookRequest) error {
	newBook := models.Book{
		Name:   book.Name,
		Author: book.Author,
		Year:   book.Year,
	}

	if err := b.bookRepo.AddBook(&newBook); err != nil {
		return err
	}
	return nil
}

func (b *bookService) GetBooks() ([]models.Book, error) {
	books, err := b.bookRepo.GetBooks()
	if err != nil {
		return nil, err
	}
	return books, nil
}

func (b *bookService) DeleteBook(req *dto.DeleteBookRequest) error {
	if err := b.bookRepo.DeleteBook(req.ID); err != nil {
		return err
	}
	return nil
}
