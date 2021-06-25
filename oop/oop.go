package main

// This example demonstrates how to do OOP in go

// Car is the public interface that others will be able to use
type Car interface {
	// We force the implementation to implement the Start() and Drive() methods.
	Start()
	Drive()
}

// AbstractCar is a struct that partially implements Car. It is exported because we want others to embed it.
type AbstractCar struct {
}

// Start is an implementation in the abstract class. Note that Drive() is not implemented, so any
// inheriting classes like *car below will need to implement it.
func (a *AbstractCar) Start() {
	println("Car now starting")
}

// NewCar the constructor. We return the Car interface instead of the *car implementation
// because we want to force the implementation to be correct.
func NewCar(param1 string, param2 int) (Car, error) {
	return &car{
		param1: param1,
		param2: param2,
	}, nil
}

// This is the internal storage of member variables.
type car struct {
	// Embed AbstractCar for inheritance.
	AbstractCar

	param1 string
	param2 int
}

// This is the implementation of the Drive method
func (c *car) Drive() {
	// You can use "c" here as you would use "this" in other languages.
	print(c.param1)
}

func main() {
	car, err := NewCar("Wrrrrrrrr", 5)
	if err != nil {
		panic(err)
	}
	car.Start()
	car.Drive()
}
