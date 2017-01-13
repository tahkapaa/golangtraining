package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2)
	go foo()
	go bar()
	wg.Wait()
	fmt.Println("Its all over!")
}

func foo() {
	for i := 0; i < 100; i++ {
		fmt.Println("Foo: ", i)
		time.Sleep(time.Duration(1) * time.Microsecond)
	}
	wg.Done()
}

func bar() {
	for i := 0; i < 100; i++ {
		fmt.Println("Bar: ", i)
		time.Sleep(time.Duration(1) * time.Microsecond)
	}
	wg.Done()
}
