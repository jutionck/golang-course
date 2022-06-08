package main

import "fmt"

func main() {
	// utils.PrintStudentMethod()
	// utils.PrintShape()
	// Emebedded Interface
	// var cube data.ShapeGeometry = &data.Cube{
	// 	Side: 4,
	// }
	// fmt.Println("Volume Kubus: ", cube.Volume())
	// fmt.Println("Luas Kubus: ", cube.Area())
	// var cubeInfo data.ShapeInformation
	// cubeInfo.GetMeasurement()

	// interface{}
	// var whatever interface{}
	// fmt.Printf("Type of whatever: %T - %p\n", whatever, &whatever)

	// whatever = "Enigma"
	// fmt.Printf("Type of whatever: %T - %p\n", whatever, &whatever)

	// whatever = 2000
	// fmt.Printf("Type of whatever: %T - %p\n", whatever, &whatever)

	// whatever = true
	// fmt.Printf("Type of whatever: %T - %p\n", whatever, &whatever)

	// whatever = [2]int{1, 2}
	// fmt.Printf("Type of whatever: %T - %p\n", whatever, &whatever)

	// whatever = []string{"mangga", "apel"}
	// fmt.Printf("Type of whatever: %T - %p\n", whatever, &whatever)

	// contoh penggunaan interface kosong
	// fmt.Print
	// fmt.Printf
	// fmt.Println
	// fmt.Sprint
	// dst...

	// Pembuatan interface kosong harus ada -> {}
	// Jika tidak ada maka itu di anggap sebagai pembuatan interface (kontrak)

	printResult(1995)
	printResult("Enigma")
	printResult(false)

	mustNumber(20)
	mustNumber("Enigma")
}

func printResult(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("%d\n", v)
	case string:
		fmt.Printf("%q memiliki panjang %v\n", v, len(v))
	default:
		fmt.Printf("Error uknown %T\n", v)
	}
}

func mustNumber(i interface{}) {
	v, ok := i.(int)
	if ok {
		fmt.Printf("%v\n", v)
	} else {
		fmt.Println("Tipe data harus (int)")
	}
}
