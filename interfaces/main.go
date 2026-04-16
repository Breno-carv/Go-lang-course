package main

import (
	"fmt"
	"math"
)

// interfaces are just a collection of method signatures

type shape interface {
	area() float64
	perimeter() float64
}

type circle struct {
	radius float64
}

type rectangle struct {
	width, height float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func (r rectangle) area() float64 {
	return r.width * r.height
}

func (r rectangle) perimeter() float64 {
	return 2 * (r.width + r.height)
}

func main() {
	c := circle{radius: 5}
	r := rectangle{width: 4, height: 6}

	shapes := []shape{c, r}

	for _, s := range shapes {
		fmt.Printf("Area: %.2f\n", s.area())
		fmt.Printf("Perimeter: %.2f\n", s.perimeter())
	}

	fmt.Printf("Circle area: %.2f\n", c.area())

	// Type assertion: This is a concept that allows you to initialize a struct with an interface type.
	// It is used to check if an interface value holds a specific type and to extract the underlying value of that type.

	// it returns the underlying value of the interface if it holds the specified type, and a boolean indicating whether the assertion succeeded or not.

	s := shape(c) // type assertion
	n := shape(r)

	if circle, ok := n.(circle); ok {
		fmt.Printf("Circle radius: %.2f\n", circle.radius)
	} else {
		fmt.Println("n is not a circle")
	}

	if circle, ok := s.(circle); ok {
		fmt.Printf("Circle radius: %.2f\n", circle.radius)
	} else {
		fmt.Println("s is not a circle")
	}

	// type switches are just assertions for multiple types inside a func

	for _, s := range shapes {
		switch v := s.(type) {
		case circle:
			fmt.Printf("%+v is a circle\n", v)
			fmt.Printf("Circle with radius: %.2f\n", v.radius)
		case rectangle:
			fmt.Printf("%+v is a rectangle\n", v)
			fmt.Printf("Rectangle with width: %.2f and height: %.2f\n", v.width, v.height)
		default:
			fmt.Println("Unknown shape")
		}

	}
}
