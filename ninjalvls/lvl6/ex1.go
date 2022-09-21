package main

import "fmt"

func main() {
	x := foo()
	y, z := bar()
	fmt.Println("This is x", x)
	fmt.Println("\n And this is y,z", y, z)
}

func foo() int {
	x := 5
	return x
}

func bar() (int, string) {
	x := 5
	y := "This is a string"
	return x, y
}
