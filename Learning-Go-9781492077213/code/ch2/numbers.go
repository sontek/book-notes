/*
Run this with:
go run numbers.go

Or, if you want to distribute the binary:
go build numbers.go
*/
package main

import "fmt"

func main() {
	var one = 3.4028234663
	two := 3.40282346634
	var x int = 10
	var y float64 = 30.2
	/* Type conversion */
	var z float64 = float64(x) + y
	var d int = x + int(y)

	fmt.Println("Hello, world!", one == two)
	fmt.Println("z", z)
	fmt.Println("d", d)
}
