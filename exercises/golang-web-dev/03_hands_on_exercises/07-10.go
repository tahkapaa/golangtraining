package main

import "fmt"

type transportation interface {
	transportationDevice() string
}

func report(t transportation) {
	fmt.Println(t.transportationDevice())
}

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

func (t truck) transportationDevice() string {
	return "I move things around"
}

type sedan struct {
	vehicle
	luxury bool
}

func (s sedan) transportationDevice() string {
	return "I move people around"
}

func main() {
	t := truck{vehicle{2, "red"}, false}
	fmt.Println(t)
	fmt.Println(t.color)
	fmt.Println(t.transportationDevice())
	report(t)

	s := sedan{vehicle{5, "blue"}, true}
	fmt.Println(s)
	fmt.Println(s.luxury)
	fmt.Println(s.transportationDevice())
	report(s)
}
