package main

// Initialize a MAP using a composite literal where the key is a string and the value is an int; print out the map; range over the map printing out just the key; range over the map printing out both the key and the value
// solution

import "fmt"

func main() {
	items := map[string]int{
		"first":  1,
		"second": 2,
		"third":  3,
	}
	fmt.Println(items)

	for k := range items {
		fmt.Println(k)
	}

	for k, v := range items {
		fmt.Printf("%s - %d\n", k, v)
	}
}
