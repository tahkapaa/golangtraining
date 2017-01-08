package main

import "fmt"

func main() {
	expression := func() {
		fmt.Println("Expression")
	}
	expression()
	fmt.Printf("type: %T\n", expression)

	greet := makeGreeter()
	greet()
	fmt.Printf("type: %T\n", greet)
}

func makeGreeter() func() string {
	return func() string {
		return fmt.Sprintln("Hello world!")
	}
}
