package main

import "fmt"

func main() {
	fmt.Println("Largest: ", largest(4, 5, 6, 7, 8, 9, 123, 5, 6, 7, 2))
}

func largest(numbers ...int) int {
	var n int
	for _, value := range numbers {
		if value > n {
			n = value
		}
	}
	return n
}
