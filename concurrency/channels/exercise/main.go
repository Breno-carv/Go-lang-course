package main

import (
	"fmt"
)

func waitForDbs(numDBs int, dbChan chan struct{}) {
	// complete the fn to block until it receives numDBs tokens on the dbChan

	for range numDBs {
		<-dbChan
	}
}

func getDatabasesChannel(numDBs int) chan struct{} {
	ch := make(chan struct{})
	go func() {
		for i := range numDBs {
			ch <- struct{}{}
			fmt.Printf("Database %v is online! \n", i+1)
		}
	}()
	return ch
}

func main() {
	ch := getDatabasesChannel(5)
	waitForDbs(5, ch)
}
