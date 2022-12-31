/*
Run this with:
go run slices.go

Or, if you want to distribute the binary:
go build slices.go
*/
package main

import "fmt"

func main() {
	var x []int
	x = append(x, 2)
	x = append(x, 3)
	x = append(x, 4, 5, 6)
	y := make([]int, 5)
	fmt.Println("Hello, world!", x)
	fmt.Println("Length", len(x), cap(x))
	fmt.Println("Static slice!", y)
	fmt.Println("Length", len(y), cap(y))
	fmt.Println("======")
	xx := []int{1, 2, 3, 4}
	yy := xx[:2]
	z := xx[1:]
	d := xx[1:3]
	e := xx[:]
	fmt.Println("x:", xx)
	fmt.Println("y:", yy)
	fmt.Println("z:", z)
	fmt.Println("d:", d)
	fmt.Println("e:", e)
	fmt.Println("====== SHARED MEMORY =====")
	new := make([]int, len(xx))
	copy(new, xx)
	z[0] = 9999
	fmt.Println("x:", xx)
	fmt.Println("z:", z)
	fmt.Println("new:", new)
}
