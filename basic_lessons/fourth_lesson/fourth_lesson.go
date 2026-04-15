package main

import "fmt"

func main() {

	// numeric types class
	// the idea of unsigned numbers is quite easy to understand: it doubles the amount of bytes allocated
	// in memory to store that certain value by simply admiting it has no sign (therefore, treated as Natural numbers)

	// some types are:
	// bool string int and uint, int8/uint8 (and byte as an alias for uint8), 16 32 and 64 and also uintptr
	// float32/64 and complex64/128
	// there's also this rune as an alias for int32

	// initialize the variables below here:
	var smsSendingLimit int = 100
	var costPerMsg float32 = 0.01
	var hasPermission bool = true
	var username string = "Derek"
	fmt.Printf(
		"%v %.02f %v %s\n",
		smsSendingLimit,
		costPerMsg,
		hasPermission,
		username,
	)

	// infering types
	empty := "empty"
	fmt.Println(empty)

}
