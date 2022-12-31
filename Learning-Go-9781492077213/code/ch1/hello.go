/*
Run this with:
go run hello.go

Or, if you want to distribute the binary:
go build hello.go
*/
package main

import "fmt"

func main() {
	var foo = 1_234
	fmt.Println("Hello, world!", foo)
}
