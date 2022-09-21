package main

import (
	"fmt"
)

func main() {
	//Show the ok idiom.
	c := make(chan int)

	go func() {
		c <- 42
	}()

	v, ok := <-c //ok == true when the channel is still Open
	fmt.Println(v, ok)

	close(c)

	v, ok = <-c // ok == false when the channel is closed.
	fmt.Println(v, ok)
}
