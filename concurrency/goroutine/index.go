package main

import (
	"fmt"
	"sync"
	"time"
)

func main() { // main ini adalah goroutine
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		greeting("Jution")
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		sayHello()
	}()

	wg.Wait() // sifatnya seperti blocking, artinya proses eksekusi program tidak akan selesai/diteruskan sebelum goroutine selesai (dipanggil)
}

func greeting(name string) {
	fmt.Println("Start greeting func")
	fmt.Println("Hello", name)
	time.Sleep(1 * time.Second)
	fmt.Println("End greeting func")
}

func sayHello() {
	fmt.Println("Start sayHello func")
	fmt.Println("Hai...")
	time.Sleep(2 * time.Second)
	fmt.Println("End sayHello func")
}

// func greeting(wg *sync.WaitGroup, name string) {
// 	defer wg.Done()
// 	fmt.Println("Start greeting func")
// 	fmt.Println("Hello", name)
// 	time.Sleep(1 * time.Second)
// 	fmt.Println("End greeting func")
// }

// func sayHello(wg *sync.WaitGroup) {
// 	defer wg.Done() // gunakan ini ketika menunggu goroutine selesai
// 	fmt.Println("Start sayHello func")
// 	fmt.Println("Hai...")
// 	time.Sleep(2 * time.Second)
// 	fmt.Println("End sayHello func")
// }

/*
	Goroutine bisa dikatan dapat menjalankan beberapa proses secara bersamaan
	Lantas bagaimana jika kita ingin menunggu beberapa proses tersebut ?
	Kita bisa menggunakan waitGroup untuk menunggu goroutine
	WaitGroup itu sendiri sudah disediakan oleh package Go
	Package Go: sync.WaitGroup
*/
