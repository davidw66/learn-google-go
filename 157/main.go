package main

import (
	"fmt"
	"math"
)

type square struct {
	length float64
	width  float64
}

func (s square) area() float64 {
	area := s.length * s.width
	return area
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	area := math.Pow(c.radius, 2) * math.Pi
	return area
}

type shape interface {
	area() float64
}

func info(s shape) {
	fmt.Println(s.area())
}

func main() {
	s1 := square{
		length: 105,
		width:  10,
	}
	// fmt.Println(s1.area())

	c1 := circle{
		radius: 1,
	}

	info(s1)
	info(c1)

}
