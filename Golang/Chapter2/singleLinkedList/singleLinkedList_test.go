package singleLinkedList

import (
	"errors"
	"log"
	"reflect"
	"testing"
)

func TestDeleteLinkedList(t *testing.T) {
	ll := buildLinkedList([]int8{2, 5, 8, 6, 7, 2, 5})
	expectedLL := buildLinkedList([]int8{2, 5, 8, 6, 7, 2})
	deleteLL := ll.DeleteElement(5, 7)

	if err := CompareLL(&expectedLL, deleteLL); err != nil {
		log.Println(deleteLL)
		t.Error(err)
	}

}

func buildLinkedList(buildElements []int8) Node {
	ll := InstantiateLinkedList()
	for _, num := range buildElements {
		ll.InsertElement(num)
	}
	return ll
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
