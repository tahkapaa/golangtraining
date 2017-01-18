package main

import (
	"fmt"
	"html/template"
	"os"
)

type person struct {
	Name string
	Age  int
}

func main() {
	tpl, err := template.ParseFiles("letter.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	p := []person{
		person{"Pasi", 41},
		person{"Miia", 40},
	}

	err = tpl.Execute(os.Stdout, p)
	if err != nil {
		fmt.Println(err)
	}

}
