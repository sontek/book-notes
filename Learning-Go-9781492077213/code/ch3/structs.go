/*
Run this with:
go run structs.go

Or, if you want to distribute the binary:
go build structs.go
*/
package main

import "fmt"

type person struct {
	name string
	age  int
	pet  string
}

func main() {
	var fred person
	fmt.Println("Fred Name", fred.name)
	bob := person{}
	fmt.Println("Bob Age", bob.age)
	julia := person{
		"Julia",
		40,
		"cat",
	}
	fmt.Println("Julia", julia)
	beth := person{
		age:  30,
		name: "Beth",
	}
	fmt.Println(beth)
}
