package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/me/", me)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Index")
}

func dog(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Dog")
}

func me(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Pasi")
}
