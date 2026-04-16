package main

// the syntax of a basic loop is this:
// for INITIAL; CONDITION; AFTER {
//     // code to execute
// }

// initial condition creates a block-scoped variable that is only accessible within the loop. It is executed once before the loop starts.

// condition is evaluated before each iteration of the loop. If it evaluates to true, the loop body is executed. If it evaluates to false, the loop terminates.

// after is executed after each iteration of the loop body. It is typically used to update the loop variable.

// in this example, each subsequent message an user sents, costs an additional 10c to send.

import (
	"errors"
	"fmt"
)

func bulkSend(numMessages int) (float64, error) {
	if numMessages <= 0 {
		return 0.0, errors.New("number of messages must be greater than zero")
	}
	totalCost := 0.0

	for i := 0; i < numMessages; i++ {
		cost := float64(i) * 0.01 // cost increases by 10c for each message

		totalCost += cost + 1.0

	}

	return totalCost, nil
}

// if we want to create the same structure to a very large number of msgs, we can ommit the condition and force the break manually

func bulkSendInfinite(numMessages int) (float64, error) {
	if numMessages <= 0 {
		return 0.0, errors.New("number of messages must be greater than zero")
	}
	totalCost := 0.0

	for i := 0; ; i++ {
		if i >= numMessages {
			break
		}
		cost := float64(i) * 0.01

		totalCost += cost + 1.0
	}

	return totalCost, nil
}

func main() {
	cost, err := bulkSend(10)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Total cost for sending 10 messages: $%.2f\n", cost)

	newCost, err := bulkSend(-1)
	if err != nil {
		fmt.Println("Error:", err)

	}
	fmt.Printf("Total cost for sending -1 messages: $%.2f\n", newCost)

	largeNumberOfMessages := 1000
	largeCost, err := bulkSendInfinite(largeNumberOfMessages)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Total cost for sending %d messages: $%.2f\n", largeNumberOfMessages, largeCost)
}
