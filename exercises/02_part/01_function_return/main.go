package main

import "fmt"

func divide(n int) (int, bool) {
	if n%2 == 0 {
		return n / 2, true
	}
	return n / 2, false
}

func main() {
	result, even := divide(11)
	fmt.Printf("Result: %d, Even: %t\n", result, even)
}
