package main

import "fmt"

func main() {
	fmt.Println(42)
	fmt.Printf("%d\t%b\n", 42, 42)          // decimal, binary
	fmt.Printf("%d\t%b\t%#x\n", 50, 50, 50) // decimal, binary, hex
	fmt.Printf("%d\t%b\t%#X\n", 50, 50, 50) // decimal, binary, hex

	for index := 0; index < 200; index++ {
		fmt.Printf("%d\t%b\t%#x\t%q\n", index, index, index, index)
	}
}
