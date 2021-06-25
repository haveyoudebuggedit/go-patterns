package main

import (
	"fmt"
)

func generatePanic() {
	panic("Help!")
}

// This example demonstrates how to call a method that might accidentally panic.
func catchPanicDemo(run func()) {
	done := make(chan struct{})
	// Recovery only works in a gofunc.
	go func() {
		// When the run() function exists run the recovery
		defer func() {
			// Recover the panic
			if err := recover(); err != nil {
				// Handle error.
				fmt.Printf("Recovered panic: %v", err)
			}
			// Close the done channel
			close(done)
		}()

		run()
	}()
	// Wait for the function to finish
	<-done
}

func main() {
	catchPanicDemo(generatePanic)
}
