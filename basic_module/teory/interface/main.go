package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
}

type square struct {
	side float64
}

func (s square) area() float64 {
	return math.Pow(s.side, s.side)
}

type circle struct {
	ray float64
}

func (c circle) area() float64 {
	return math.Pi * (c.ray * c.ray)
}

func viewArea(s shape) {
	fmt.Printf("the area of ​​the shape is %0.2f\n", s.area())
}

func main() {

	c := circle{10}
	viewArea(c)

	s := square{2}
	viewArea(s)

}
