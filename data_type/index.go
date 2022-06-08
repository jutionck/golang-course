package main

import (
	"fmt"
)

func main() {

	// Number: int, uint, float
	// var numberOne int = -1
	// var numberTwo uint = 1
	// var numberThree float32 = 20

	// fmt.Printf("NumberOne: %d \n", numberOne)
	// fmt.Printf("NumberTwo: %d \n", numberTwo)
	// fmt.Printf("NumberThree: %.2f \n", numberThree)

	// // Boolean
	// var isValid bool = true
	// fmt.Printf("isValid: %t \n", isValid)

	// // String
	// var messsage string = "Halo salam kenal"
	// fmt.Printf("message: %s \n", messsage)

	// // Working with String

	// // String Literal
	// getMessage := `Hello....
	// Hai
	// Gracias
	// Hola
	// `
	// fmt.Println(getMessage)

	// getData := "Hari ini adalah hari jum'at ke pasar membeli \"apel, garem, gula\""
	// fmt.Println(getData)

	// // Concat
	// firstName := "Jution"
	// lastName := "Kirana"
	// // age := 27
	// // fullName := firstName + " " + lastName + age // tidak bisa berbeda tipe data
	// fullName := firstName + lastName
	// fmt.Println(fullName)

	// fmt.Println(strings.ToUpper(firstName))
	// fmt.Println(strings.ToLower(firstName))
	// fmt.Println(strings.Contains(firstName, "J")) // Case Sensitive
	// fmt.Println(strings.Count(lastName, "a"))     // Case Sensitive
	// fmt.Println(len(firstName))                   // jumlah karakter

	// // Convert Data Type
	// var index int8 = 15
	// // index = 128
	// // var otherIndex int64 = int64(index)
	// otherIndex := int64(index)
	// // otherIndex = 32768
	// fmt.Println(otherIndex)

	var number int64 = 129
	otherNumber := int8(number)               // 127
	fmt.Println("Number: ", number)           // 129
	fmt.Println("otherNumber: ", otherNumber) // ? 127

	// // Cek Tipe Data
	// fmt.Printf("Tipe data index: %T\n", index)
	// fmt.Printf("Tipe data otherIndex: %T\n", otherIndex)

	// var x int64 = 57
	// y := float64(x)
	// fmt.Printf("%d \n", x)
	// fmt.Printf("%f \n", y)

	// var f float32 = 80.99
	// i := int32(f)
	// fmt.Printf("%f \n", f)
	// fmt.Printf("%d", i)

	// Number to String
	// var a int = 12
	// b := fmt.Sprintf("%d", a)
	// fmt.Printf("print with sprinft: %T \n", b)

	// c := strconv.Itoa(a)
	// fmt.Printf("print with strconv.Itoa %T \n: ", c)

	// String to Number
	// a := "12"
	// b := "13"
	// result := a + b
	// fmt.Println(result)

	// c, _ := strconv.Atoi(a)
	// d, _ := strconv.Atoi(b)
	// result2 := c + d
	// fmt.Println(result2)

	// Array, Slice, Map, Struct, dll
}
