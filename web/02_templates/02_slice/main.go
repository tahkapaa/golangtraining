package main

import (
	"fmt"
	"html/template"
	"os"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("*.gohtml"))
}

func main() {
	xs := []string{"Pasi", "Miia", "Miuku"}
	err := tpl.Execute(os.Stdout, xs)
	if err != nil {
		fmt.Println(err)
	}
}
