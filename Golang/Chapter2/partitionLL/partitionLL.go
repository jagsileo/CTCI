package partitionLL

import (
	"github.com/CTCI/Golang/Chapter2/singleLinkedList"
)

func partitionLL(ll *singleLinkedList.Node, k int8) *singleLinkedList.Node {

	n := ll

	idx := -1
	counter := -1

	for n != nil {
		counter++

		if n.Value == 5 && idx > -1 {
			idx = min(idx, counter)
		} else if n.Value == 5 && idx == -1 {
			idx = counter
		}
		n = n.Next
	}
	n = ll
	lessHead := &singleLinkedList.Node{
		Value: 0,
		Next:  nil,
	}
	less := lessHead
	prev := n
	counter = 0

	for n != nil {
		counter++
		if n.Value < k && counter > idx {
			less.Next = n
			prev.Next = n.Next
			less = less.Next
			less.Next = nil
			n = prev.Next
		} else {
			prev = n
			n = n.Next
		}
	}

	less.Next = ll
	ll = lessHead

	return ll

}

func min(x, y int) int {
	if x <= y {
		return x
	}
	return y
}
