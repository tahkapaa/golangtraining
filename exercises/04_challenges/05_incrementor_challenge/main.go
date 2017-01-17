package main

import "fmt"

func main() {
	c1 := incrementor("1")
	c2 := incrementor("2")

	count := adder(c1, c2)

	fmt.Println("Final Counter:", count)
}

func adder(channels ...<-chan int64) int64 {
	var count int64
	for _, ch := range channels {
		for i := range ch {
			count += i
		}
	}
	return count
}

func incrementor(s string) <-chan int64 {
	out := make(chan int64)
	go func() {
		for i := 0; i < 20; i++ {
			out <- 1
			fmt.Println("Process: "+s+" printing:", i)
		}
		close(out)
	}()
	return out
}

/*
CHALLENGE #1
-- Take the code from above and change it to use channels instead of wait groups and atomicity
*/
