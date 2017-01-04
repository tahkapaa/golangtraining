package main

import "fmt"

func main() {
	a := 50
	b := true
	c := "Joo"
	d := 50.1234
	e := []byte(c)
	fmt.Printf("%T: %v\n", a, a)
	fmt.Printf("%T: %v\n", b, b)
	fmt.Printf("%T: %v\n", c, c)
	fmt.Printf("%T: %v\n", d, d)
	fmt.Printf("%T: %v\n", e, e)
	fmt.Printf("%T: %#x\n", e, e)
}
