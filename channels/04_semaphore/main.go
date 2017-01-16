package main

import "fmt"

type message struct {
	sender string
	value  int
}

func (m message) String() string {
	return fmt.Sprintf("%s: %d", m.sender, m.value)
}

func main() {
	c := make(chan message)
	done := make(chan bool)

	go sender("Foo", c, done)
	go sender("Bar", c, done)
	go sender("FooBar", c, done)

	go func() {
		for i := 0; i < 3; i++ {
			fmt.Println("Waiting")
			<-done
			fmt.Println("Done")
		}
		close(c)
	}()

	for msg := range c {
		fmt.Println(msg)
	}
}

func sender(s string, c chan message, done chan bool) {
	for i := 0; i < 10; i++ {
		c <- message{s, i}
	}
	done <- true
}
