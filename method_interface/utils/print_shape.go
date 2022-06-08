package utils

import (
	"fmt"
	"golang-fundamental/method_interface/data"
)

func PrintShape() {

	// rectangle := new(data.Rectangle)
	// rectangle.Length = 5
	// rectangle.Width = 5

	// getRectangleArea := rectangle.Area()
	// getRectanglePerimeter := rectangle.Perimeter()

	// fmt.Printf("Luas Persegi Panjang (P = %2.f, L= %2.f) adalah: %2.f \n", rectangle.Length, rectangle.Width, getRectangleArea)
	// fmt.Printf("Keliling Persegi Panjang (P = %2.f, L= %2.f) adalah: %2.f \n", rectangle.Length, rectangle.Width, getRectanglePerimeter)

	// With Interface
	var s data.Shape
	s = &data.Rectangle{
		Length: 5,
		Width:  5,
	}
	fmt.Println("Luas Persegi Panjang: ", s.Area())
	fmt.Println("Keliling Persegi Panjang: ", s.Perimeter())

	s = &data.Circle{
		Radius: 5,
	}
	fmt.Println("Luas Lingkaran: ", s.Area())
	fmt.Println("Keliling Lingkaran: ", s.Perimeter())

	// fmt.Println(s.(*data.Circle).Radius) // casting

}
