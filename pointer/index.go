package main

import "fmt"

func main() {

	// // variabel biasa
	// var num int = 10
	// fmt.Println(num) // -> 0 (zero value)

	// // variable dengan pointer
	// var ptr *int
	// fmt.Println(ptr) // -> <nil> (zero value)

	// // Mendapatkan alamat memori
	// // fmt.Println("num (address)", &num)
	// // fmt.Println("ptr (address)", &ptr)

	// // set address dari num ke ptr
	// ptr = &num
	// fmt.Println("num (address)", &num)
	// fmt.Println("ptr (address)", ptr)

	// // get value
	// fmt.Println("num (value) 1", num)
	// fmt.Println("ptr (value) 1", *ptr)

	// // Simulasi perubahan
	// *ptr = *ptr / 2 // *ptr = 5
	// fmt.Println("num (value) 2", num)
	// fmt.Println("ptr (value) 2", *ptr)

	// Memmbuat sebuah variabel pointer menggunakan keyword new() -> function
	// Ketika membuat variabel dengan keyword new(), sebenernya dia akan mengembalikan sebuah
	// pointer ke data kosong.
	// ex: var name = new(string) === var name *string

	// var name = new(int)
	// *name = 0
	// fmt.Println("name (value)", *name)

	// var name3 *int // <nil>
	// *name3 = 0     // invalid memory address or nil pointer dereference
	// // var name = 0
	// // name3 = &name

	// // fmt.Println(*name)
	// // fmt.Println(&name)
	// // fmt.Println(&name2)
	// // fmt.Println(name3)
	// fmt.Println("name3 (value)", *name3)
	// // fmt.Println(name)

	number := 20
	passByValue(number)
	fmt.Println("passByValue", number)
	fmt.Println()
	fmt.Println("number (address)", &number)
	passByReference(&number)
	fmt.Println("passByReference", number)

}

// Menerapkan pointer di Function
func passByReference(number *int) {
	fmt.Println("passByReference: number (address)", number)
	*number = 10
}

func passByValue(number int) {
	fmt.Println("passByValue: number (address)", &number)
	number = 10
}

// Pass By Value ? -> Duplikasi
// Pass By Reference ? -> Mengirim sebuah alamat ke sebuah variabel, function, method dll
// Contoh penerapan Pass By Reference adalah Pointer.
