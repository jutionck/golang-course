package main

import "fmt"

func main() {

	// 2 ? genap atau ganjil
	// a := 210
	// if a%2 == 0 {
	// 	fmt.Println("This is even number")
	// } else {
	// 	fmt.Println("This is odd number")
	// }

	// Simulasi untuk menghitung Nilai (GPA)
	/*
		GPA: >=3.5 s.d <=4.0 -> Cumlaude
		GPA: >=3.0 s.d <3.5 -> Memuaskan
		GPA: >= 2.75 s.d < 3.0 -> Cukup
		GPA: < 2.75 -> Kurang Memuaskan
	*/

	// GPA := 2.0

	// if GPA > 4.0 || GPA < 0 {
	// 	fmt.Printf("GPA %.2f: Not Found", GPA)
	// } else if GPA >= 3.5 && GPA <= 4.0 {
	// 	fmt.Printf("GPA %.2f: Cumlaude", GPA)
	// } else if GPA >= 3.0 && GPA < 3.5 {
	// 	fmt.Printf("GPA %.2f: Memuaskan", GPA)
	// } else if GPA >= 2.75 && GPA < 3.0 {
	// 	fmt.Printf("GPA %.2f: Cukup", GPA)
	// } else {
	// 	fmt.Printf("GPA %.2f: Kurang Memuaskan", GPA)
	// }

	// fmt.Println()

	// Simulasi Perhitungan Nilai
	/*
		Percent >=100% -> Perfect
		Percent >70% -> Good
		Percent <70% -> Not Bad

		Point / 100 = ?%
		10000 / 100 = 100 -> Perfect
		8000 / 100 = 80 -> Good
	*/

	point := 9000.0

	if percent := point / 100; percent >= 100 {
		fmt.Printf("%.1f%% perfect ", percent)
	} else if percent >= 70 {
		fmt.Printf("%.1f%s good ", percent, "%")
	} else {
		fmt.Printf("%.1f%s not bad ", percent, "%")
	}
	fmt.Println()
}
