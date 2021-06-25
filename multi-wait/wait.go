package main

import (
	"sync"
	"time"
)

// This example demonstrates how to create a method where multiple goroutines wait for a single event to finish
func main() {
	// Create a channel that will be closed when the main go func is done
	done := make(chan struct{})
	wg := &sync.WaitGroup{}
	wg.Add(2)
	go func() {
		time.Sleep(10 * time.Second)
		// Close the channel to trigger the other goroutines to continue
		close(done)
	}()
	go func() {
		// Wait for done to be closed. Optionally, you can add waiting for a timeout or a context to expire.
		select {
		case <-done:
			print("Go func A is now running!")
		}
		wg.Done()
	}()
	go func() {
		// Wait for done to be closed. Optionally, you can add waiting for a timeout or a context to expire.
		select {
		case <-done:
			print("Go func B is now running!")
		}
		wg.Done()
	}()
	// Wait for both processes to complete
	wg.Wait()
}
