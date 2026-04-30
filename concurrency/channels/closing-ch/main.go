// It is possible to close a ch by using the built-in fn close(), this will always be done by
// the sending side of the ch

// chs also return the ok parameter to verify wether it's still open intCh, ok := <- someCh

package main

import (
	"fmt"
	"time"
)

func emailSender(ch chan string) {
	// TODO:
	// Loop over the channel and process messages
	// Print: "Sending email: <msg>"
	// Simulate delay with time.Sleep
	// When channel closes, print: "All emails processed"
	for msg := range ch {
		fmt.Printf("Sending email: %s \n", msg)
		time.Sleep(125 * time.Millisecond)
	}
	fmt.Println("All emails processed!")
}

func main() {
	messages := []string{
		"Order confirmed",
		"Payment received",
		"Order shipped",
		"Order delivered",
	}

	ch := make(chan string)
	// Start consumer
	go emailSender(ch)

	// TODO:
	// Send all messages into the channel
	for _, msg := range messages {
		ch <- msg
	}

	// TODO:
	// Close the channel
	close(ch)

	// Prevent main from exiting too early
	time.Sleep(3 * time.Second)
}
