package main

import "fmt"

func main() {
	// name := "Jution"
	// address := "Ragunan"
	// nama, alamat := printMessage(name, address)
	// fmt.Println(nama, alamat)

	// province, region := location("Jakarta Selatan")
	// fmt.Printf("Saya tinggal di %s propinsi %s \n", region, province)
	// province, region = location("Bandung")
	// fmt.Printf("Saya tinggal di %s propinsi %s \n", region, province)
	// province, region = location("Pesawaran")
	// fmt.Printf("Saya tinggal di %s propinsi %s \n", region, province)

	// oddEven(20) // GENAP
	// oddEven(0)  // Invalid
	// oddEven(11) // GANJIL

	// var animals = []string{"Ayam", "Bebek", "Kucing", "Merpati"}
	// enigmaZoo(animals)
	// fmt.Println()
	// enigmaZooVariadic(animals...) // cara pertama
	// fmt.Println()
	// enigmaZooVariadic("Kucing", "Kadal", "Buaya", "Singa") // cara kedua

	// portofolio("Jution", "Ragunan", "S1", "Tiduran", "Menghilang")

	// Closure
	// printGreetingFunc := greeting
	// fmt.Println(printGreetingFunc("Jution"))

	// Callback Function
	// getValidEmail("jutionck@gmail.com", emailFiter)
	// getValidEmail("jutionck.gmail.com", emailFiter)

	var greeting = func() {
		fmt.Println("Hai salam kenal...!")
	}

	greeting()

}

// func printMessage(name, address string) (string, string) {
// 	getName := fmt.Sprintf("Halo nama saya %s", name)
// 	getAddress := fmt.Sprintf("Saya tinggal di %s", address)
// 	return getName, getAddress
// }

/*
	Membuat sebuah function yang dia akan mengecek dari kota yang diinputkan itu provinsi dan negara mana.

	Jakarta Selatan -> DKI Jakarta, Indonesia
	Cianjur -> Jawa Barat, Indonesia
*/

// func location(city string) (province, region string) { // Named Result
// 	// var province string
// 	// var region string
// 	switch city {
// 	case "Jakarta Selatan", "Jakarta Barat", "Jakarta Timur":
// 		province, region = "DKI Jakarta", "Indonesia"
// 	case "Bandar Lampung", "Lampung Selatan", "Pesawaran":
// 		province, region = "Lampung", "Indonesia"
// 	case "Bogor", "Cianjur", "Bandung":
// 		province, region = "Jawa Barat", "Indonesia"
// 	default:
// 		province, region = "Tidak ada", "Tidak ada"
// 	}
// 	return
// }

// func oddEven(number int) {
// 	if number == 0 {
// 		fmt.Println("Invalid number, can't zero number")
// 		return
// 	}
// 	if number%2 == 0 {
// 		fmt.Printf("Angka %d merupakan bilangan GENAP\n", number)
// 	} else {
// 		fmt.Printf("Angka %d merupakan bilangan GANJIL\n", number)
// 	}
// }

// Variadic Function
/*
	Variadic Function
		-> Pembuatan Function dengan parameter seperti biasa
		-> Dengan penerapan variadic itu sama hal nya dgn kita membuat sebuah slice []string, []int dll
		-> Pembuatannya cukup dengan (...tipeData) ex: ...string, ...int dst
		-> Mempunyai sebuah index
		-> Catatan: Penerapan Variadic Function ini harus di akhir parameter yang dibuat
		-> ex: func numbers(number ...int) (v) | func numbers(number ...int, x, y int) (x)
*/

// func enigmaZoo(animal []string) { // slice
// 	for _, animal := range animal {
// 		fmt.Println(animal)
// 	}
// }

// func enigmaZooVariadic(animal ...string) { // variadic
// 	for _, animal := range animal {
// 		fmt.Println(animal)
// 	}
// }

/*
	Buatkan aku sebuah function yang menerima beberapa argumen:
	-> Nama : Jution
	-> Alamat : Ragunan
	-> Pendidikan: S1
	-> Skills: ["Tidur", "Berlari", "Berenang", "Menghilang"] (variadic)
	func portofolio(...) {
		...
	}
*/

// func portofolio(name, address, education string, skills ...string) {
// 	var mySkill = strings.Join(skills, ", ")
// 	fmt.Printf("Nama \t\t: %s \n", name)
// 	fmt.Printf("Alamat \t\t: %s\n", address)
// 	fmt.Printf("Pendidikan \t: %s\n", name)
// 	fmt.Printf("Skills \t\t: %s \n", mySkill)
// }

// Function Value (closure)
/*
	-> Sebuah function yang bisa dijadikan sebuah value dalam sebuah variabel
	-> Sehingga variabel tersebut bertindak selayaknya sebuah function
*/

// func greeting(name string) string {
// 	return "Salam kenal " + name
// }

// Function as Parameter

/*
	-> Kita akan membuat sebuah callback function yang ditaruh di parameter
	-> Artinya function itu bertindak sebagai parameter
	-> ex: Saya mau ngefilter email, '@' -> jutionck@gmail.com, jution.kirana@enigmacamp.com, dst
*/

// func getValidEmail(email string, callback func(string) string) {
// 	emailFitered := callback(email)
// 	fmt.Println("Your email", emailFitered)
// }

// func emailFiter(email string) string {
// 	if strings.Contains(email, "@") {
// 		return email
// 	} else {
// 		return "Not Valid!"
// 	}
// }

// Anonymous Function
/*
	-> Function yang tidak didefinisikan nama function-nya
	-> Disimpan dalam sebuah variabel
	-> Pembuatannya bisa di lakukan di dalam func main()
	-> ex: var greet = func(param string) string {
				// body function
		}
	-> greet()
*/
