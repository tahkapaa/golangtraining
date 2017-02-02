package main

import (
	"html/template"
	"io"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("*.gohtml"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/me/", me)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	if err := tpl.Execute(w, "index"); err != nil {
		io.WriteString(w, err.Error())
	}
}

func dog(w http.ResponseWriter, r *http.Request) {
	if err := tpl.Execute(w, "dog"); err != nil {
		io.WriteString(w, err.Error())
	}
}

func me(w http.ResponseWriter, r *http.Request) {
	if err := tpl.Execute(w, "Pasi"); err != nil {
		io.WriteString(w, err.Error())
	}
}
