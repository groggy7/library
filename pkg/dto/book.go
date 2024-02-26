package dto

type AddBookRequest struct {
	Name   string `json:"name"`
	Author string `json:"author"`
	Year   int16  `json:"year"`
}

type DeleteBookRequest struct {
	ID int `json:"id"`
}
