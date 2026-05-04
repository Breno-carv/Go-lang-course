package main

// the SELECT keyword is actually used on a goroutine that needs to handle multiple ch's simultaneously

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Println("Worker started...")

	time.Sleep(time.Second * 3)
	done <- true
}

func main() {
	done := make(chan bool)

	go worker(done)

	select {
	case <-done:
		fmt.Println("Worker has finished successfully")
	case <-time.After(4 * time.Second): // so if this time is longer than what it takes to execute the routine, nothing wrong happens...
		fmt.Println("The job has been canceled due to overtime...")
	}

	// 2nd example with multiple ch's
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(time.Second * 1)
		ch1 <- "Message from channel 1"
		close(ch1)
	}()
	go func() {
		time.Sleep(time.Second * 2)
		ch2 <- "Message from channel 2"
		close(ch2)
	}()

	for ch1 != nil || ch2 != nil {
		select {
		case msg1, ok := <-ch1:
			if !ok {
				ch1 = nil
				break
			}
			fmt.Println(msg1)
		case msg2, ok := <-ch2:
			if !ok {
				ch2 = nil
				break
			}
			fmt.Println(msg2)

		}
	}

}
