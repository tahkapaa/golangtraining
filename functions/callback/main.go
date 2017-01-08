package main

import "fmt"

func main() {
	numSlice := []int{10, 20, 30, 40}
	printer := func(number int) {
		fmt.Println("Number: ", number)
	}
	visit(numSlice, printer)

	filterFunc := func(n int) bool {
		return n > 20
	}

	fmt.Println(filter(numSlice, filterFunc))
}

func visit(numbers []int, callback func(int)) {
	for _, value := range numbers {
		callback(value)
	}
}

func filter(numbers []int, filterFunc func(int) bool) []int {
	var filteredSlice []int
	for _, value := range numbers {
		if filterFunc(value) {
			filteredSlice = append(filteredSlice, value)
		}
	}
	return filteredSlice
}
