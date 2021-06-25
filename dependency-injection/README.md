# Dependency injection

Let's say you have module A that depends on module B. The simplest solution is to hard-code creating a copy of module B. However, then how do you test it? Dependency injection provides a means to "inject" this dependency, which can then be replaced for test purposes.

Note: this example is based on the [OOP example](../oop/), please read that one before this.

This example demonstrates how to do that in Go.

[See the code &raquo;](dependency-injection.go)