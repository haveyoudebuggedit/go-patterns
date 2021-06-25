package main

import (
	"errors"
	"fmt"
)

// region Error struct

// Let's create a custom error type
type myErrorStruct struct {
}

func (m *myErrorStruct) Error() string {
	return "My error!"
}

func (m *myErrorStruct) SayHi() string {
	return "Hello world!"
}

// We return error here because otherwise the nil check won't work.
func generateStructError() error {
	return &myErrorStruct{}
}

func checkStructError() {
	if err := generateStructError(); err != nil {
		myErr := &myErrorStruct{}
		if errors.As(err, &myErr) {
			print(myErr.SayHi())
		} else {
			panic(myErr)
		}
	}
}

// endregion

//region Error variable

// This example demonstrates how to use a static error

var errMyVarError = fmt.Errorf("this is an error")

// We return the error stored above
func generateVarError() error {
	return errMyVarError
}

func checkVarError() {
	if err := generateVarError(); err != nil {
		if errors.Is(err, errMyVarError) {
			print("A my var error happened!")
		} else {
			// Handle error that is not a myError
		}
	}
}

// endregion

func main() {
	// Generate a struct-error
	checkStructError()
	// Generate a var-error
	checkVarError()
}
