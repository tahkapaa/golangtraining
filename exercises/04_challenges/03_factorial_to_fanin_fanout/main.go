package main

import (
	"fmt"
)

func main() {

	in := gen()

	f := factorial(in)

	for _, ch := range f {
		go func(c <-chan int) {
			for n := range c {
				fmt.Println(n)
			}
		}(ch)
	}
}

func gen() <-chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < 10000; i++ {
			for j := 3; j < 13; j++ {
				out <- j
			}
		}
		close(out)
	}()
	return out
}

func factorial(in <-chan int) []<-chan int {
	out := make([]<-chan int, 100)
	for n := range in {
		out = append(out, fact(n))
	}
	return out
}

func fact(n int) <-chan int {
	out := make(chan int)
	go func() {
		total := 1
		for i := n; i > 0; i-- {
			total *= i
		}
		out <- total
		close(out)
	}()
	return out
}

/*
CHALLENGE #1:
-- Change the code above to execute 1,000 factorial computations concurrently and in parallel.
-- Use the "fan out / fan in" pattern

CHALLENGE #2:
WATCH MY SOLUTION BEFORE DOING THIS CHALLENGE #2
-- While running the factorial computations, try to find how much of your resources are being used.
-- Post the percentage of your resources being used to this discussion: https://goo.gl/BxKnOL
*/
