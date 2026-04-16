package main

import (
	"fmt"
	"math"
)

// the exercise: As an easteregg, we decided to reward users with a free text message if they
// send a prime number of messages:
// complete de printPrime fn: it should print all primes until it reaches the input value
// it should skip non-prime numbers

func checkPrime(inputNumber int) bool {
	if inputNumber < 2 {
		return false
	} else if inputNumber <= 3 {
		fmt.Println(inputNumber)
		return true
	} else if inputNumber%2 == 0 || inputNumber%3 == 0 {
		return false
	}

	limit := int(math.Sqrt(float64(inputNumber)))

	for i := 5; i <= inputNumber; i += 6 {
		num1 := i
		num2 := i + 2

		if num1 <= limit && inputNumber%num1 == 0 {
			return false
		}
		if num2 <= limit && inputNumber%num2 == 0 {
			return false
		}

	}
	fmt.Printf("%d is a prime number.\n", inputNumber)
	return true

}

func main() {
	const inputNumber = 71

	if checkPrime(inputNumber) {
		fmt.Printf("Congratulations! You've sent a prime number of messages: %d\n", inputNumber)
	} else {
		fmt.Printf("Sorry, %d is not a prime number. Try again!\n", inputNumber)
	}
}
