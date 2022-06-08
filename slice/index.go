package main

import "fmt"

func main() {

	/*
		Slice -> sama seperti array, bedanya tidak perlu definisikan jumlah elemen
		Slcie -> merupakan reference array
	*/

	// var personArray = [2]string{"Antonio", "Cristiano"} // Array
	// // var person2 = [...]string{"Antonio", "Cristiano"}   // Array
	// var personSlice = []string{"Antonio", "Cristiano"} // Slice

	// // fmt.Printf("person: %T \n", person)
	// // fmt.Printf("person2: %T \n", person2)
	// // fmt.Printf("person3: %T \n", person3)

	// fmt.Println(personArray[0], personArray[1])
	// fmt.Println(personSlice[0:3]) // membuat sebuah array

	/*
		len -> untuk menghitung jumlah elemen slice
		cap -> untuk menghitung lebar atau kapasitas slice
	*/

	// var fruits = []string{"Apel", "Mangga", "Jeruk", "Alpukat"}
	// fmt.Println("len: ", len(fruits)) //4
	// fmt.Println("cap: ", cap(fruits)) //4

	// getFruit := fruits[0:2]
	// fmt.Println(getFruit)
	// fmt.Println("len: ", len(getFruit))
	// fmt.Println("cap: ", cap(getFruit))

	// append() -> menambahkan elemen pada slice.
	// Elemen yang di tambahkan, posisinya di akhir
	// Jumlah Kapasitas memenuhi, pasti masuk ke array yang mempunya referensi sama (lama)
	// Jumlah Kapasitas melebihi, otomatis yang masuk itu akan di jadikan sebuah referensi baru
	// Ex: Cap: 4, 5 (ini akan di anggap sebuah array baru)

	// fmt.Printf("%p \n", &fruits[0])
	// getFruit = append(getFruit, "Kelapa", "Kiwi")
	// fmt.Println(getFruit)
	// fmt.Println("len: ", len(getFruit))
	// fmt.Println("cap: ", cap(getFruit))
	// fmt.Printf("%p \n", &getFruit[0])

	// fmt.Println(strings.Repeat("-", 20))
	// getFruit = append(getFruit, "Strawberry")
	// fmt.Println(getFruit)
	// fmt.Println("len: ", len(getFruit))
	// fmt.Println("cap: ", cap(getFruit))
	// fmt.Printf("%p \n", &getFruit[0])

	// make -> Membuat slice kosong dengan spesifik length
	cities := make([]string, 3)
	fmt.Printf("cities %T \n", cities)

	cities[0] = "Jakarta Selatan"
	cities[1] = "Jakarta Pusat"
	cities[2] = "Jakarta Barat"
	// cities[10] = "Bandung" // index out of range
	fmt.Printf("%q ", cities)
}
