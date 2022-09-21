package main

import "fmt"

func main() {
	x := func() string {
		return "This is an anon func"
	}
	fmt.Println(x())
}
