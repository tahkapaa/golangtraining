package main

import "fmt"

func main() {
	out := make(chan int64)
	factorial(20, out)
	factorial(4, out)
	factorial(13, out)
	factorial(2, out)

	for i := 0; i < 4; i++ {
		fmt.Println("Result:", <-out)
	}
}

func factorial(n int64, out chan int64) {
	go func() {
		total := int64(1)
		for i := int64(1); i <= n; i++ {
			total *= i
		}
		out <- total
	}()
}
