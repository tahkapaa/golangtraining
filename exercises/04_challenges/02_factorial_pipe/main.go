package main

import (
	"fmt"
	"math/rand"
)

func main() {
	c := factorial(gen(100))
	for n := range c {
		fmt.Println(n)
	}
}

func gen(n int) chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < n; i++ {
			out <- rand.Intn(16) + 3
		}
		close(out)
	}()
	return out
}

func factorial(c chan int) chan int {
	out := make(chan int)
	go func() {
		for n := range c {
			out <- fact(n)
		}
		close(out)
	}()
	return out
}

func fact(n int) int {
	total := 1
	for i := 1; i <= n; i++ {
		total *= i
	}
	return total
}
