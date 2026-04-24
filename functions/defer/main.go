package main

/*
defer is just a keyword that signals that the code after it should run before the function exits
*/

import "fmt"

func mockDeferExample() {
	fmt.Println("Start function")

	defer fmt.Println("Deferred 1 (runs last)")
	defer fmt.Println("Deferred 2")
	defer fmt.Println("Deferred 3 (runs first)")

	fmt.Println("Doing some work...")
	fmt.Println("End function")
}

func main() {
	mockDeferExample()
}
