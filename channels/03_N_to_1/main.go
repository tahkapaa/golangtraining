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

	sender("Foo", c, &wg)
	sender("Bar", c, &wg)
	sender("FooBar", c, &wg)
	sender("Fizz", c, &wg)
	sender("Buzz", c, &wg)
	sender("FizzBuzz", c, &wg)

	go func() {
		wg.Wait()
		close(c)
	}()

	for msg := range c {
		fmt.Println(msg)
	}
}

func sender(s string, c chan message, wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		for i := 0; i < 10; i++ {
			c <- message{s, i}
		}
		wg.Done()
	}()
}
