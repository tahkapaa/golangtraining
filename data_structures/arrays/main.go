package main

import "fmt"

func main() {
	var strings [10]string
	fmt.Printf("strings[1]: '%s'\n", strings[1])
	fmt.Println(strings)

	for i := range strings {
		strings[i] = fmt.Sprintf("string: '%d'", i)
	}
	fmt.Printf("string: '%s'\n", strings[1])
	fmt.Println(strings)

	/*var ints [256]int

	for i := 0; i < 256; i++ {
		ints[i] = i
	}

	fmt.Printf("%v - %T - %b\n", ints, ints, ints)*/

	var bytes [256]byte

	for i := 0; i < 256; i++ {
		bytes[i] = uint8(i) //byte(i)
	}

	for _, value := range bytes {
		fmt.Printf("%v - %T - %b\n", value, value, value)
	}
}
