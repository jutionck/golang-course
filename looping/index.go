package main

import "fmt"

func main() {

	/*
		Saya programmer Golang
		Saya programmer Golang
		Saya programmer Golang
		Saya programmer Golang
		Saya programmer Golang
	*/

	// fmt.Println("Saya programmer Golang")
	// fmt.Println("Saya programmer Golang")
	// fmt.Println("Saya programmer Golang")
	// fmt.Println("Saya programmer Golang")
	// fmt.Println("Saya programmer Golang")
	// fmt.Println("Saya programmer Golang")
	// fmt.Println("Saya programmer Golang")
	// fmt.Println("Saya programmer Golang")
	// fmt.Println("Saya programmer Golang")
	// fmt.Println("Saya programmer Golang")

	// Basic For
	// for i := 0; ; i += 2 {
	// 	fmt.Println(i, "- Saya programmer Golang")
	// }

	// For a While
	// i := 0
	// for i < 10 { // hanya kondisi
	// 	fmt.Println(i, "- Saya programmer Golang")
	// 	i += 2 // ini harus ada agar tidak terjadi infinite loop (perulangan tak henti)
	// }

	// For Ever -> for tanpa argumen
	// i := 0
	// j := 0
	// for {
	// 	fmt.Println(i, "- Saya programmer Golang")
	// 	i++

	// 	// kondisi untuk menghentikan perulangan
	// 	if i == 9 {
	// 		break
	// 	}
	// }

	// for {
	// 	fmt.Println(i, "- Saya programmer Golang")
	// 	i++
	// 	time.Sleep(1 * time.Second)
	// 	for {
	// 		fmt.Println(i, "- Saya programmer React")
	// 		j++
	// 		time.Sleep(1 * time.Second)
	// 		if j == 4 {
	// 			break
	// 		}
	// 	}
	// 	if i == 4 {
	// 		break
	// 	}
	// }

	// Break ? -> Memaksa untuk berhenti dalam kondisi tertentu
	// Continue ? -> Memaksa untuk lanjut

	// Simulasi even and odd
	// 1, 3, 5, 7, 9 dst.. (continue)
	// 2, 4, 6, 8 dst.. (break)

	for i := 1; i < 12; i++ {
		if i%2 == 1 {
			break
		}

		// if i > 8 {
		// 	break
		// }

		fmt.Println("Number: ", i) // output: 2, 4, 6, 8
	}

	fmt.Println("Selesai")

	
}
