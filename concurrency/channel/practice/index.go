package main

import (
	"fmt"
	"sync"
)

/*
	Contoh penggunaan channel for - range - close
	Ex: square()
*/

func main() {
	var wg sync.WaitGroup
	squareChan := make(chan int)
	cubeChan := make(chan int)
	doneChan := make(chan bool)
	numbers := []int{2, 3, 4, 5, 6, 7, 8}
	wg.Add(2)
	go square(numbers, squareChan, &wg)
	go cube(numbers, cubeChan, &wg)

	// go func() {
	// 	for i := range squareChan {
	// 		fmt.Println("Square result: ", i)
	// 	}
	// }()

	// go func() {
	// 	for i := range cubeChan {
	// 		fmt.Println("Cube result: ", i)
	// 	}
	// }()
	// wg.Wait() //  blocking (error: deadlock!)
	// doneChan <- true
	go func() {
		wg.Wait() //  blocking (error: deadlock!)
		doneChan <- true
	}()

	for {
		select {
		case msg1 := <-squareChan:
			fmt.Println("Square result: ", msg1)
		case msg2 := <-cubeChan:
			fmt.Println("Cube result: ", msg2)
		case <-doneChan:
			fmt.Println("Done")
			return
		}
	}
}

func square(numbers []int, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for _, v := range numbers {
		ch <- v * v
		fmt.Println("Berhasil mengirim data (square)", v, "ke channel")
	}
}

func cube(numbers []int, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for _, v := range numbers {
		ch <- v * v * v
		fmt.Println("Berhasil mengirim data (cube)", v, "ke channel")
	}
}
