package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/lib/pq"
)

type book struct {
	Isbn   string
	Title  string
	Author string
	Price  float32
}

func (b book) String() string {
	return fmt.Sprintf("ISBN: %s, TITLE: %s, AUTHOR: %s, PRICE: %f", b.Isbn, b.Title, b.Author, b.Price)
}

func main() {
	db, err := sql.Open("postgres", "postgres://pasi:pasi@localhost/company?sslmode=disable")
	if err != nil {
		panic(err)
	}

	http.Handle("/", http.HandlerFunc(getBooks(db)))
	http.ListenAndServe(":3000", nil)
}

func getBooks(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		bs, err := fetchBooks(db)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		b, err := json.MarshalIndent(bs, " ", "  ")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		_, err = w.Write(b)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func fetchBooks(db *sql.DB) ([]book, error) {
	rows, err := db.Query("SELECT * FROM books")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	bs := make([]book, 0)
	for rows.Next() {
		var b book
		if err = rows.Scan(&b.Isbn, &b.Title, &b.Author, &b.Price); err != nil {
			return nil, err
		}
		bs = append(bs, b)
	}
	return bs, err
}
