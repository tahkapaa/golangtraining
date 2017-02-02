package main

import "fmt"

// HANDS ON 3
// create an interface type that both person and secretAgent implement
// declare a func with a parameter of the interface’s type
// call that func in main and pass in a value of type person
// call that func in main and pass in a value of type secretAgent

type person struct {
	name string
	age  int
}

type secretAgent struct {
	person
	code string
}

type speaker interface {
	speak()
}

func (p person) speak() {
	fmt.Println("Hello there")
}

func (sa secretAgent) speak() {
	fmt.Println("Shaken, not stirred")
}

func speak(s speaker) {
	s.speak()
}

func main() {
	p := person{"Pasi", 41}
	sa := secretAgent{person{"Miia", 40}, "001"}
	speak(p)
	speak(sa)
}
