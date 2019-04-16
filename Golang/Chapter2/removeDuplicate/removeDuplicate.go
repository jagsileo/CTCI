package removeDuplicate

import "github.com/CTCI/Golang/Chapter2/singleLinkedList"

type void struct{}

func removeDuplicate(ll *singleLinkedList.Node) *singleLinkedList.Node {
	buf := make(map[int8]void)
	n := ll
	var prev *singleLinkedList.Node
	prev = nil
	for n != nil {
		if _, ok := buf[n.Value]; !ok {
			buf[n.Value] = void{}
			prev = n
		} else {
			prev.Next = n.Next
		}

		n = n.Next
	}

	return ll
}
