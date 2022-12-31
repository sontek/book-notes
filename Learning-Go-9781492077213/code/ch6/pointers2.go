/*
Run this with:
go run pointers2.go

Or, if you want to distribute the binary:
go build pointers2.go
*/
package main

import "fmt"

func failedUpdate(px *int) {
	x2 := 20
	px = &x2
}

func update(px *int) {
	*px = 20
}

func main() {
	x := 10
	failedUpdate(&x)
	fmt.Println(x) // prints 10
	update(&x)
	fmt.Println(x) // prints 20
}
