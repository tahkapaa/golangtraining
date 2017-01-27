package main

import (
	"fmt"
	"html/template"
	"os"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseGlob("*.gohtml"))
}

func main() {
	xi := []int{1, 2, 3, 4, 5}
	err := tpl.ExecuteTemplate(os.Stdout, "page.gohtml", xi)
	if err != nil {
		fmt.Println(err)
	}
}

var fm = template.FuncMap{
	"tt": tenTimer,
	"sq": square,
}

func tenTimer(n int) int {
	return n * 10
}

func square(n int) int {
	return n * n
}
