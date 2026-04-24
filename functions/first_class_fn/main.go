/*
Higher order fns are just functions that receive another function as a parameter, or returns a fn as its return value;

First class fns, on the other hand, are just functions that are used as "data"
*/

package main

import "fmt"

// first easy scenarios:

func sum(a, b int) int {
	return a + b
}

func multiply(a, b int) int {
	return a * b
}

func aggregate(a, b, c int, arithmetic func(int, int) int) int {
	return arithmetic(arithmetic(a, b), c)
}

func main() {
	a, b, c := 10, 20, 5

	fmt.Println(aggregate(a, b, c, multiply))
}
