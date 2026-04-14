package main

import "fmt"

func main() {

	// Interpolate the default representation:
	// There are some formatting verbs such as : %v, %T, %d, %f, %s, %t, etc.

	// %v is the default format for the value
	fmt.Printf("The default format for 100 is: %v\n", 100)

	// %T is the type of the value
	fmt.Printf("The type of 100 is: %T\n", 100)

	// %d is for integers (decimal formatting)
	fmt.Printf("The integer format for 100 is: %d\n", 100)

	// %f is for floating-point numbers
	fmt.Printf("The float format for 3.14 is: %f\n", 3.14)
	// we can also specify the precision for floating-point numbers
	fmt.Printf("The float format for 3.14 with 2 decimal places is: %.2f\n", 3.14)

	// %s is for strings
	fmt.Printf("The string format for 'Hello' is: %s\n", "Hello")

	// %t is for booleans
	fmt.Printf("The boolean format for true is: %t\n", true)

	// The exercise here is the following: Create a var that holds a dynamic string and print it out using Println

	const name = "Alice"
	const openRate = 30.5

	msg := fmt.Sprintf("%s has an open rate of %.1f%%", name, openRate)
	fmt.Println(msg)
}
