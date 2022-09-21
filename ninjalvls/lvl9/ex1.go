package main

import (
	"fmt"
	"sync"
)

func main() {

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		fmt.Println("First go routine")
		wg.Done()
	}()
	go func() {
		fmt.Println("Second")
		wg.Done()
	}()
	wg.Wait()
}
