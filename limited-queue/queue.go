package main

import (
	"sync"
	"time"
)

// This example demonstrates how to do queue that is processed in a limited number of parallel goroutines
func parallelProcessDemo(items []string, parallelism int) {
	wg := &sync.WaitGroup{}
	// Create a semaphore that can contain only the number of items specified in "parallelism"
	sem := make(chan struct{}, parallelism)
	// Add all items to wait for
	wg.Add(len(items))
	for _, item := range items {
		// Capture item in loop
		i := item
		// Run processing in background
		go func() {
			// This call will block if the semaphore is full.
			sem <- struct{}{}
			defer func() {
				// Release one slot in the sem
				<-sem
				// Signal that the item is done.
				wg.Done()
			}()

			// process item
			print(i)
			time.Sleep(time.Second)
		}()
	}
	// Wait for all items to finish
	wg.Wait()
}

func main() {
	parallelProcessDemo([]string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}, 3)
}
