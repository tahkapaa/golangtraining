package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	var i int
	c, err := r.Cookie("counter")
	if err != nil {
		fmt.Println(err)
		c = &http.Cookie{Name: "counter", Value: "1", HttpOnly: true}
	}
	i, err = strconv.Atoi(c.Value)
	if err != nil {
		fmt.Println(err)
		c = &http.Cookie{Name: "counter", Value: "1", HttpOnly: true}
	}
	i++
	c.Value = strconv.Itoa(i)
	fmt.Println(c.Name, c.Value)
	http.SetCookie(w, c)
	fmt.Fprintln(w, c)
}
