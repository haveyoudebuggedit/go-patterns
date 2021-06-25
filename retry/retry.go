package main

import (
	"context"
	"fmt"
	"time"
)

func someNetworkBoundCall() error {
	return fmt.Errorf("network connection aborted")
}

// This example demonstrates how to do a retry loop.
func retryDemo(ctx context.Context) (lastError error) {
	// We define a label to break to when we are done
loop:
	for {
		println("Trying with hard-coded demo...")
		// We call someNetworkBoundCall() that can fail. Ideally, this call
		// also takes a context (ctx) as a parameter so it can't hang indefinitely.
		if lastError = someNetworkBoundCall(); lastError == nil {
			// No error, so let's return
			return nil
		}
		select {
		case <-ctx.Done():
			// The context has expired, break the loop and fail permanently.
			break loop
		case <-time.After(10 * time.Second):
		}
	}
	return fmt.Errorf("failed to call someNetworkBoundCall repeatedly, giving up (%w)", lastError)
}

// This pattern was suggested by @bwplotka. It demonstrates how to create a reusable retrier that executes the "what" function.
// It is somewhat limited because it cannot return values, but it is a single implementation for retrying.

func retry(forHowLong context.Context, what func() error, howOften time.Duration) (err error) {
	for {
		println("Trying with callback demo...")
		if err = what(); err == nil {
			return nil
		}
		select {
		case <-forHowLong.Done():
			return forHowLong.Err()
		case <-time.After(howOften):
		}
	}
}

func main() {
	// Demo 1, with hard-coded call
	ctx, cancelFunc := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancelFunc()
	if err := retryDemo(ctx); err != nil {
		fmt.Printf("%v\n", err)
	}

	// Demo 2 with callbacks
	ctx2, cancelFunc2 := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancelFunc2()
	if err := retry(ctx2, someNetworkBoundCall, 10*time.Second); err != nil {
		fmt.Printf("%v\n", err)
	}
}
