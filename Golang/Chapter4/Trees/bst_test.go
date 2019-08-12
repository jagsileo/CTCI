package bst

import (
	"testing"
)

func TestBSTOperations(t *testing.T) {
	arr := []int{2, 5, 1, 4, 8, 6}

	bstOperations(arr)
}

func TestCalcHeight(t *testing.T) {
	arr := []int{2, 5, 1, 4, 8, 6}

	n := bstOperations(arr)

	h := n.calcHeight()

	if h != 6 {
		t.Error("Test failed - Incorrect height")
	}
}
