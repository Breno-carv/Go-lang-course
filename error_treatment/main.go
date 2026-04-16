package main

import (
	"fmt"
)

// a priori, the error handling in Go is done by a simple interface called Error that has only one method:
// type error interface { Error () string }

// there are two main values that the Error() string could return: nil and not nil
// Basically, everything that's not nil is an error

// exercise below:
func costPerSMSToCouple(msgToCustomer, msgToSpouse string) (float64, error) {
	costToCustomer, err := costPerSMS(msgToCustomer)
	if err != nil {
		return 0.0, fmt.Errorf("Could not send a message to the customer: %v", err)
	}
	costToSpouse, err := costPerSMS(msgToSpouse)
	if err != nil {
		return 0.0, fmt.Errorf("Could not send the message to the spouse: %v", err)
	}

	totalValue := costToCustomer + costToSpouse
	return totalValue, nil
}

func costPerSMS(message string) (float64, error) {
	const maxTextLen int = 125
	const costPerChar float64 = 0.0002
	if len(message) > maxTextLen {
		return 0.0, fmt.Errorf("Text should have %d characters length", maxTextLen)
	}

	totalValue := costPerChar * float64(len(message))
	return totalValue, nil
}

func main() {
	msgToCustomer := "Hello, this is a message to the customer."
	msgToSpouse := "Hello, this is a message to the spouse."

	totalCost, err := costPerSMSToCouple(msgToCustomer, msgToSpouse)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Total cost of sending SMS: $%.4f\n", totalCost)

}
