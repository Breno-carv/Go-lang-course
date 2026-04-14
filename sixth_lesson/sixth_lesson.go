package main

import "fmt"

func main() {
	// defining consts -> CONST should always be assigned before compilling

	const firstName = "Carlos"
	var typeOfPlan = []string{"Basic", "Plus"}

	isPlus := true

	plan := typeOfPlan[map[bool]int{true: 1, false: 0}[isPlus]]

	fmt.Printf("%s has a %s plan right now.\n", firstName, plan)

}
