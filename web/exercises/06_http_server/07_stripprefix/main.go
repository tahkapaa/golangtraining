package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("public"))
	stripped := http.StripPrefix("/resources", fs)
	http.Handle("/resources/", stripped)

	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles("templates/index.gohtml")
	if err != nil {
		log.Fatal(err)
	}
	if err := tpl.Execute(w, nil); err != nil {
		log.Fatal(err)
	}
}
