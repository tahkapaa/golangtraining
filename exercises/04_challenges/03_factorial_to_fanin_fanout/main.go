package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {
	in := gen()
	f := factorial(in)
	var i int
	for s := range receiver(f) {
		i++
		fmt.Println(i, s)
	}
}

func receiver(channels []<-chan int) <-chan string {
	out := make(chan string)
	var wg sync.WaitGroup
	wg.Add(len(channels))

	go func() {
		for _, ch := range channels {
			go func(c <-chan int) {
				for n := range c {
					out <- fmt.Sprint(n)
				}
				wg.Done()
			}(ch)
		}
		wg.Wait()
		close(out)
	}()
	return out
}

func gen() <-chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < 50000; i++ {
			out <- rand.Intn(10) + 3
		}
		close(out)
	}()
	return out
}

func factorial(in <-chan int) []<-chan int {
	out := make([]<-chan int, 300)
	for i := range out {
		out[i] = fact(in)
	}
	return out
}

func fact(ch <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range ch {
			total := 1
			for i := n; i > 0; i-- {
				total *= i
			}
			out <- total
		}
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
