package deleteMiddleElement

import "github.com/CTCI/Golang/Chapter2/singleLinkedList"

func deleteMiddleElement(ll *singleLinkedList.Node) *singleLinkedList.Node {
	n := ll
	var prev *singleLinkedList.Node
	prev = nil
	for n.Next != nil {
		n.Value = n.Next.Value
		prev = n
		n = n.Next
	}

	prev.Next = nil

	return ll
}
