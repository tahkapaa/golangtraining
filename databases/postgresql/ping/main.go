package main

import "database/sql"
import _ "github.com/lib/pq"

func main() {
	db, err := sql.Open("postgres", "postgres://pasi:pasi@localhost/company?sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	if err := db.Ping(); err != nil {
		panic(err)
	}

	println("DB Connection opened")
}
