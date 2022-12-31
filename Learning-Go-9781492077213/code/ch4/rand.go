/*
Run this with:
go run rand.go

Or, if you want to distribute the binary:
go build rand.go
*/
package main

import (
	"fmt"
	"math/rand"
)

type person struct {
	name string
	age  int
	pet  string
}

func main() {
	/* TODO: Why is this always 1? */
	n := rand.Intn(10)
	if n == 0 {
		fmt.Println("That's too low")
	} else if n > 5 {
		fmt.Println("That's too big:", n)
	} else {
		fmt.Println("That's a good number:", n)
	}
}
