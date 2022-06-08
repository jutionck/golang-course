package main

/*
	Buatlah sebuah program untuk mencetak jenis mobil "TOYOTA" yang dibagi berdasarkan jenis mobilnya
	dan CC.
	TOYOTA jenismobil: ->
				MVP, SUV, Sedan, dll
				cc: ->
						3000 -> "Inova", "Fortuner"
						2000 -> "Raize", "Avanza"
						2500 -> "Camry", "Rush"
						2000 -> dll
*/

/*
	Tolong hapus aku dari antrian
	[]int{1, 3, 5, 6, 7, 8, 10}

	-> 1
	[]int{3, 5, 6, 7, 8, 10}

	-> 3
	[]int{5, 6, 7, 8, 10}
*/

import "fmt"

func main() {

	// code

	toyota := map[string]map[string][]string{
		"mvp":   {"3000": {"Inova", "Fortuner"}},
		"suv":   {"2000": {"Raize", "Avanza"}},
		"sedan": {"2500": {"Camry", "Rush"}},
	}

	fmt.Println()
	fmt.Println(toyota["mvp"]["3000"]) // output: Inova, Fortuner
	fmt.Println(toyota["suv"]["2000"]) // output: Raize, Avanza
}
