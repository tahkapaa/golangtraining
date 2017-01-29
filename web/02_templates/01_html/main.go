package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

// PageData ...
type PageData struct {
	Name string
	Time time.Time
}

var friendsPageTpl *template.Template

func init() {
	tpl, err := template.ParseFiles("friends.html")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	friendsPageTpl = tpl
}

func main() {
	http.HandleFunc("/foo", foo)
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/friends", friends)
	http.ListenAndServe(":3333", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "foo")
}

func bar(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "bar")
}

func friends(w http.ResponseWriter, r *http.Request) {
	if err := friendsPageTpl.Execute(w, PageData{"Pasi", time.Now()}); err != nil {
		log.Fatalln(err)
	}
}
