package removeDuplicate

import "github.com/CTCI/Golang/Chapter2/singleLinkedList"

type void struct{}

func removeDuplicate(ll *singleLinkedList.Node) *singleLinkedList.Node {
	buf := make(map[int8]void)
	counter := 0
	n := ll
	for n.Next != nil {
		counter++
		if _, ok := buf[n.Value]; !ok {
			buf[n.Value] = void{}
		} else {
			n = n.DeleteElement(n.Value, counter+1)
		}

		n = n.Next
	}

	return ll
}
