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

/*

	This is an interesting case study to understand pointers in Go
	If a reference variable is copied to another variable and the variable is changed,
	then the original reference is changed - agreed
	So when I reverse a linked list, I copy the dereferenced variable - see line 33, with the understanding that
	if I modify the copy , the original variable will be unaffected.
	But it does affect, my reasoning - even though we dereference and make a copy in line 33,
	it dereferences the starting head node ,
	the "Next" still carries the original "Next" reference.


	To test this, instead of just making a copy of the dereferenced head, I loop the original linked list and copy their values
	- see line 68 - 72
	This seems to resolve the issue!



*/

// func reverseLL(input *singleLinkedList.Node) singleLinkedList.Node {
// 	prev := &singleLinkedList.Node{}
// 	prev = nil
// 	inputCopy := *input
// 	curr := &inputCopy
// 	next := curr.Next
// 	for next != nil {

// 		curr.Next = prev
// 		prev = curr
// 		curr = next
// 		next = curr.Next

// 	}
// 	curr.Next = prev
// 	print("\n<><><><>\n")
// 	input.PrintLinkedList()
// 	print("REVERSED ....\n")
// 	curr.PrintLinkedList()
// 	return *curr
// }

func reverseLL(input *singleLinkedList.Node) singleLinkedList.Node {
	inputCopy := &singleLinkedList.Node{}
	prev := &singleLinkedList.Node{}
	prev = nil
	n := input
	nCopy := inputCopy

	for n != nil {
		nCopy.InsertElement(n.Value)
		n = n.Next

	}
	curr := inputCopy
	next := curr.Next
	for next != nil {

		curr.Next = prev
		prev = curr
		curr = next
		next = curr.Next

	}
	curr.Next = prev
	print("\n<><><><>\n")
	input.PrintLinkedList()
	print("REVERSED ....\n")
	curr.PrintLinkedList()
	return *curr

}
