package main

import (
	"fmt"
)

func main() {
	quit := make(chan struct{})
	c := gen(quit)
	receive(c, quit)

	fmt.Println("about to exit")
}

func receive(c <-chan int, quit <-chan struct{}) {
	for {
		select {
		case v := <-c:
			if v != 0 {
				fmt.Println(v)
			}
		case _ = <-quit:
			return
		}
	}
}

func gen(quit chan struct{}) <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		quit <- struct{}{}
	}()

	return c
}
