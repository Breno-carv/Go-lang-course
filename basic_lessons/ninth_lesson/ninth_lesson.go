package main

import "fmt"

func concat(s1 string, s2 string) string { // the problem was that both variables were not defined as strings
	return s1 + s2
}

// also, if all args are from the same type, we can omit the type for all but the last one, like this:
// func concat(s1, s2 string) string { it also works if you have more than 2 args, and some from different types, like this:
// func concat(s1, s2 string, n int) string {

func main() {

	fmt.Println(concat("Hello, ", "World!"))
	fmt.Println(concat("Go ", "Programming"))

}
