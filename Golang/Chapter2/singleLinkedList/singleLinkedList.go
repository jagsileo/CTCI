package singleLinkedList

import (
	"errors"
	"fmt"
	"reflect"
)

type Node struct {
	Value int8
	Next  *Node
}

func InstantiateLinkedList() Node {
	exampleLL := Node{}
	return exampleLL
}

func (ll *Node) InsertElement(val int8) *Node {

	n := ll

	if n == nil {
		return nil
	}

	for n.Next != nil {
		n = n.Next
	}

	n.Next = &Node{
		Value: val,
		Next:  nil,
	}

	return ll

}

func (ll *Node) DeleteElement(val int8, idx int) *Node {

	n := ll
	counter := 0
	if n == nil {
		return nil
	}

	if n.Value == val && counter == idx {
		return n.Next
	}

	for n.Next != nil && counter+1 != idx {

		counter++
		n = n.Next
	}

	if counter+1 == idx && n.Next.Value == val {
		n.Next = n.Next.Next
	}

	return ll

}

func (ll *Node) PrintLinkedList() {
	n := ll
	for n != nil {
		fmt.Println(n.Value)
		n = n.Next
	}
}

func CompareLL(llA, llB *Node) error {
	arrA := []int8{}
	arrB := []int8{}

	n := llA
	for n != nil {
		arrA = append(arrA, n.Value)
		n = n.Next
	}

	n = llB
	for n != nil {
		arrB = append(arrB, n.Value)
		n = n.Next
	}

	if !reflect.DeepEqual(arrA, arrB) {
		return errors.New("Test Failed")
	}
	return nil
}
