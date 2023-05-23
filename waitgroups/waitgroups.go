// To wait for multiple goroutines to finish, we can use a wait
// group.

package main

import (
	"fmt"
	"sync"
	"time"
)

// This is the function we’ll run in every goroutine.
func worker(id int) {
	fmt.Printf("worker %d starting\n", id)

	// Sleep to simulate an expensive task.
	time.Sleep(time.Second)
	fmt.Printf("worker %d done\n", id)
}

func main() {
	// This WaitGroup is used to wait for all the goroutines launched
	// here to finish. Note: if a WaitGroup is explicitly passed into
	// functions, it should be done by pointer.
	var wg sync.WaitGroup

	// Launch several goroutines and increment the WaitGroup counter
	// for each.
	for i := 1; i <= 5; i++ {
		wg.Add(1)

		// Avoid re-use of the same i value in each goroutine closure. See
		// the FAQ for more details.
		i := i

		// Wrap the worker call in a closure that makes sure to tell the
		// WaitGroup that this worker is done. This way the worker itself
		// does not have to be aware of the concurrency primitives
		// involved in its execution.
		go func() {
			defer wg.Done()
			worker(i)
		}()
	}
	// Block until the WaitGroup counter goes back to 0; all the
	// workers notified they’re done.
	wg.Wait()

	// Note that this approach has no straightforward way to propagate
	// errors from workers. For more advanced use cases, consider
	// using the errgroup package.

}

// The order of workers starting up and finishing is likely to be
// different for each invocation.
// $ go run waitgroups/waitgroups.go
// worker 5 starting
// worker 1 starting
// worker 2 starting
// worker 3 starting
// worker 4 starting
// worker 1 done
// worker 4 done
// worker 2 done
// worker 5 done
// worker 3 done
