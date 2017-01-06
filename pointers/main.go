package main

import "fmt"

const metersToYards = 1.09361

func main() {
	a := 42
	fmt.Println("a: ", a)
	fmt.Println("Memory address: ", &a)
	fmt.Printf("Memory address(int): %d\n", &a)

	var b *int = &a
	fmt.Println("b: ", b)
	fmt.Println("dereferenced b: ", *b)

	*b = 43
	fmt.Println("a: ", a)

	var meters float64

	fmt.Println("Enter meters")
	fmt.Scanf("%v", &meters)
	yards := meters * metersToYards
	fmt.Println("Yards: ", yards)
}
