package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	var wg sync.WaitGroup
	wg.Add(100)
	// To lock the variables
	var m sync.Mutex

	counter := 0
	for v := 0; v < 100; v++ {
		go func() {
			m.Lock()
			i := counter
			i++
			counter = i
			fmt.Println(counter) // We should be reading the value (via the print) before we unlock it.
			// otherwise we have issues that we get stuff out of order (obv.)
			m.Unlock()
			wg.Done()
		}()

	}
	wg.Wait()
	fmt.Println("CPU's", runtime.NumCPU)
}
