package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Go OS:", runtime.GOOS)
	fmt.Println("Architecture:", runtime.GOARCH)
}
