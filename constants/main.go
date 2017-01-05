package main

import "fmt"

const p = "death & taxes"
const (
	a = iota
	b = iota
	c = iota
)
const (
	d = iota
	e
	f
)
const (
	_  = iota
	kb = 1 << (iota * 10) // 1*10
	mb = 1 << (iota * 10) // 2*10
	gb = 1 << (iota * 10) // 3*10
	tb = 1 << (iota * 10) // 4*10
)

func main() {
	const q int = 42
	fmt.Println("p - ", p)
	fmt.Println("q - ", q)

	fmt.Println("a - ", a)
	fmt.Println("b - ", b)
	fmt.Println("c - ", c)

	fmt.Println("d - ", d)
	fmt.Println("e - ", e)
	fmt.Println("f - ", f)

	fmt.Printf("%b - %d\n", kb, kb)
	fmt.Printf("%b - %d\n", mb, mb)
	fmt.Printf("%b - %d\n", gb, gb)
	fmt.Printf("%b - %d\n", tb, tb)

	var i uint
	j := 0 + i

	fmt.Printf("%d", i+j)

	for ; i < 10; i++ {
		b := 1 << (i * 10)
		fmt.Printf("%b - %d\n", b, b)
	}

	hello := "Hello, 世界"
	fmt.Printf("%T: %v\n", "Hello, 世界", "Hello, 世界")
	fmt.Printf("%T: %v\n", hello, hello)

}
