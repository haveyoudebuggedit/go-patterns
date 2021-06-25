# Typed error handling

[Error handling](https://blog.golang.org/error-handling-and-go) in Go always returns the `error` type. There are design decisions behind this and you should always return the generic `error` type, otherwise nil-checks won't work.

This example demonstrates how you can get a typed error from the Go errors.

[See the code &raquo;](errors.go)