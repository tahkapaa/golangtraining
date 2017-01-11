package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string
	last  string // Unexported - not included in json representations
	Age   int    `json:"-"` // Not included in json representations
}

type climber struct {
	person
	TopGrade string `json:"top grade"`
}

func main() {
	p1 := climber{
		person{"Pasi", "Tähkäpää", 41},
		"6c"}

	s, err := json.Marshal(p1)
	if err != nil {
		fmt.Println("Failed to marshal: ", err.Error())
	}
	fmt.Println(s)
	fmt.Printf("%T\n", s)
	fmt.Println(string(s))
}

func (p person) fullName() string {
	if p.last == "" {
		return p.First
	}
	return p.First + " " + p.last
}
