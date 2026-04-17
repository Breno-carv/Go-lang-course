package main

import "fmt"

func getMessageWithRetries() [3]string { // That's a function that returns an array of strings with a fixed size of 3. The array contains three different messages that can be used for sending to users.
	return [3]string{
		"Click here to sign up!",
		"Please, click here!",
		"We're begging you to sign up!",
	}

}

// the difference between an array and a slice is that an array has a fixed size, while a slice is a dynamically sized, flexible view into the elements of an array.
//  Slices are more commonly used in Go because they provide more functionality and are easier to work with than arrays.

// there's a very interesting syntax on "slicing" arrays that already exists, you can create a "subarray"
// (but it's actually a slice) by using the syntax array[start(inclusive):end(exclusive)].

// exempli gratia:

var arrayA = [5]int{1, 2, 3, 4, 5}
var sliceFromA = arrayA[1:4] // this will create a slice that contains the elements from index 1 to index 3 of the arrayA, which are 2, 3, and 4.

func send(name string, doneAt int) {
	fmt.Printf("Message sent to %s at %d\n", name, doneAt)

	messages := getMessageWithRetries()
	for i := 0; i < len(messages); i++ {
		msg := messages[i]
		fmt.Printf("Processing message: %s\n", msg)

	}
}

func main() {
	firstUser := User{Name: "Alice", HasPlan: true, Plan: freePlan}
	secondUser := User{Name: "Bob", HasPlan: false}
	thirdUser := User{Name: "Charlie", HasPlan: true, Plan: proPlan}

	firstUserMessages, err := getMessagesForUser(firstUser)
	if err != nil {
		fmt.Printf("Error getting messages for user %s: %v\n", firstUser.Name, err)
	}
	secondUserMessages, err := getMessagesForUser(secondUser)
	if err != nil {
		fmt.Printf("Error getting messages for user %s: %v\n", secondUser.Name, err)
	}
	thirdUserMessages, err := getMessagesForUser(thirdUser)
	if err != nil {
		fmt.Printf("Error getting messages for user %s: %v\n", thirdUser.Name, err)
	}

	fmt.Printf("Messages for %s: %v\n", firstUser.Name, firstUserMessages)
	fmt.Printf("Messages for %s: %v\n", secondUser.Name, secondUserMessages)
	fmt.Printf("Messages for %s: %v\n", thirdUser.Name, thirdUserMessages)

}

// exercise:
// if the user has a freePlan, return only the first string of the array
// if the user has no plan, return the whole array
// if the user has the pro plan, return only the last 2 strings of the array

type User struct {
	Name    string
	HasPlan bool
	Plan    string
}

const (
	freePlan = "free"
	proPlan  = "pro"
)

func verifyPlan(user User) (string, error) {
	if user.HasPlan {
		return user.Plan, nil
	}
	return "", fmt.Errorf("user %s does not have a plan", user.Name)
}

func getMessagesForUser(user User) ([]string, error) {
	messages := getMessageWithRetries()

	plan, err := verifyPlan(user)
	if err != nil {
		return nil, fmt.Errorf("error verifying plan for user %s: %v", user.Name, err)
	}
	switch plan {
	case freePlan:
		return messages[:1], nil // return the first string of the array
	case proPlan:
		return messages[len(messages)-2:], nil // return the last 2 strings of the array
	default:
		return messages[:], nil // return the whole array
	}
}
