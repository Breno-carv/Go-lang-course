package main

import (
	"fmt"
	"time"

	tinytime "github.com/wagslane/go-tinytime" // go get {remote}/{user}/{pkg}
)

func main() {
	tt := tinytime.New(1585750374)

	tt = tt.Add(time.Hour * 48)
	fmt.Println(tt)
}
