package kToLastElements

import "github.com/CTCI/Golang/Chapter2/singleLinkedList"

func kToLastElements(ll *singleLinkedList.Node, k int) *singleLinkedList.Node {
	idx := 0
	n := ll

	for n != nil && idx < k {
		idx++
		n = n.Next
	}

	return n
}
