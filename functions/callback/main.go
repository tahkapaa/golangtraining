package main

import "fmt"

func main() {
	numSlice := []int{10, 20, 30, 40}
	printer := func(number int) {
		fmt.Println("Number: ", number)
	}
	visit(numSlice, printer)
}

func visit(numbers []int, callback func(int)) {
	for _, value := range numbers {
		callback(value)
	}
}
