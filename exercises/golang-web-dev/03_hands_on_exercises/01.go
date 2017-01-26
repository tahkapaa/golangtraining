package main

// Initialize a SLICE of int using a composite literal; print out the slice; range over the slice printing out just the index; range over the slice printing out both the index and the value
// solution

import "fmt"

func main() {
	xi := []int{1, 3, 4, 5, 8}
	fmt.Println(xi)
	for i, v := range xi {
		fmt.Printf("%d - %d\n", i, v)
	}
}
