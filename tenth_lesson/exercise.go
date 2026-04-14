package main

import "fmt"

func incrementSends(sendsSoFar int, sendsToAdd int) int {
	return sendsSoFar + sendsToAdd
}

func main() {

	sendsSoFar := 430
	const sendsToAdd = 25
	sendsSoFar = incrementSends(sendsSoFar, sendsToAdd)
	fmt.Println("you've sent", sendsSoFar, "messages.")
}
