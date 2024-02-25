package main

import (
	"fmt"
	"library/internal/db"
	"library/pkg/repositories"
	"log"
)

func main() {
	db, err := db.NewDB()
	if err != nil {
		log.Fatal(err)
	}

	rp := repositories.NewBookRepository(db)

	/*
		rp.AddBook(&models.Book{
			Name:   "Martin Eden",
			Author: "Jack London",
			Year:   2001,
		})
	*/

	rp.RemoveBook(2)

	fmt.Println(rp.GetBooks())
}
