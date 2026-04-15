package main

import "fmt"

func main() {
	centsPerText := 2.00

	fmt.Printf("The type of centsPerText is %T\n", centsPerText)

	// infering types of complex numbers
	c := 0.867 + 0.5i

	fmt.Printf("The type of the c variable is %T\n", c)
}
