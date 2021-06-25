# Limited queue

Sometimes we want to parallel-process something on a limited amount of parallel processes.

This can be achieved in Go through a clever use of [channels](https://gobyexample.com/channels).

[See the code &raquo;](queue.go) 