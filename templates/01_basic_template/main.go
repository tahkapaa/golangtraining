package main

import (
	"fmt"
	"html/template"
	"os"
)

func main() {
	tpl, err := template.ParseFiles("letter.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	err = tpl.Execute(os.Stdout, "Pasi")
	if err != nil {
		fmt.Println(err)
	}

}
