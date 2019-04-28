package partitionLL

import (
	"testing"

	"github.com/CTCI/Golang/Chapter2/singleLinkedList"
)

func TestPartitionLL(t *testing.T) {
	ll := buildLinkedList([]int8{3, 5, 8, 1, 5, 10, 2, 1, 3, 7, 8})
	n := ll
	n.PrintLinkedList()
	print("\n<><><><>")
	expectedLL := buildLinkedList([]int8{1, 2, 1, 3, 3, 5, 8, 5, 10, 7, 8})
	partitioned := partitionLL(&ll, 5)
	partitioned.PrintLinkedList()

	if err := singleLinkedList.CompareLL(&expectedLL, partitioned); err != nil {
		t.Error("Buf Test Failed")
	}
}

func buildLinkedList(buildElements []int8) singleLinkedList.Node {
	ll := singleLinkedList.Node{}
	for _, num := range buildElements {
		ll.InsertElement(num)
	}

	return ll
}
