package main

import "fmt"

func main() {
	c1 := incrementor("Foo")
	c2 := incrementor("Bar")
	c3 := incrementor("FooBar")
	p1 := puller(c1)
	p2 := puller(c2)
	p3 := puller(c3)
	fmt.Println("Counters: ", <-p1, <-p2, <-p3)
}

func incrementor(s string) <-chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			out <- i
			fmt.Println(s, i)
		}
		close(out)
	}()
	return out
}

func puller(c <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		var sum int
		for n := range c {
			sum += n
		}
		out <- sum
		close(out)
	}()
	return out
}
