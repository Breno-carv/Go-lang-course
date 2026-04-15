package main

import "fmt"

func main() {
	//conditional statements: if, else if, else
	// go don't use parentheses for the condition, but it does use curly braces for the block of code

	messageLen := 10
	maxMessageLen := 20
	fmt.Println("Trying to send a message of length:", messageLen)

	if messageLen > maxMessageLen {
		fmt.Println("Message not sent")
	} else {
		fmt.Println("Message sent")
	}

	// So, when dealing with a "one time use" variable, we can declare it
	// "inside" the if statement, this will keep the variable only accessible within the scope of the if statement

	if msg := "Hello, World!"; len(msg) > maxMessageLen {
		fmt.Println("Message not sent")
	} else {
		fmt.Println("Message sent:", msg)
	}
}
