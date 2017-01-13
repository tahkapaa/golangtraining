package main

import (
	"fmt"
	"sort"
)

type people []string

func (p people) Len() int {
	return len(p)
}

func (p people) Less(i, j int) bool {
	return p[i] < p[j]
}

func (p people) Swap(i, j int) {
	tmp := p[i]
	p[i] = p[j]
	p[j] = tmp
}

func main() {
	studyGroup := people{"Zeno", "John", "Al", "Jenny"}
	fmt.Println("Original: ", studyGroup)
	sort.Sort(studyGroup)
	fmt.Println("Sort: ", studyGroup)
	sort.Sort(sort.Reverse(studyGroup))
	fmt.Println("Reversed: ", studyGroup)

	s := []string{"Zeno", "John", "Al", "Jenny"}
	fmt.Println("Original", s)
	sort.Strings(s)
	fmt.Println("Sorted", s)
	sort.Sort(sort.Reverse(sort.StringSlice(s)))
	fmt.Println("Reversed", s)

	n := []int{7, 4, 8, 2, 9, 19, 12, 32, 3}
	fmt.Println("Original: ", n)
	sort.Ints(n)
	fmt.Println("Sorted: ", n)
	sort.Sort(sort.Reverse(sort.IntSlice(n)))
	fmt.Println("Reverse: ", n)
}
