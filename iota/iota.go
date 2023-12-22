package main

import "fmt"

type ByteSize int

const (
	_           = iota
	KB ByteSize = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
)

func main() {
	fmt.Printf("KB in decimal %d in binary %b\n", KB, KB)
	fmt.Printf("MB in decimal %d in binary %b\n", MB, MB)
	fmt.Printf("GB in decimal %d in binary %b\n", GB, GB)
	fmt.Printf("TB in decimal %d in binary %b\n", TB, TB)
	fmt.Printf("PB in decimal %d in binary %b\n", PB, PB)
	fmt.Printf("EB in decimal %d in binary %b\n", EB, EB)
}
