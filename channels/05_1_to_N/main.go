package main

import (
	"fmt"
	"sort"
)

type message struct {
	id       int
	msgCount int
	value    int
}

func (m message) String() string {
	return fmt.Sprintf("ID: %d, Messages: %d, Value: %d\n", m.id, m.msgCount, m.value)
}

type messages []message

func (m messages) Len() int {
	return len(m)
}

func (m messages) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m messages) Less(i, j int) bool {
	return m[i].id < m[j].id
}

func main() {
	n := 20
	c := make(chan int)
	result := make(chan message)

	go func() {
		for i := 0; i < 100000; i++ {
			c <- i
		}
		close(c)
	}()

	for i := 0; i < n; i++ {
		go catchMessage(c, result, i)
	}

	msgs := make(messages, n)
	for i := 0; i < n; i++ {
		msgs[i] = <-result
	}

	sort.Sort(msgs)
	fmt.Println(msgs)
}

func catchMessage(c chan int, result chan message, id int) {
	count := 0
	tv := 0
	for value := range c {
		count++
		tv += value
	}
	result <- message{id, count, tv}
}
