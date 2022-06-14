package main

import (
	"errors"
	"fmt"
)

/*
	Jenis Linked List:
	1. Single
	2. Double
	3. Circular
*/

type node struct {
	data string
	next *node
}

type mySingleLinkedList struct {
	size int
	head *node
}

func (sl *mySingleLinkedList) validateDuplicate(node node) error {
	current := sl.head
	for current.next != nil {
		if current.next.data == node.data {
			return errors.New("no duplicate number")
		}
	}
	return nil
}

func (sl *mySingleLinkedList) addToHead(name string) {
	newNode := &node{
		data: name,
	}
	if sl.head == nil {
		sl.head = newNode
	} else {
		err := sl.validateDuplicate(*newNode)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		newNode.next = sl.head
		sl.head = newNode
	}
	sl.size++
}

func (sl *mySingleLinkedList) addToTail(name string) {
	newNode := &node{
		data: name,
	}
	if sl.head == nil {
		sl.head = newNode
	} else {
		current := sl.head
		for current.next != nil {
			current = current.next
		}
		current.next = newNode
	}
	sl.size++
}

func (sl *mySingleLinkedList) iterateList() {
	for node := sl.head; node != nil; node = node.next {
		fmt.Println(node.data)
	}
}

func NewLinkedList() *mySingleLinkedList {
	return &mySingleLinkedList{}
}

func main() {
	singleList := NewLinkedList()
	singleList.addToHead("8")
	singleList.addToHead("8")
	singleList.addToHead("10")
	singleList.addToHead("8")
	// singleList.addToHead("29")
	singleList.addToTail("15")
	singleList.addToTail("1")
	singleList.iterateList()
	fmt.Println()
	fmt.Println("Size: ", singleList.size)
}
