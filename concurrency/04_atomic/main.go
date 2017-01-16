package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

var counter int32
var wg sync.WaitGroup

func incrementer(s string) {
	for i := 0; i < 5; i++ {
		// Atomic pretects for writing values, but reading is not protected
		// -> Use the return value if you want to read
		x := atomic.AddInt32(&counter, 1)
		time.Sleep(time.Duration(rand.Intn(20)) * time.Microsecond)
		fmt.Println(s, i, "Counter:", x) // x is probably not the current value of counter anymore
	}
	wg.Done()
}

func main() {
	wg.Add(3)
	go incrementer("Foo")
	go incrementer("Bar")
	go incrementer("FooBar")
	wg.Wait()
	fmt.Println("Final counter:", counter)
}
