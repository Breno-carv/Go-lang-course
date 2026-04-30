package main

// the main idea behind pointers is that we can actually change a value without having to access it directly
// in other words, we can pass values into functions without creating copies for that value or having to return and storing a new value

// syntax:
// the * creates a pointer TYPE, so var p *int creates a blank pointer that point to an int.
// the default null value for a pointer is nil

// myStr := "Hello there!"
// myStrPtr := &myStr creates a pointer in memory to that myStr string

// the * operator dereferences a pointer, so that you can read it's value (the value that it points to), instead of it's
// address in memo. Also, it can be used to reassign the underlying value:

import (
	"fmt"
	"math"
)

type circle struct {
	radius        int
	circumference int
}

func (c *circle) SetRadius(r int) { // you need to pass the pointer reference in order to change the real value
	c.radius = r
}

func (c circle) getCircleArea() float64 {
	return float64(c.radius*c.radius) * math.Pi
}

func main() {
	myStr := "Hello, there!"
	myStrPtr := &myStr

	fmt.Println(*myStrPtr)
	*myStrPtr = "Hello, just hello..."
	fmt.Println(*myStrPtr, myStr)

	newCircle := circle{
		radius:        5,
		circumference: 10,
	}

	fmt.Printf("%.3f it's the circle current area, with r = 5", newCircle.getCircleArea())

	newCircle.SetRadius(20)
	fmt.Printf("%.3f it's the circle current area, with r = 20", newCircle.getCircleArea())

}
