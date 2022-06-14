package main

import "fmt"

type queue []string

func (q *queue) enqueue(name string) {
	*q = append(*q, name)
}

func (q *queue) dequeue() string {
	data := (*q)[0]
	*q = (*q)[1:]
	return data
}

func main() {
	q := make(queue, 0)
	q.enqueue("51") // address 1
	q.enqueue("62") // address 1
	q.enqueue("2")  // address 1
	q.enqueue("48")
	fmt.Println("List:")
	fmt.Println(q) // address 1

	q.dequeue()
	q.dequeue()
	fmt.Println("List:")
	fmt.Println(q)
}
