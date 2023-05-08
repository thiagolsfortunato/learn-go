package main

import (
	"fmt"
	"math"
)

type form interface {
	area() float64
}

type rectangle struct {
	height float64
	width float64
}

func (r rectangle) area() float64 {
	return r.height * r.width
}

type circle struct {
	ratio float64
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.ratio, 2)
}

func getArea(f form) {
	fmt.Printf("Area: %0.2f", f.area())
}

func main() {
	r := rectangle{5, 5}
	getArea(r)

	c := circle{10}
	getArea(c)
}