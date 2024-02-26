package repositories

import "library/pkg/models"

type BookRepository interface {
	AddBook(book *models.Book) error
	GetBooks() ([]models.Book, error)
	DeleteBook(id int) error
}
