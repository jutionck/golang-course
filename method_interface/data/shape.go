package data

type Shape interface {
	Area() float64
	Perimeter() float64
}

type ShapeInformation interface {
	GetMeasurement() float64
}
