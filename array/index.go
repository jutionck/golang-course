package main

import "fmt"

func main() {

	/*
		Array -> Kumpulan Data dengan tipe data yang sama
		Array -> Harus didefinisikan alokasinya
	*/

	// var fruits [4]string
	// fruits[0] = "Mangga"
	// fruits[1] = "Apel"
	// fruits[2] = "Jeruk"
	// fruits[3] = "Alpukat"

	// // Cara memanggil / cetak
	// fmt.Println(fruits)
	// fmt.Println()
	// fmt.Printf("%s %s %s %s", fruits[0], fruits[1], fruits[2], fruits[3])

	// Inisialisasi array di awal
	// var fruits = [4]string{"Mangga Nanas", "Apel", "Jeruk", "Alpukat"}
	// fmt.Println(fruits)
	// fmt.Printf("%s %s %s %s", fruits[0], fruits[1], fruits[2], fruits[3])
	// fmt.Println("Jumlah element array fruits adalah: ", len(fruits))

	// fmt.Printf("%q", fruits) // %q -> cetak tiap elemen di tandai dengan double quote [""]

	// bootcamp := [2]string{"Enigma", "Enigma Camp"}
	// fmt.Println(bootcamp)
	// fmt.Printf("%q", bootcamp)

	/*
		Inisialiasi Value:
		- Horizontal -> ex: bootcamp := [2]string{"Enigma", "Enigma Camp"}
		- Vertikal -> ex: bootcamp := [2]string{
																		"Enigma",
																		"Enigma Camp",
																	}
	*/

	// var cars [2]string
	// // Horizontal
	// cars = [2]string{"BMW", "VW"}
	// fmt.Println(cars)

	// // Vertikal
	// cars = [2]string{
	// 	"TOYOTA",
	// 	"HONDA",
	// }
	// fmt.Println(cars)

	// // Iniasialisasi tanpa diketahui (tak terbatas) jumlah elemennya -> [...]
	// var days = [...]string{"Senin", "Selasa", "Rabu", "Kamis"}
	// fmt.Println(len(days))
	// fmt.Println(days)

	// var month = [...]string{"Januari", "Februari"}
	// fmt.Println(len(month))
	// fmt.Println(month)

	// Multimedimensi Array -> [2][3]int
	// var numbers = [2][3]int{
	// 	{1, 2, 3},
	// 	{4, 5, 6},
	// 	// {4, 5, 6},
	// }

	// var numbers2 [1][2][3][4]int
	// fmt.Print(numbers2)

	// fmt.Println(numbers)
	// fmt.Println(len(numbers)) // 2

	// Looping Array
	var students = [4]string{
		"Jack",
		"Milla",
		"Thomas",
		"Anna",
	}

	// for loop
	fmt.Println("For Loop Basic")
	for i := 0; i < len(students); i++ {
		fmt.Printf("Mahasiswa ke - %d: %s \n", i, students[i])
	}

	fmt.Println("For Loop Range")
	for i, student := range students {
		fmt.Printf("Mahasiswa ke - %d: %s \n", i, student)
	}

	// Cetak value
	for _, student := range students {
		fmt.Printf("Mahasiswa: %s \n", student)
	}

	// Cetak index
	for i := range students {
		fmt.Printf("Mahasiswa index ke - %d\n", i)
	}

	fmt.Printf("%T \n", students)

}
