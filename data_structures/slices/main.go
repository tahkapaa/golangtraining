package main

import "fmt"

func main() {
	myArray := [6]int{1, 2, 3, 4, 5}
	fmt.Printf("%T\n", myArray)
	fmt.Println(myArray)
	fmt.Println(cap(myArray))

	mySlice := []int{1, 2, 3, 4, 5}
	fmt.Printf("%T\n", mySlice)
	fmt.Println(mySlice)
	fmt.Println(cap(mySlice))

	fmt.Println(cap(mySlice[1:2]))

	myString := "Tähkäpää" //[84 195 164 104 107 195 164 112 195 164 195 164]
	fmt.Println(myString)
	fmt.Println([]byte(myString))
	fmt.Println(myString[1]) // 195
	fmt.Println(myString[2]) // 164 -> ä = [195 164]

	myBytes := []byte{84, 195, 164, 104, 107, 195, 164, 112, 195, 164, 195, 164}
	fmt.Println(string(myBytes))

	ints := make([][]int, 0)
	for i := 0; i < 4; i++ {
		a := make([]int, 0)
		for j := 0; j < 4; j++ {
			a = append(a, j)
		}
		ints = append(ints, a)
	}
	fmt.Println(ints)

	fmt.Println(mySlice)
	plusOne(mySlice)
	fmt.Println(mySlice)
}

// Slice is a reference type (like map and channel)
// What is passed, is a memory address, which is copied by value
func plusOne(slice []int) {
	for i := range slice {
		slice[i]++
	}
}
