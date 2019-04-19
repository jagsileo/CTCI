package sumLists

import "github.com/CTCI/Golang/Chapter2/singleLinkedList"

func sumLists(listA *singleLinkedList.Node, listB *singleLinkedList.Node) *singleLinkedList.Node {
	a := listA
	b := listB
	result := singleLinkedList.InstantiateLinkedList()
	c := &result
	var quotient, sum int8

	for a != nil || b != nil {
		if a != nil {
			sum += a.Value
		}
		if b != nil {
			sum += b.Value
		}

		c.Next = &singleLinkedList.Node{
			Value: (quotient + sum) % 10,
			Next:  nil,
		}
		quotient = sum / 10
		a = a.Next
		b = b.Next
		c = c.Next
	}
	return result.Next
}
