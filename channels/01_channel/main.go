package main

import (
	"fmt"
	"sync"
)

type message struct {
	sender string
	value  int
}

func (m message) String() string {
	return fmt.Sprintf("%s: %d", m.sender, m.value)
}

func main() {
	c := make(chan message)
	var wg sync.WaitGroup
	wg.Add(4)

	go sender("Foo", c, &wg)
	go sender("Bar", c, &wg)
	go sender("FooBar", c, &wg)

	go func() {
		for i := 0; i < 30; i++ {
			fmt.Println(<-c)
		}
		wg.Done()
	}()
	wg.Wait()
}

func sender(s string, c chan message, wg *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		c <- message{s, i}
	}
	wg.Done()
}
