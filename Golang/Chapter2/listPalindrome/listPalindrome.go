package listPalindrome

import (
	"github.com/CTCI/Golang/Chapter2/singleLinkedList"
)

func isListPalindrome(ll *singleLinkedList.Node) bool {

	rev := reverseLL(ll)

	if err := singleLinkedList.CompareLL(&rev, ll); err != nil {
		return false
	}

	return true
}

func reverseLL(input *singleLinkedList.Node) singleLinkedList.Node {
	prev := &singleLinkedList.Node{}
	prev = nil
	inputCopy := *input
	curr := &inputCopy
	next := curr.Next
	for next != nil {

		curr.Next = prev
		prev = curr
		curr = next
		next = curr.Next

	}
	curr.Next = prev
	print("REVERSED ....\n")
	curr.PrintLinkedList()
	return *curr
}
