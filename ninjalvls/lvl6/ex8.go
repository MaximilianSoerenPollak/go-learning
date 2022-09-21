package main

import "fmt"

func main() {
	x := rf()
	fmt.Println(x())
}

func rf() func() string {
	return func() string {
		return "This is a return func string"
	}
}
