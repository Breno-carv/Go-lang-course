package main

import "fmt"

// the fmt.Sprintf always returns a string value, so, it can be used in memory as a normal var

func formatUserInfo(username string, age int) string {

	return fmt.Sprintf("%s has %d years old.", username, age)

}

type divideError struct {
	dividend float64
}

func (e divideError) Error() string {
	return fmt.Sprintf("cannot divide by zero: %v", e.dividend)
}

func divide(dividend, divisor float64) (float64, error) {
	if divisor == 0 {
		return 0, error(divideError{dividend: dividend})
	}

	return dividend / divisor, nil
}

func main() {
	const username string = "Alice"
	age := 28

	formattedInfo := formatUserInfo(username, age)

	fmt.Println(formattedInfo)

	dividend := 0.0
	divisor := 0.0

	division, err := divide(dividend, divisor)
	if err != nil {
		fmt.Println(fmt.Errorf("error occurred while dividing: %w", err))
		return
	}
	fmt.Println(division)
}
