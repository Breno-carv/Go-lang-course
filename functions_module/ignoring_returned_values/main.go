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

}
