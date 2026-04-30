package main

import (
	"fmt"
	"sync"
	"time"
)

func sendEmail(msg string, wg *sync.WaitGroup) {
	// The go keyword is used to trigger a concurrent process inside a function, it cannot wait / hold it's value as
	// in a first class fn tho, it's literally used to run side processes

	go func() {
		defer wg.Done()
		time.Sleep(250 * time.Millisecond)
		fmt.Printf("E-mail received: %s \n", msg)
	}()
	fmt.Printf("Email sent: '%s'\n", msg)
}

func main() {
	var wg sync.WaitGroup
	message := "Your order is currently being processed, thanks for the purchase!"

	wg.Add(1)
	sendEmail(message, &wg)

	wg.Wait()
}
