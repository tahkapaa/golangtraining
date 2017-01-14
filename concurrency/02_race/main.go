package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var counter int

func main() {
	wg.Add(2)
	go incrementor("Foo")
	go incrementor("Bar")
	wg.Wait()
	fmt.Println("Its all over! Counter:", counter)
}

func incrementor(s string) {
	for i := 0; i < 5; i++ {
		counter++
		fmt.Println(s, i, "Counter:", counter)
	}
	wg.Done()
}
