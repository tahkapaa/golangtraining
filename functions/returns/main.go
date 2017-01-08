package main

import "fmt"

func main() {
	fmt.Println(greet("Pasi", "Tähkäpää"))
}

func greet(fn, sn string) string {
	return fmt.Sprintf("Hello %s %s!", fn, sn)
}
