package singleLinkedList

import (
	"log"
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
