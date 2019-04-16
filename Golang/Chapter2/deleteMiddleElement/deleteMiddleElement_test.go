package deleteMiddleElement

import (
	"testing"

	"github.com/CTCI/Golang/Chapter2/singleLinkedList"
)

func TestDeleteMiddleElement(t *testing.T) {
	ll := buildLinkedList([]int8{2, 5, 8, 6, 7, 2, 5})

	n := &ll
	for i := 0; i < 5; i++ {
		n = n.Next
	}

	deleted := deleteMiddleElement(n)
	deleted.PrintLinkedList()
	n = &ll
	for i := 0; i < 4; i++ {
		n = n.Next
	}

	n.Next = deleted

	expectedLL := buildLinkedList([]int8{2, 5, 8, 6, 2, 5})

	if err := singleLinkedList.CompareLL(&ll, &expectedLL); err != nil {
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
