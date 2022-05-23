package sturct

import "math"

type Rectangle struct {
	height float64
	width  float64
}
type Circle struct {
	Radius float64
}

type Shape interface {
	Area() float64
}

func (r Rectangle) Area() float64 {
	return r.height * r.width
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}
