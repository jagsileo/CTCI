package listPalindrome

import (
	"testing"

	"github.com/CTCI/Golang/Chapter2/singleLinkedList"
)

func TestListPalindrome(t *testing.T) {
	llA := buildLinkedList([]int8{2, 5, 8, 5, 2})
	if !isListPalindrome(&llA) {
		t.Error("Test Failed for #1")
	}

	llB := buildLinkedList([]int8{2, 5, 5, 2})
	if !isListPalindrome(&llB) {
		t.Error("Test Failed for #2")
	}

	llC := buildLinkedList([]int8{6, 5, 5, 8, 9})
	if isListPalindrome(&llC) {
		t.Error("Test Failed for #3")
	}
}

func buildLinkedList(buildElements []int8) singleLinkedList.Node {
	ll := singleLinkedList.Node{}
	for _, num := range buildElements {
		ll.InsertElement(num)
	}

	return ll
}
