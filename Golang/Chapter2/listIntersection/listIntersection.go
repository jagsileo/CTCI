package listIntersection

import (
	"github.com/CTCI/Golang/Chapter2/singleLinkedList"
)

func getListIntersection(llA *singleLinkedList.Node, llB *singleLinkedList.Node) *singleLinkedList.Node {
	n1 := llA
	n2 := llB

	isFirst := true

	for {
		if isFirst == false && n1 == llA && n2 == llB {
			return nil
		}
		if err := singleLinkedList.CompareLL(n1, n2); err != nil {
			if n1.Next != nil {
				n1 = n1.Next
			} else {
				n1 = llA
			}

			if n2.Next != nil {
				n2 = n2.Next
			} else {
				n2 = llB
			}

			isFirst = false

		} else {
			return n1
		}

	}

	return nil
}
