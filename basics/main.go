package main

import "fmt"

func main() {
	xi := []int{1, 2, 3, 4, 5, 6}
	x := bar(xi...)
	y := foo(xi)
	fmt.Println(x)
	fmt.Println(y)
}

func bar(x ...int) int {
	total := 0
	for _, v := range x {
		total += v

	}
	return (total)
}
func foo(x []int) int {
	total := 0
	for _, v := range x {
		total += v
	}
	return (total)
}
