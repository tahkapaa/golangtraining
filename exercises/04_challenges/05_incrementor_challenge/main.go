package main

import "fmt"

func main() {
	c1 := incrementor(2)
	count := adder(c1)
	fmt.Println("Final Counter:", count)
}

func adder(ch <-chan int64) int64 {
	var count int64
	for i := range ch {
		count += i
	}
	return count
}

func incrementor(n int) <-chan int64 {
	out := make(chan int64)
	done := make(chan bool)

	for i := 0; i < n; i++ {
		go func(id int) {
			for j := 0; j < 20; j++ {
				fmt.Printf("Process: %d printing %d\n", id, j)
				out <- 1
			}
			done <- true
		}(i + 1)
	}

	go func() {
		for i := 0; i < n; i++ {
			<-done
		}
		close(out)
	}()
	return out
}

/*
CHALLENGE #1
-- Take the code from above and change it to use channels instead of wait groups and atomicity
*/
