package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Please enter large and small number")

	var a, b int

	_, err := fmt.Scanf("%d %d", &a, &b)
	if err != nil {
		fmt.Println("Error reading input: ", err.Error())
		os.Exit(1)
	}

	if a > b {
		fmt.Printf("Remainder for %d/%d: %d\n", a, b, a%b)
	} else if b > a {
		fmt.Printf("Remainder for %d/%d: %d\n", b, a, b%a)
	} else {
		fmt.Println("Numbers are equal")
	}
}
