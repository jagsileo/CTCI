package listPalindrome

import (
	"github.com/CTCI/Golang/Chapter2/singleLinkedList"
)

func isListPalindrome(ll *singleLinkedList.Node) bool {

	n := ll
	len := getLen(ll)

	ctr := 1

	for n != nil && ctr <= len/2 {
		n = n.Next
		ctr++
	}

	n1 := ll

	for n != nil {
		if n.Value != 
	}
	return false
}

func getLen(ll *singleLinkedList.Node) int {
	n := ll
	len := 0

	for n != nil {
		n = n.Next
		len++
	}

	return len
}
