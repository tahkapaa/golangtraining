package main

import "fmt"

type person struct {
	name string
	age  int
}

func (p person) String() string {
	return fmt.Sprintf("Name: %s, Age: %d", p.name, p.age)
}

func main() {
	var s interface{} = "Test"

	fmt.Printf("%T\n", s)
	fmt.Printf("%s\n", s.(string))
	ss, ok := s.(float64)
	if !ok {
		fmt.Printf("Not ok: %v\n", ss)
	}

	var n interface{} = 9
	nn, ok := n.(string)
	if !ok {
		fmt.Printf("Not ok: %v\n", nn)
	}

	var p interface{} = person{"Pasi", 41}
	pp, ok := p.(person)
	if ok {
		fmt.Println(pp)
	}
}
