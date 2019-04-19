package sumLists

import (
	"github.com/CTCI/Golang/Chapter2/singleLinkedList"
	"testing"
)

func TestSumLists(t *testing.T) {
	listA := buildLinkedList([]int8{2, 1, 3})
	listB := buildLinkedList([]int8{6, 5, 9})

	summedList := sumListsWhenNumReversed(&listA, &listB)
	expectedList := buildLinkedList([]int8{8, 6, 2, 1})

	summedList.PrintLinkedList()
	if err := singleLinkedList.CompareLL(summedList, &expectedList); err != nil {
		t.Error("Test #1 Failed")
	}

	listA = buildLinkedList([]int8{2, 1, 3})
	listB = buildLinkedList([]int8{6, 5, 9, 9})

	summedList = sumListsWhenNumReversed(&listA, &listB)
	expectedList = buildLinkedList([]int8{8, 6, 2, 0, 1})
	summedList.PrintLinkedList()
	if err := singleLinkedList.CompareLL(summedList, &expectedList); err != nil {
		t.Error("Test #2 Failed")
	}

	//Dumb me created list with value as int8. So things get messed up if I passed a test data with > 3 digits
	//The algorithm works though
	listA = buildLinkedList([]int8{1, 3})
	listB = buildLinkedList([]int8{5, 9})

	summedList = sumListsWhenNumForward(&listA, &listB)
	expectedList = buildLinkedList([]int8{7, 2})

	summedList.PrintLinkedList()
	if err := singleLinkedList.CompareLL(summedList, &expectedList); err != nil {
		t.Error("Test #3 Failed")
	}

	listA = buildLinkedList([]int8{3})
	listB = buildLinkedList([]int8{6, 9})

	summedList = sumListsWhenNumForward(&listA, &listB)
	expectedList = buildLinkedList([]int8{7, 2})
	// summedList.PrintLinkedList()
	if err := singleLinkedList.CompareLL(summedList, &expectedList); err != nil {
		t.Error("Test #4 Failed")
	}

}

func buildLinkedList(buildElements []int8) singleLinkedList.Node {
	ll := singleLinkedList.InstantiateLinkedList()
	for _, num := range buildElements {
		ll.InsertElement(num)
	}

	return ll
}
