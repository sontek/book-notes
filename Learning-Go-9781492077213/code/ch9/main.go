// Main package which tries to get the double of two
package main

import (
	"fmt"

	print "ch9/formatter"
	"ch9/math"
)

func main() {
	num := math.Double(2)
	output := print.Format(num)
	fmt.Println(output)
}
