package main

import "fmt"

func main() {

	/*
		switch statement {
		case condition:
			// output
			break (x) // tidak perlu
		default:
			// output
		}
	*/

	point := 100
	// switch point {
	// case 100:
	// 	fmt.Println("Perfect!")
	// case 80:
	// 	fmt.Println("Good!")
	// default:
	// 	fmt.Println("Not Bad!")
	// }

	// Multiple Case
	// switch point {
	// case 80, 90, 100:
	// 	fmt.Println("Perfect!")
	// case 60, 70:
	// 	fmt.Println("Good!")
	// case 50:
	// 	fmt.Println("Awesome!")
	// default:
	// 	fmt.Println("Not Bad!")
	// }

	// Switch Like If Else
	// switch {
	// case point == 100:
	// 	fmt.Println("Perfect!")
	// case (point >= 80) && (point < 100):
	// 	fmt.Println("Good!")
	// case (point > 60) && (point < 80):
	// 	fmt.Println("Awesome!")
	// default:
	// 	fmt.Println("Not Bad!")
	// 	fmt.Println("You need to learn more")
	// }

	// Keyword fallthrough -> memaksa pengecekan 1 tingkat di bawah
	switch {
	case point == 100:
		fmt.Println("Perfect!")
		fallthrough
	case (point >= 80) && (point < 100):
		fmt.Println("Good!")
	case (point > 60) && (point < 80):
		fmt.Println("Awesome!")
	default:
		fmt.Println("Not Bad!")
		fmt.Println("You need to learn more")
	}
}
