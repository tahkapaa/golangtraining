package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/pasi", pasi)
	http.ListenAndServe(":3333", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "FOO")
}

func bar(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "BAR")
}

func pasi(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Pasi")
}
