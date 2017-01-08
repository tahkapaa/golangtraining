package main

import "fmt"

func main() {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("Average: ", average(data...))
}

func average(number ...int) int {
	var sum int
	for _, value := range number {
		sum += value
	}

	return sum / len(number)
}
