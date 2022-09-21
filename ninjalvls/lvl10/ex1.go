package main

import (
	"fmt"
)

func main() {
	// This is how you can make it work without buffered channels. Just use an anon func.
	// c := make(chan int)
	//
	// go func() {
	// 	c <- 42
	// }()
	//
	// fmt.Println(<-c)

	// With buffered channels you do not need to do this (although they are not as wanted as unbuffered ones).
	c := make(chan int, 1)

	c <- 42

	fmt.Println(<-c)

}
