package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Please enter you name:")

	var s string
	_, err := fmt.Scanln(&s)
	if err != nil {
		fmt.Println("Error reading input: ", err.Error())
		os.Exit(1)
	}

	fmt.Printf("Hello %s!\n", s)
}
