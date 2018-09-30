package main

import "fmt"

func main() {
	x := 42
	y := "James Bond"
	z := true

	fmt.Println(x, y, z)

	fmt.Printf("%+v\n", x)
	fmt.Printf("%+v\n", z)
	fmt.Printf("%+v\n", y)
}
