package main

import (
	"fmt"
	"strings"
)

func getTrimmedName(s string) (string, string) {
	return strings.Split(s, " ")[0], strings.Split(s, " ")[1]
}
func main() {

	// Since go doesn't allow you to keep unused variables
	// you can simply ignore some variables by using the blank identifier "_"
	// in this case in particular, the compiler will indeed IGNORE the variable

	firstName, _ := getTrimmedName("John Doe") //ignore lastName
	fmt.Println(firstName)
	_, lastName := getTrimmedName("John Doe") //ignore firstName
	fmt.Println(lastName)

	//something interesting about fns is that you can initialize a variable inside the function itself,
	// also, you can simply return (nothing) and it will automatically return the variable that you initialized inside the function, for example:
	// it's called implicit (or naked) returns, honestly, should not be used that much

	fmt.Println(divideTwoNumbers(10, 2)) // returns 5, nil
	fmt.Println(divideTwoNumbers(10, 0)) // returns 0, error
}

// Below we're going to work a little bit on this concept called EARLY RETURNS
// so, basically, what this means is that we can return a value before the actual return statement inside a fn

func divideTwoNumbers(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}
	return a / b, nil
}
