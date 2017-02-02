package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/dog/dog.jpg", dogPic)
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "foo ran")
}

func dog(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles("dog.gohtml")
	if err != nil {
		log.Fatal(err)
	}
	if err := tpl.Execute(w, nil); err != nil {
		log.Fatal(err)
	}
}

func dogPic(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "dog.jpg")
}
