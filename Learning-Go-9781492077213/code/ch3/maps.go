/*
Run this with:
go run maps.go

Or, if you want to distribute the binary:
go build maps.go
*/
package main

import "fmt"

func main() {
	totalWins := map[string]int{}
	totalWins["Orcas"] = 1
	totalWins["Lions"] = 2
	fmt.Println("Orcas", totalWins["Orcas"])
	fmt.Println("Kittens", totalWins["Kittens"])
	totalWins["Kittens"]++
	fmt.Println("Kittens", totalWins["Kittens"])
	totalWins["Lions"] = 3
	fmt.Println("Lions", totalWins["Lions"])
}
