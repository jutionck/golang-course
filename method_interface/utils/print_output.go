package utils

import (
	"fmt"
	"golang-fundamental/method_interface/data"
)

func PrintStudentMethod() {
	// fmt.Println(d.SayHi("Jono"))

	// fmt.Println(data.sayHello()) // tidak akan bisa karena dia unexported di package data

	// fmt.Println(d.SayHello())

	// cara akses sebuah method
	// greeting := d.Greeting{Name: "Jack"}
	// greeting.SayHi().SayGoodBye("Bulan") // method chaining

	// fmt.Printf("greeting (address) %p \n: ", &greeting)
	// greeting2 := d.Greeting2{
	// 	Name:    "Bulan",
	// 	Address: "Bali",
	// }
	// greeting2.SayHi()

	// Method Pointer
	// Pass By Value & Pass By Reference
	student := data.Student{Name: "Balabala", Age: 20}
	fmt.Println("student (value): ", student.Name)

	// Panggil method PBV
	student.ChangeNamePBV("Horehore", 25)
	fmt.Println("student.Name (value) after ChangeNamePBV: ", student.Name)
	fmt.Println("student.Age (value) after ChangeNamePBV: ", student.Age)

	// Panggil method PBR
	student.ChangeNamePBR("Horehore", 25)
	fmt.Println("student.Name (value) after ChangeNamePBR: ", student.Name)
	fmt.Println("student.Age (value) after ChangeNamePBR: ", student.Age)
}
