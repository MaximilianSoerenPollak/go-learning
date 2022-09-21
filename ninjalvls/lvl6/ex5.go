package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}

type square struct {
	sidelength float64
}

func (c circle) area() float64 {
	return (c.radius * c.radius) * math.Pi
}

func (s square) area() float64 {
	return (s.sidelength * s.sidelength)
}

type shape interface {
	area() float64
}

func info(s shape) {
	fmt.Println(s.area())
}

func main() {
	c := circle{
		radius: 4,
	}
	fmt.Println(c.area())
	s := square{
		sidelength: 3,
	}
	fmt.Println(s.area())

}
