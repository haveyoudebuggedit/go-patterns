# Contexts

When writing an application that needs to shut down in an orderly fashion the signal to shut down needs to be carried across module boundaries. Thankfully, Go provides a mechanism called [contexts](https://golang.org/pkg/context/) to achieve this goal. Network-related functions often accept a context variable that, if canceled, causes the network function to abort.

This pattern demonstrates how to properly set up a timeout context to abort the execution of a function when the timeout expires.

[See the code &raquo;](context.go)