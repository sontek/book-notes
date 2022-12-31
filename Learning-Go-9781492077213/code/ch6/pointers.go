/*
Run this with:
go run pointers.go

Or, if you want to distribute the binary:
go build pointers.go
*/
package main

import "fmt"

func main() {
	x := 10
	pointerToX := &x
	valueOfX := *&x
	fmt.Println(valueOfX)    // prints 10
	fmt.Println(pointerToX)  // prints a memory address
	fmt.Println(*pointerToX) // prints 10
	z := 5 + *pointerToX
	fmt.Println(z) // prints 15

	var y *int
	fmt.Println(y == nil) // prints true
	fmt.Println(*y)       // panics”
}
