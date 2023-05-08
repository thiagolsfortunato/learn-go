package forms

import (
	"math"
)

type Form interface {
	area() float64
}

type Rectangle struct {
	height float64
	width float64
}

func (r Rectangle) Area() float64 {
	return r.height * r.width
}

type Circle struct {
	ratio float64
}

func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.ratio, 2)
}

// func GetArea(f Form) {
// 	fmt.Printf("Area: %0.2f", f.area())
// }
