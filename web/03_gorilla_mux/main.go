package main

import (
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", index)
	r.HandleFunc("/foo", foo).Methods("GET")
	r.HandleFunc("/bar", bar)
	http.ListenAndServe(":3333", r)
}

func index(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Index")
}

func foo(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "foo")
}

func bar(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "foo")
}
