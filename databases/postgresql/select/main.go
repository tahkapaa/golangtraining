package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type book struct {
	isbn   string
	title  string
	author string
	price  float32
}

func (b book) String() string {
	return fmt.Sprintf("ISBN: %s, TITLE: %s, AUTHOR: %s, PRICE: %f", b.isbn, b.title, b.author, b.price)
}

func main() {
	db, err := sql.Open("postgres", "postgres://pasi:pasi@localhost/company?sslmode=disable")
	if err != nil {
		panic(err)
	}

	bs := fetchBooks(db)

	for _, b := range bs {
		fmt.Printf("%v\n", b)
	}
}

func fetchBooks(db *sql.DB) []book {
	rows, err := db.Query("SELECT * FROM books")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	bs := make([]book, 0)
	for rows.Next() {
		var b book
		if err := rows.Scan(&b.isbn, &b.title, &b.author, &b.price); err != nil {
			panic(err)
		}
		bs = append(bs, b)
	}
	return bs
}
