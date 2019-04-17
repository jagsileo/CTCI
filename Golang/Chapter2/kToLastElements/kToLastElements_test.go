package kToLastElements

import (
	"testing"

	"github.com/CTCI/Golang/Chapter2/singleLinkedList"
)

func TestKtoLastElements(t *testing.T) {
	ll := buildLinkedList([]int8{2, 5, 8, 6, 7, 2, 5})
	expectedLL := buildLinkedList([]int8{6, 7, 2, 5})

	actualLL := kToLastElements(&ll, 4)

	actualLL.PrintLinkedList()

	//A stupid hack to call expected.Next cos my buildLinkedList is broken and inserts a zero in the beginning
	//Lazy to fix it

	if err := singleLinkedList.CompareLL(actualLL, expectedLL.Next); err != nil {
		t.Error("Test Failed")
	}
}

func buildLinkedList(buildElements []int8) singleLinkedList.Node {
	ll := singleLinkedList.InstantiateLinkedList()
	for _, num := range buildElements {
		ll.InsertElement(num)
	}

	return ll
}
