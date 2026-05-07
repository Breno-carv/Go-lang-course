// mutexes are used to lock access to data.
// This ensures that we can control which goroutine can have access to certain data at a certain time.

/*
e.g.:
func protected () {
mux.Lock() >> blocks the data until the fn is done
defer mux.Unlock() >> unblocks the data when the fn is done
}

A good usage for this kind of locking system is with MAPS, maps are no thread-safe in go
*/

package main

import (
	"fmt"
	"sync"
)

/*	counter := 0
	var wg sync.WaitGroup
	var mux sync.Mutex

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			mux.Lock()
			defer mux.Unlock()
			defer wg.Done()
			counter++ // race condition
		}()

	}
	wg.Wait()
	fmt.Println("Final Counter:", counter)

*/

// new example:
type Bank struct {
	balance float64
	rw      sync.RWMutex
}

func (b *Bank) Deposit(amount float64) {
	b.rw.Lock()

	b.balance += amount
	// read it inside the lock, so the value is held, if trying to read b.balance outside the lock, another goroutine can change it in the meantime
	b.rw.Unlock()
}

func (b *Bank) Balance() float64 {
	b.rw.RLock()
	defer b.rw.RUnlock()

	return b.balance
}

func (b *Bank) Withdraw(amount float64) {
	b.rw.Lock()

	b.balance = b.balance - amount
	b.rw.Unlock()
}

func main() {
	Inter := Bank{
		balance: 2552.23,
	}
	var wg sync.WaitGroup

	Inter.Balance()

	for range 500 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			Inter.Deposit(1000)
			Inter.Withdraw(100)
		}()
	}

	wg.Wait()

	fmt.Printf("Your new total is: %.2f", Inter.Balance())
}
