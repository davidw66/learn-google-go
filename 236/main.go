package main

import (
	"fmt"
)

type customErr struct {
	errInfo string
}

func (ce customErr) Error() string {
	return fmt.Sprintf("This is a custom error: %v", ce.errInfo)
}

func main() {
	ce := customErr{
		errInfo: "this is the custom error description",
	}
	foo(ce)
}

func foo(e error) {
	fmt.Println("foo ran - ", e)
	// Using assertion to print ce.errInfo
	fmt.Println("Using assertion - ", e.(customErr).errInfo)
}
