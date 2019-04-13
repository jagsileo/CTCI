package main

import (
	"fmt"
)

type node struct {
	value int8
	next  *node
}

func main() {
	exampleLL := instantiateLinkedList()
	llPtr := &exampleLL
	llPtr = llPtr.insertElement(1)
	llPtr = llPtr.insertElement(5)
	llPtr = llPtr.insertElement(3)
	llPtr = llPtr.insertElement(2)
	fmt.Println("Inserted 1,5,3,2")
	llPtr.printLinkedList()

	llPtr = llPtr.deleteElement(3)
	fmt.Println("Deleted 3")
	llPtr.printLinkedList()

	llPtr = llPtr.deleteElement(0)
	fmt.Println("Deleted 0")
	llPtr.printLinkedList()

	llPtr = llPtr.deleteElement(2)
	fmt.Println("Deleted 2")
	llPtr.printLinkedList()

}

func instantiateLinkedList() node {
	exampleLL := node{}
	exampleLL.value = 0
	exampleLL.next = nil
	return exampleLL
}

func (ll *node) insertElement(val int8) *node {

	n := ll

	if n == nil {
		return nil
	}

	for n.next != nil {
		n = n.next
	}

	n.next = &node{
		value: val,
		next:  nil,
	}

	return ll

}

func (ll *node) deleteElement(val int8) *node {
	n := ll

	if n == nil {
		return nil
	}

	if n.value == val {
		return n.next
	}

	for n.next != nil && n.next.value != val {
		n = n.next
	}

	n.next = n.next.next

	return ll

}

func (ll *node) printLinkedList() {
	n := ll
	for n != nil {
		fmt.Println(n.value)
		n = n.next
	}
}
