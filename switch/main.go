package main

import "fmt"

type contact struct {
	a string
	i int
}

func switchOnType(x interface{}) {
	switch x.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case contact:
		fmt.Println("contact")
	case rune:
		fmt.Println("rune")
	default:
		fmt.Println("Unknown")
	}
}

func main() {

	switch "Miuku" {
	case "Miia":
		fmt.Println("Heippa Miia")
	case "Pasi":
		fmt.Println("Moikka Pasi")
	case "Miuku":
		fmt.Println("Moikka Miuku")
		fallthrough
	case "Pasi2", "Miia2":
		fmt.Println("Moikka Pasi, tai Miia")
	}

	switchOnType(1)
	switchOnType("1")
	switchOnType('1')
	switchOnType(uint8(1))
	switchOnType(contact{a: "1", i: 2})
}
