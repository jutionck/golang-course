package main

import "fmt"

func main() {

	var customerName string = "James" // camelCase
	var customerAge uint8 = 26
	var customerAddress string = "LA"

	// var 2sss // gak boleh
	// var coba2 string // ini boleh
	// var _name  string // ini boleh

	/*
		%s -> string
		%d -> bilangan bulat
		%f -> bilangan desimal
		%t -> boolean
		%v -> untuk semua
	*/
	fmt.Printf("Nama %s umur %d alamat %s \n", customerName, customerAge, customerAddress)
	fmt.Println("Halo")

	var (
		firstName, lastName string
		age                 int
	)

	firstName = "Jono"
	lastName = "Jini"
	age = 27

	fmt.Println(firstName, lastName, age)

	// Multiple Initialize
	var (
		bootcampName, bootcampAddress = "Enigmacamp", "Jakarta Selatan"
	)
	fmt.Println(bootcampName, bootcampAddress)

	// Inferred Type
	days, month, year := "Monday", "May", 2022
	fmt.Println(days, month, year)

	const phi = 3.14
	fmt.Println(phi)
	const (
		StatusOK      = 200
		StatusCreated = 201
	)

	fmt.Println(StatusOK, StatusCreated)
	// Imutable & mutable

	officeName, _ := "PT Enigma Cipta Humanika", "Bootcamp"
	fmt.Println(officeName)

	// a := "Harus isi"
	var b float32
	fmt.Printf("Nilai B: %f", b) // zero value

}
