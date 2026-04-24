// the idea of a closure is that you initialize a variable in the high level on a high order fn so that you can always "keep"
// track of it in memory:

package main

import "fmt"

func concatter() func(string) string {
	doc := ""

	return func(word string) string {
		doc += word + " "
		return doc
	}
}

func main() {
	phrase := concatter()

	fmt.Println(phrase("Hi"))
	fmt.Println(phrase("There!"))
}
