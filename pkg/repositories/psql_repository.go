package repositories

import (
	"database/sql"
	"library/pkg/models"
)

type psqlBookRepository struct {
	db *sql.DB
}

func NewBookRepository(db *sql.DB) BookRepository {
	return &psqlBookRepository{db: db}
}

func (p *psqlBookRepository) AddBook(book *models.Book) error {
	return nil
}

func (p *psqlBookRepository) GetBooks() ([]models.Book, error) {
	return nil, nil
}

func (p *psqlBookRepository) RemoveBook(id int) error {
	return nil
}
