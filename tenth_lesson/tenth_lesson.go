package main

import "fmt"

func main() {

	// about memory location:
	// After declaring a x variable, and the creating a y varible and assigning to the value of x,
	// we actually have a "copy" of X, but pointing to the different values in memory,
	// if we then change the value of x to something else, y will still point to the original value in memory, and will not be affected by the change in x.

	var x int = 10
	y := x // y is a copy of x, but both point to the different values in memory
	x = 20

	fmt.Println("x:", x) // x is now 20
	fmt.Println("y:", y) // y is still 10, because it points to the original value in memory

	// But now, if we want y to POINT to the same value in memory as X, then we should initiallize it as a pointer:

	var a int = 10
	b := &a // b is a pointer to a, it points to the same value in memory as a
	a = 20

	fmt.Println("a:", a)   // a is now 20
	fmt.Println("b:", b)   // b is a pointer to a, so it will show the address of a
	fmt.Println("*b:", *b) // *b dereferences the pointer, so it will show the value of a
}
