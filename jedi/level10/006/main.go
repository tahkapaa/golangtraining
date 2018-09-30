package main

import "log"

func main() {
	c := generate()

	for v := range c {
		log.Printf("v: %#+v\n", v)
	}
}

func generate() <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
	}()

	return c
}
