package repositories

import "library/pkg/models"

type mockBookRepository struct {
	books []models.Book
}

func NewMockBookRepository() mockBookRepository {
	return mockBookRepository{
		books: []models.Book{
			{ID: 1, Name: "Martin Eden", Author: "Jack London", Year: 1909},
			{ID: 2, Name: "To Kill a Mockingbird", Author: "Harper Lee", Year: 1960},
		},
	}
}

func (m *mockBookRepository) AddBook(book *models.Book) error {
	m.books = append(m.books, *book)
	return nil
}

func (m *mockBookRepository) GetBooks() ([]models.Book, error) {
	return m.books, nil
}

func (m *mockBookRepository) DeleteBook(id int) error {
	for i, book := range m.books {
		if id == book.ID {
			m.books = append(m.books[:i], m.books[i+1:]...)
		}
	}
	return nil
}
