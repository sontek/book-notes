/*
Run this with:
go run for.go

Or, if you want to distribute the binary:
go build for.go
*/
package main

import (
	"fmt"
)

type person struct {
	name string
	age  int
	pet  string
}

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	fmt.Println("=====")
	foo := []int{2, 4, 6, 8}
	for index, value := range foo {
		fmt.Println(index, value)
	}
	fmt.Println("=====")
	bar := map[string]string{
		"foo": "bar",
		"baz": "qux",
	}
	for key, value := range bar {
		fmt.Println(key, value)
	}

}
