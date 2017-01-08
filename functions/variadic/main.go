package main

import "fmt"

func main() {
	fmt.Println("Average: ", average(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
}

func average(number ...int) int {
	var sum int
	for i := range number {
		sum += i
	}

	return sum / len(number)
}
