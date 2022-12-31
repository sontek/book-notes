package main

import "fmt"

type MyInt int
type MyIntArray []int

func main() {
	var i interface{}
	var mine MyInt = 20
	i = mine
	i2 := i.(MyInt)
	fmt.Println(i2 + 1)
	i3 := i.(MyIntArray)
	fmt.Println(i3)
}
