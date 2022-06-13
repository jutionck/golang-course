package main

import (
	"fmt"
	"sync"
)

/*
	Program kita mengalami suatu kejadian yang bernama "race condition"
	race condition: ?
		Kondisi dimana 2 atau lebih goroutine mengakses data yang sama secara bersamaan
	Efek:
		Ex: perhitungan
		Pastinya perhitungan yang terjadi akan kacau

	Go menyediakan solusi, yaitu "mutex"
	Dengan mutex: ketika terjadi rc, hanya goroutine yang beruntung yang dapat mengakses data tsb.
	Kemudian goroutine yang bagaimana ? Status: Wait(paksa), sampai yang mendapatkan data tersebut selesai.

	Go: package yang sama seperti WaitGroup yaitu sync.Mutex
	Ex: Counter
*/

type counter struct {
	val int
}

func (c *counter) Add(x int) {
	c.val = c.val + x
}

func (c *counter) Value() int {
	return c.val
}

func print(till int, message *counter, wg *sync.WaitGroup, mtx *sync.Mutex) {
	for i := 0; i < till; i++ {
		mtx.Lock() // kunci akses: memastikan hanya 1 goroutine yang akses data
		message.Add(1)
		mtx.Unlock() // membuka akses
	}
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	var mtx sync.Mutex
	var meter counter
	till := 100
	for i := 0; i < till; i++ {
		wg.Add(1)
		go print(till, &meter, &wg, &mtx)
	}
	wg.Wait()
	fmt.Println(meter.Value())
}