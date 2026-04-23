package main

import (
	"fmt"
	"unicode/utf8"
)

// complete the getUserMap function, it takes a slice of names and a slice of phone numbers
// and returns a map where the keys are the names and the values are the phone numbers and potentially an error.

type user struct {
	name        string
	phoneNumber int
}

// the first name in the slice matches the first phoneNumber and so on..

func getUserMap(names []string, phoneNum []int) map[string]int {
	if len(names) != len(phoneNum) {
		fmt.Println("Error: The length of names and phone numbers must be the same.")
		return nil
	}

	userMap := make(map[string]int)

	for i := range names {
		userMap[names[i]] = phoneNum[i]
	}
	return userMap
}

func main() {
	names := []string{"Alice", "Bob", "Charlie"}
	phoneNumbers := []int{1234567890, 9876543210, 5555555555}

	userMap := getUserMap(names, phoneNumbers)
	fmt.Println(userMap)

	newNames := []string{"Bob", "Brian", "Brian", "Charlie", "Doug", "Charlie", "Diego", "Adrian"}

	fmt.Println(getNameCounts(newNames))
}

// syntaxes for manipulating maps
// create a map:
/*
myMap := make(map[string]int) // creates an empty map with string keys and int values
or myMap := map[string]int{ // creates a map with initial values
	"Alice": 1234567890,
	"Bob": 9876543210,
}

inserting elements:
myMap["Charlie"] = 5555555555 // adds a new key-value pair to the map

retrieving values:
phoneNumber := myMap["Alice"] // retrieves the value associated with the key "Alice"

deleting
delete(myMap, "Bob") // removes the key "Bob" and its associated value from the map

checking for existence:
value, exists := myMap["Charlie"] // checks if the key "Charlie" exists in the map and retrieves its value if it does

*/

func getCounts(userIDs []string) map[string]int {

	counts := make(map[string]int)

	for _, userID := range userIDs {
		count := counts[userID]
		count++
		counts[userID] = count
	}
	return counts
}

/*
New assignment:
complete the getNameCounts fn so that it returns a map[rune]map[string]int, the idea is that we create a list by
the first letter of each name, so, in the b key we should have a map with all the names associated with the b letter
(the ones that start with B) and then the number of messages that each one has sent

e.g.: names := []string{"billy","billy","joe","bob"}
should return:
b : {
billy: 2,
bob: 1
},
j:{
joe: 1
}
*/

func getNameCounts(names []string) map[rune]map[string]int {
	finalMap := make(map[rune]map[string]int)

	for _, name := range names {
		firstLetter, _ := utf8.DecodeRuneInString(name)

		if finalMap[firstLetter] == nil {
			finalMap[firstLetter] = make(map[string]int)
		}

		finalMap[firstLetter][name]++
	}
	return finalMap
}
