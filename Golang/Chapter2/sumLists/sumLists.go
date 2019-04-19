package sumLists

import "github.com/CTCI/Golang/Chapter2/singleLinkedList"

func sumListsWhenNumReversed(listA *singleLinkedList.Node, listB *singleLinkedList.Node) *singleLinkedList.Node {
	a := listA
	b := listB
	result := singleLinkedList.InstantiateLinkedList()
	c := &result
	var quotient, sum int8

	for a != nil || b != nil || quotient != 0 {
		sum += quotient
		if a != nil {
			sum += a.Value
			a = a.Next
		}
		if b != nil {
			sum += b.Value
			b = b.Next
		}

		c.Next = &singleLinkedList.Node{
			Value: sum % 10,
			Next:  nil,
		}
		quotient = sum / 10
		c = c.Next
		sum = 0
	}
	return result.Next
}
