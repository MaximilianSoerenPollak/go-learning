package main

import (
	"fmt"
)

func main() {
	// Pull channel values off using select statements.
	// ????????????
	q := make(chan int)
	c := gen(q)

	receive(c, q)

	fmt.Println("about to exit")
}

func gen(q chan<- int) <-chan int {
	c := make(chan int)
	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		q <- 1   // Do not understand yet fully why we do this.
		close(c) // Need to close channel here so it wont get blocked.
	}()

	return c
}

func receive(c, q <-chan int) {
	for {
		select {
		case v := <-c: // Interesting that here you just do ':' not '{}'. Remember this.
			fmt.Println("This is in channel c", v)
		case <-q:
			return
		}
	}
}
