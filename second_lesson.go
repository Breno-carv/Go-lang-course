package main

import "fmt"

func main() {

	messagesFromDoris := []string{
		"You doing anything later??",
		"Did you get my last message?",
		"Don't leave my hanging...",
		"Please respond I'm lonely...",
	} // wtf are those example msgs tho... lmao

	numMessages := float64(len(messagesFromDoris))
	costPerMessage := .02

	// assume everything above is correct

	totalCost := costPerMessage * numMessages // the problem was in this line, it should not be + but * operation

	fmt.Printf("Doris spent %.2f on text messages today. \n", totalCost)

}
