# Callbacks

Go treats functions as first class objects: they can be passed from function to function and then called.

In this example we demonstrate how this happens: the `GetEnv` function is passed to the `PrintFOO` method as a variable.

This allows for injecting a custom `GetEnv` for testing purposes.

[See the code &raquo;](callback.go)
