package main

import "fmt"

func main() {

	/*
		Defer -> sebuah baris eksekusi yang terakhir
	*/

	// fmt.Println(1)
	// fmt.Println(3)
	// defer fmt.Println(2) // penerapan defer
	defer fmt.Println("Ini ke cetak")
	// defer b()
	// defer c()
	a()
	b()

	// number := 6

	// if number == 6 {
	// 	fmt.Println(4)
	// 	defer fmt.Println(6)
	// }

	// fmt.Println(5)

	// for i := 0; i < 10; i++ {
	// 	defer fmt.Println(i)
	// }

}

func a() {
	fmt.Println("a Function")
	// defer fmt.Println("a function with defer")
}

func b() {
	fmt.Println("b Function")
	panic("")
}

// func c() {
// 	fmt.Println("c Function")
// }

// func d() {
// 	fmt.Println("d Function")
// }
