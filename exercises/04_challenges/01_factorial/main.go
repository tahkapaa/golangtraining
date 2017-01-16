package main

import (
	"fmt"
	"math/big"
)

func main() {
	c := factorReceiver(factorial(20), factorial(100), factorial(13), factorial(4))

	for s := range c {
		fmt.Println("Result:", s)
	}
}

func factorial(n int64) <-chan big.Int {
	out := make(chan big.Int)
	go func() {
		total := big.NewInt(n)
		for i := big.NewInt(1); i.Cmp(big.NewInt(n)) == -1; i.Add(i, big.NewInt(1)) {
			total.Mul(total, i)
		}
		out <- *total
		close(out)
	}()
	return out
}

func factorReceiver(channels ...<-chan big.Int) <-chan string {
	out := make(chan string)
	go func() {
		for _, c := range channels {
			for n := range c {
				out <- fmt.Sprint(n)
			}
		}
		close(out)
	}()
	return out
}
