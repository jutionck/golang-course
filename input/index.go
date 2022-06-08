package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	// var (
	// 	fullName, address string
	// 	age               int
	// )

	// fmt.Println("Data Diri Trainee PT Enigma Cipta Humanika")
	// fmt.Println(strings.Repeat("-", 48))

	// fmt.Print("Masukkan nama anda: ")
	// fmt.Scanln(&fullName)

	// fmt.Print("Masukkan alamat anda: ")
	// fmt.Scanln(&address)

	// fmt.Print("Masukkan umur anda: ")
	// fmt.Scanln(&age)

	// fmt.Println(strings.Repeat("-", 48))
	// fmt.Print("Nama \t Alamat \t\t Umur\n")
	// fmt.Println(strings.Repeat("-", 48))
	// fmt.Printf("%s \t %s \t\t %d\n", fullName, address, age)
	// fmt.Println(strings.Repeat("-", 48))

	// bufio.NewReader(os.Stdin)
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Print("Masukkan Nama Anda: ")

	name, _ := inputReader.ReadString('\n')
	fmt.Printf("Nama saya: %s", name)

}
