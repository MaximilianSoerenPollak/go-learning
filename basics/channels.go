package main

import (
	"fmt"
)

func main() {
	// c := make(chan int)
    // This passes the batton just normaly. We have an unbuffered channel and we pick the value of from the go routine
	// go func() {
	// 	c <- 42
	// }()
	// fmt.Println(<-c)
    c := make(chan int, 1) // This channel has now a buffer with size 1. It can hold 1 value no matter if it's pulled or not.
    c <- 42
    fmt.Println(<-c)

    //Note: Try to use unbuffered channels. And code in a way that makes your programm always able to deal with the things.
}
