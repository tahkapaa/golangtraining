package main

import "fmt"

// Create a new type called “gator”. The underlying type of “gator” is an int. Using var, declare an identifier “g1” as type gator (var g1 gator). Assign a value to “g1”. Print out “g1”. Print the type of “g1” using fmt.Printf(“%T\n”, g1)
// solution
// Adding onto this code: Using var, declare an identifier “x” as type int (var x int). Print out “x”. Print the type of “x” using fmt.Printf(“%T\n”, x)
// solution
// Adding onto this code: Can you assign “g1” to “x”? Why or why not?
// solution
// Adding onto this code: You will now learn about CONVERSION. This is called “CASTING” in a lot of other languages. Since “g1” is of type “gator” but its underlying type is an “int”, we can use “CONVERSION” to convert the value to an int. Here is how you do it:
// solution
// Now add a method to type gator with this signature ...
// 	greeting()
// … and have it print “Hello, I am a gator”. Create a value of type gator. Call the greeting() method from that value.
// solution

type gator int

func (g gator) greeting() {
	fmt.Println("Hello, I am a gator")
}

type flamingo bool

func (f flamingo) greeting() {
	fmt.Println("Hello, I am pink and beautiful and wonderful")
}

type swampCreature interface {
	greeting()
}

func bayou(sC swampCreature) {
	sC.greeting()
}

func main() {
	var g1 gator
	g1 = 1
	fmt.Println(g1)
	fmt.Printf("%T\n", g1)
	g1.greeting()

	var x int
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = int(g1)

	var f1 flamingo
	bayou(f1)
	bayou(g1)

}
