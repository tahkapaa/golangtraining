package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

type climber struct {
	person
	topGrade string
}

func main() {
	p1 := climber{
		person{"Pasi", "Tähkäpää", 41},
		"6c"}
	climbers := make([]climber, 2)
	climbers[0] = p1
	climbers[1] = climber{
		person{
			first: "Miia",
			age:   40},
		"6a"}

	for _, p := range climbers {
		fmt.Printf("Name: %s, Age: %d Grade: %s\n", p.fullName(), p.age, p.topGrade)
	}

	p1.greeting()
}

func (c climber) greeting() {
	fmt.Println("Hello! any toppings yet?")
}

func (p person) greeting() {
	fmt.Println("Hello!")
}

func (p person) fullName() string {
	if p.last == "" {
		return p.first
	}
	return p.first + " " + p.last
}
