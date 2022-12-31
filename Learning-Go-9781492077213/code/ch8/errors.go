package main

import (
	"errors"
	"fmt"
)

type CustomError struct{}

func (ce CustomError) Error() string {
	return "Custom Error"
}

func bam() error {
	return fmt.Errorf("external failure")
}

/* Wrapped external failure with %w */
func boom() error {
	err := bam()
	if err != nil {
		return fmt.Errorf("internal failure: %w", err)
	}
	return nil
}

func kapow() error {
	return CustomError{}
}

func main() {
	err := boom()
	if err != nil {
		fmt.Printf("Caught an error: %s foo\n", err)
		fmt.Println(err)
		wrapepdError := errors.Unwrap(err)
		fmt.Println(wrapepdError)
	}
	err2 := kapow()
	if err2 != nil {
		if errors.Is(err2, CustomError{}) {
			fmt.Println("Got a custom error")
			fmt.Printf("%+v", err)
		}
	}

}
