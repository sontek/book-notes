/*
Run this with:
go run slices.go

Or, if you want to distribute the binary:
go build slices.go
*/
package main

import "fmt"

func add(slice []int, num int) {
	slice[0] = 999
	slice = append(slice, num)
}

func main() {
	slice := []int{1, 2, 3}
	for i := 0; i < 100; i++ {
		add(slice, i)
		fmt.Println(slice)
	}
	slice = append(slice, 5555)
	fmt.Println(slice)
}
