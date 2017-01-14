package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup
var counter int
var m sync.Mutex

func main() {
	wg.Add(2)
	go incrementor("Foo")
	go incrementor("Bar")
	wg.Wait()
	fmt.Println("Its all over! Counter:", counter)
}

func incrementor(s string) {
	for i := 0; i < 5; i++ {
		m.Lock()
		counter++
		fmt.Println(s, i, "Counter:", counter)
		m.Unlock()
		time.Sleep(time.Duration(rand.Intn(3)) * time.Millisecond)
	}
	wg.Done()
}
