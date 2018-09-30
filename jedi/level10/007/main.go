package main

import "log"

func main() {
	c := make(chan int)
	for i := 0; i < 10; i++ {
		go func() {
			for i := 0; i < 10; i++ {
				c <- i
			}
		}()
	}

	for i := 0; i < 100; i++ {
		log.Printf("v: %#+v\n", <-c)
	}
}
