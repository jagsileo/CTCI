package removeDuplicate

import (
	"testing"

	"github.com/CTCI/Golang/Chapter2/singleLinkedList"
)

func TestRemoveDuplicate(t *testing.T) {
	ll := buildLinkedList([]int8{2, 5, 8, 6, 7, 2, 5})
	expectedLL := buildLinkedList([]int8{2, 5, 8, 6, 7})
	dedupedLL := removeDuplicate(&ll)
	dedupedLL.PrintLinkedList()

	if err := singleLinkedList.CompareLL(&expectedLL, dedupedLL); err != nil {
		t.Error("Buf Test Failed")
	}

	dedupedLL = removeDuplicateNoBuf(&ll)
	dedupedLL.PrintLinkedList()

	if err := singleLinkedList.CompareLL(&expectedLL, dedupedLL); err != nil {
		t.Error("No Buf Test Failed")
	}
}

func buildLinkedList(buildElements []int8) singleLinkedList.Node {
	ll := singleLinkedList.InstantiateLinkedList()
	for _, num := range buildElements {
		ll.InsertElement(num)
	}
	return ll
}
