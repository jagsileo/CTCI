package listIntersection

import (
	"testing"

	"github.com/CTCI/Golang/Chapter2/singleLinkedList"
)

func TestGetListIntersection(t *testing.T) {
	llA := buildLinkedList([]int8{2, 5, 3, 9, 6, 8, 9})
	llB := buildLinkedList([]int8{6, 7, 6, 8, 9})
	expectedResult := buildLinkedList([]int8{6, 8, 9})

	intersectNode := getListIntersection(&llA, &llB)

	if err := singleLinkedList.CompareLL(intersectNode, &expectedResult); err != nil {
		t.Error("Test Failed #1")
	}

	llB = buildLinkedList([]int8{1, 5, 6, 7, 6, 8, 9})

	intersectNode = getListIntersection(&llA, &llB)
	if err := singleLinkedList.CompareLL(intersectNode, &expectedResult); err != nil {
		t.Error("Test Failed #2")
	}

	llA = buildLinkedList([]int8{1, 5, 6, 7, 6, 8, 9})
	expectedResult = llA

	intersectNode = getListIntersection(&llA, &llB)

	if err := singleLinkedList.CompareLL(intersectNode, &expectedResult); err != nil {
		t.Error("Test Failed #3")
	}

	llA = buildLinkedList([]int8{1, 2, 3, 4, 5, 6, 7})

	intersectNode = getListIntersection(&llA, &llB)

	if intersectNode != nil {
		t.Error("Test Failed #4")
	}

}

func buildLinkedList(buildElements []int8) singleLinkedList.Node {
	ll := singleLinkedList.InstantiateLinkedList()
	for _, num := range buildElements {
		ll.InsertElement(num)
	}
	return ll
}
