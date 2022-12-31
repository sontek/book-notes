package main

import "fmt"

func main() {
	a := []int{2, 4, 6, 8, 10}
	ch := make(chan int, len(a))
	for _, v := range a {
		// v is shared by all go routines and they won't notice
		// it is changed.   `go vet` will catch this.
		// https://go.dev/doc/faq#closures_and_goroutines
		go func() {
			ch <- v * 2
		}()
	}
	for i := 0; i < len(a); i++ {
		fmt.Println(<-ch)
	}
}
