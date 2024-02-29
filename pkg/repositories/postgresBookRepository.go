package repositories

import (
	"library/pkg/models"

	"gorm.io/gorm"
)

type psqlBookRepository struct {
	db *gorm.DB
}

func NewPsqlBookRepository(db *gorm.DB) BookRepository {
	return &psqlBookRepository{db: db}
}

func (p *psqlBookRepository) AddBook(book *models.Book) error {
	result := p.db.Create(&book)
	return result.Error
}

func (p *psqlBookRepository) GetBooks() ([]models.Book, error) {
	var books []models.Book
	result := p.db.Find(&books)

	return books, result.Error
}

func (p *psqlBookRepository) DeleteBook(id int) error {
	result := p.db.Delete(&models.Book{}, id)
	return result.Error
}
