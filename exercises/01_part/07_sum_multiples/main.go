package main

import "fmt"

func main() {
	var sum int
	for i := 0; i < 1000; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}
	fmt.Println("Sum of all multiples of 3 and 5 under 1000: ", sum)
}
