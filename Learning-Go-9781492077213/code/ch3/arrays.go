/*
Run this with:
go run arrays.go

Or, if you want to distribute the binary:
go build arrays.go
*/
package main

import "fmt"

func main() {
	var x [3]int
	x[0] = 2
	x[1] = 3
	x[2] = 4
	fmt.Println("Hello, world!", x)
}
