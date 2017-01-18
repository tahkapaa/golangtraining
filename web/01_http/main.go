package main

import (
	"io"
	"net/http"
)

func main() {
	h := http.NewServeMux()
	h.HandleFunc("/", index)
	h.HandleFunc("/bar", bar)
	h.HandleFunc("/foo", foo)
	http.ListenAndServe(":3300", h)
}

func foo(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "FOO")
}

func bar(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "BAR")
}

func index(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Index")
}
