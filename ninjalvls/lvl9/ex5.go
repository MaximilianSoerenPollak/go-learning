// Re-write of the ex3 code with atomic this time.
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {

	var wg sync.WaitGroup
	var counter int64
	wg.Add(100)
	// To lock the variables
	for v := 0; v < 100; v++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			fmt.Println(atomic.LoadInt64(&counter)) // We should be reading the value (via the print) before we unlock it.
			wg.Done()
		}()

	}
	wg.Wait()
}
