package main

import "fmt"

func main() {
	count := 5
	for i := 0; i < count; i++ {
		for j := 0; j < count; j++ {
			fmt.Println(i, " - ", j)
		}
	}

	i := 0
	for {
		i++
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
		if i >= 50 {
			break
		}
	}
}
