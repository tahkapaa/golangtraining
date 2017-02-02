package main

import "fmt"

// HANDS ON 2
// create a struct that holds person fields
// create a struct that holds secret agent fields and embeds person type
// attach a method to person: pSpeak
// attach a method to secret agent: saSpeak
// create a variable of type person
// create a variable of type secret agent
// print a field from person
// run pSpeak attached to the variable of type person
// print a field from secret agent
// run saSpeak attached to the variable of type secret agent
// run pSpeak attached to the variable of type secret agent

type person struct {
	name string
	age  int
}

type secretAgent struct {
	person
	code string
}

func (p person) pSpeak() {
	fmt.Println("Hello there")
}

func (sa secretAgent) saSpeak() {
	fmt.Println("Shaken, not stirred")
}

func main() {
	p := person{"Pasi", 41}
	sa := secretAgent{person{"Miia", 40}, "001"}
	fmt.Println(p.name)
	p.pSpeak()
	fmt.Println(sa.code)
	sa.saSpeak()
	sa.person.pSpeak()
}
