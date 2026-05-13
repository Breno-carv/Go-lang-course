// generics are, as in other languages, a "anything" Type, it's useful when a certain function
// is designed to treat multiple dataTypes at once, eg:

package main

import "fmt"

func splitAnySlice[T any](slice []T) ([]T, []T) {
	mid := len(slice) / 2

	return slice[:mid], slice[mid:]
}

// for example, let's say that i only want to deal with dataTypes that have the String() method callable, so, i can create a custom
// interface to use it as a "generic" type:

type stringer interface {
	String() string
}

func concat[T stringer](vals []T) string {
	result := ""

	for _, val := range vals {
		result += val.String()
	}

	return result
}

func main() {
	intSlice := []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}
	stringSlice := []string{"Bob", "Alice", "Ana", "Breno", "Carlos", "Douglas"}

	fmt.Println(intSlice)
	fmt.Println(splitAnySlice(intSlice))

	fmt.Println(stringSlice)
	fmt.Println(splitAnySlice(stringSlice))

}
