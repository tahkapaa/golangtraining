package main

import "fmt"

func main() {
	m := make(map[byte]string)
	fillMap(m)

	value, ok := m[255]
	fmt.Printf("Value: %v - Exists: %t\n", value, ok)
	value, ok = m[105]
	fmt.Printf("Value: %v - Exists: %t\n", value, ok)
}

func fillMap(byteMap map[byte]string) {
	for i := byte(0); i < 255; i++ {
		byteMap[i] = string(i)
	}
}
