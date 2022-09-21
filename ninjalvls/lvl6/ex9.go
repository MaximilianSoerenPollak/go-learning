package main

import "fmt"

func add(x, y int) int {
	return x + y
}

func multiply(x int) int {
	return 10 * x
}

func main() {
	fmt.Println(multiply(add(2, 3)))
}
