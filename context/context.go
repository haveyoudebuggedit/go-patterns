package main

import (
	"context"
	"time"
)

// someNetworkCall simulates a network call. It will wait for 20 seconds, or until the
func someNetworkCall(ctx context.Context) error {
	select {
	case <-time.After(20 * time.Second):
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// This example demonstrates how to create a context to pass on to network functions
func main() {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 10*time.Second)
	// Call the cancelFunc at the end of this method so we don't leave a goroutine hanging
	defer cancelFunc()
	// Call the function that needs a context.
	if err := someNetworkCall(ctx); err != nil {
		panic(err)
	}
}
