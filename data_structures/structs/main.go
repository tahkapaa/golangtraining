package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func main() {
	p1 := person{"Pasi", "Tähkäpää", 41}
	persons := make([]person, 2)
	persons[0] = p1
	persons[1] = person{"Miia", "Nmäki", 40}

	for _, p := range persons {
		fmt.Printf("Name: %s %s, Age: %d\n", p.first, p.last, p.age)
	}
}
