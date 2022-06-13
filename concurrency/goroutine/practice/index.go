package main

import (
	"fmt"
	"sync"
	"time"
)

/*
	Mencari Total Penjualan masing-masing toko yang sudah disediakan
	Manfaatkan concurrency: WaitGroup
*/

type PenjualanToko struct {
	KodeToko  string
	Penjualan []float64
}

func main() {

	var waitgroup sync.WaitGroup
	// sample data
	dataToko := []PenjualanToko{
		{"123", []float64{1, 2, 3, 4, 5}},
		{"124", []float64{13, 21, 33}},
		{"125", []float64{4, 2}},
		{"126", []float64{21, 22, 23, 9, 1, 22}},
		{"127", []float64{8, 4, 10}},
	}

	for i := 0; i < len(dataToko); i++ {
		waitgroup.Add(1)
		println(i)

		tugas := func(j int) {
			defer waitgroup.Done()
			HitungPenjualan(dataToko[j])
			println(j)
		}

		go tugas(i)
	}
	waitgroup.Wait()
}

func HitungPenjualan(n PenjualanToko) {
	fmt.Printf("Start for %s : %v\n", n.KodeToko, time.Now())
	var totalPenjualan float64
	for _, penjualan := range n.Penjualan {
		totalPenjualan = totalPenjualan + penjualan
	}
	time.Sleep(time.Second * time.Duration(len(n.Penjualan)))
	fmt.Printf("Total Penjualan %s = %f\n", n.KodeToko, totalPenjualan)
	fmt.Printf("Finish for %s : %v\n", n.KodeToko, time.Now())
}
