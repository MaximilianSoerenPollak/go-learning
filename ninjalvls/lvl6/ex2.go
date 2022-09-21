package main

import "fmt"

func main() {
	xi := []int{1, 2, 3, 4, 5}
	x := foo(xi...)
	fmt.Println("This is the sum", x)
	y := bar(xi)
	fmt.Println("Hello there this is", y)
}

func foo(x ...int) int {
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}

func bar(x []int) int {
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}
