package main

import (
	"fmt"
	"time"
)

// This example will be for channels and how they operate, they block code and are used to create
// communication through my goroutines

func prepareOrder(order string, ch chan string) {
	fmt.Println("Preparing your order...")
	time.Sleep(2 * time.Second)
	ch <- "Order ready: " + order

}

func main() {
	ch := make(chan string)
	order := "Pepperoni pizza, extra large"

	go prepareOrder(order, ch)
	fmt.Println("Waiting for order...")

	result := <-ch // this blocks the main process until the data is properly received and passed to result var
	fmt.Println(result)
}
