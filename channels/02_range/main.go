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

	go sender("Foo", c)

	for msg := range c {
		fmt.Println(msg)
	}
}

func sender(s string, c chan message) {
	for i := 0; i < 100000; i++ {
		c <- message{s, i}
	}
	close(c)
}
