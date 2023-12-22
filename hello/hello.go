package main

import (
	"fmt"
)

func main() {
	const age = 22
	x := age * 7
	fmt.Println("Hello 😇 🐠", x)
	fmt.Println(`This is
	a
	string literal
	`)
	adams := 42

	fmt.Printf("adams in binary: %b\n", adams)

	fmt.Printf("adams in hex: %X\n", adams)

	a, b, c, d, e, f := 0, 1, 2, 3, 4, 5

	fmt.Printf("hex: 0 is %x, 1 is %X, 2 is %X, 3 is %X, 4 is %X, 5 is %X\n", a, b, c, d, e, f)
	fmt.Printf("binary: 0 is %b, 1 is %b, 2 is %b, 3 is %b, 4 is %b, 5 is %b", a, b, c, d, e, f)
}

// 💪 shift+ctrl+e then space bar for emojis
