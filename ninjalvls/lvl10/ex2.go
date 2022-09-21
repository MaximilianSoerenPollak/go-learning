package main

import (
	"fmt"
)

func main() {
	// Make this code run
	// cs := make(chan<- int)
	//
	// go func() {
	// 	cs <- 42
	// }()
	// fmt.Println(<-cs)
	//
	// fmt.Printf("------\n")
	// fmt.Printf("cs\t%T\n", cs)

	cs := make(chan int)

	go func() {
		cs <- 42
	}()
	fmt.Println(<-cs)

	fmt.Printf("------\n")
	fmt.Printf("cs\t%T\n", cs)

}
