package detectLoop

import "github.com/CTCI/Golang/Chapter2/singleLinkedList"

func detectLoop(ll *singleLinkedList.Node) *singleLinkedList.Node {
	n1 := ll
	n2 := ll
	isFirst := true

	for n1 != nil && n2 != nil {
		if !isFirst && n1 == n2 {
			break
		}

		n1 = n1.Next
		if n2.Next != nil {
			n2 = n2.Next.Next
		}

		isFirst = false
	}

	n1 = ll
	for n1 != nil && n2 != nil {

		if n1 == n2 {
			return n1
		}
		n1 = n1.Next
		n2 = n2.Next
	}
	return nil

}
