package main

import (
	"fmt"
)

func main() {
	// penggunaan buffered channel
	message := make(chan string, 3)
	message <- "Joko"
	message <- "Bulan"
	// message <- "Sandy"

	go func() {
		close(message)
	}()

	for n := range message {
		sayHi(n)
	}
}

func sayHi(m string) {
	fmt.Println(m)
}

// func main() {
// 	message := make(chan string) // default un-beffered

// 	// Implementasi Channel Range & Close
// 	// go func() {
// 	// 	// menandakan channel yang digunakan sudah selesai (di close)
// 	// 	// men-non-aktifkan
// 	// 	defer close(message)
// 	// 	message <- "Joko"
// 	// }()

// 	message <- "Joko"
// 	close(message)

// 	for n := range message {
// 		sayHi(n)
// 	}
// }

// func sayHi(m string) {
// 	fmt.Println(m)
// }

// func main() {
// 	var wg sync.WaitGroup
// 	message := make(chan string) // default un-beffered
// 	// sayHi := func() {
// 	// 	// name := <-message
// 	// 	// fmt.Println(name)
// 	// }
// 	wg.Add(1)
// 	go sayHi(message, &wg)
// 	message <- "Joko"
// 	// vMessage := <-message
// 	// fmt.Println(vMessage) // error: deadlock

// 	// v := <-message // penerimaan dari channel ini yang disebut sebuah "blocking"
// 	// fmt.Println(v)
// 	wg.Wait()

// }

// func sayHi(m chan string, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	name := <-m
// 	fmt.Println(name)
// }

/*
	Channel: Sebuah komunikasi antar goroutine
	Channel: dapat sebagai pengirim dan penerima (<-)
	Ilustrasinya:
	Pengirim: channel <- value | mengirim data/value ke sebuah channel
	Penerima: value := <- channel | menerima data dari sebuah channel

	Cara buatnya, sama seperti pembuatan map, slice, menggunakan keyword make(chan dt)
	Ex: message := make(chan string) -> sebuah channel dengan tipe data string
*/
