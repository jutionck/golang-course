package data

import (
	"math"
)

type Circle struct {
	Radius float64
}

func (c *Circle) Area() float64 {
	return math.Pi * math.Pow(c.Radius, 2)
}

func (c *Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}
