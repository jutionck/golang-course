package data

import "math"

type Cube struct {
	Side float64
}

func (c *Cube) Volume() float64 {
	return math.Pow(c.Side, 3)
}

func (c *Cube) Area() float64 {
	return math.Pow(c.Side, 2) * 6
}

func (c *Cube) Perimeter() float64 {
	return 0.0
}

func (c *Cube) GetMeasurement() float64 {
	return c.Side
}
