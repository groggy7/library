package repositories

import (
	"database/sql"
	"library/pkg/models"
)

type psqlBookRepository struct {
	db *sql.DB
}

func NewPsqlBookRepository(db *sql.DB) BookRepository {
	return &psqlBookRepository{db: db}
}

func (p *psqlBookRepository) AddBook(book *models.Book) error {
	stmt, err := p.db.Prepare("INSERT INTO books (name, author, year) VALUES ($1, $2, $3)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(book.Name, book.Author, book.Year)
	if err != nil {
		return err
	}
	return nil
}

func (p *psqlBookRepository) GetBooks() ([]models.Book, error) {
	rows, err := p.db.Query("SELECT name, author, year FROM books")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []models.Book
	for rows.Next() {
		var book models.Book
		err = rows.Scan(&book.Name, &book.Author, &book.Year)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return books, nil
}

func (p *psqlBookRepository) RemoveBook(id int) error {
	stmt, err := p.db.Prepare("DELETE FROM books WHERE ID = $1")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		return nil
	}
	return nil
}
