package detectLoop

import (
	"testing"

	"github.com/CTCI/Golang/Chapter2/singleLinkedList"
)

func TestDetectLoop(t *testing.T) {
	ll := buildLinkedList([]int8{2, 5, 6, 7, 8})
	loopedLL := buildLoopedLinkedList(3, &ll)
	expectedLL := buildLinkedList([]int8{6, 7, 8})

	detectedLoop := detectLoop(loopedLL)
	unloop := unloopLL(detectedLoop)
	if err := singleLinkedList.CompareLL(unloop, &expectedLL); err != nil {
		t.Error("Test Failed #1")
	}

	ll = buildLinkedList([]int8{2, 5, 6, 7, 8})
	detectedLoop = detectLoop(loopedLL)

	if detectedLoop != nil {
		t.Error("Test Failed #2")
	}

}

func buildLoopedLinkedList(idx int, ll *singleLinkedList.Node) *singleLinkedList.Node {
	counter := 1
	n := ll
	buf := &singleLinkedList.Node{}

	for n.Next != nil {
		if counter == idx {
			buf = n
		}

		counter++
		n = n.Next
	}
	n.Next = buf
	return ll
}

func buildLinkedList(buildElements []int8) singleLinkedList.Node {
	ll := singleLinkedList.InstantiateLinkedList()
	for _, num := range buildElements {
		ll.InsertElement(num)
	}
	return ll
}

func unloopLL(ll *singleLinkedList.Node) *singleLinkedList.Node {
	buf := make(map[int8]bool)
	n := ll
	buf[n.Value] = true
	for n.Next != nil {
		if _, ok := buf[n.Next.Value]; !ok {
			buf[n.Next.Value] = true
			n = n.Next
		} else {
			n.Next = nil
		}

	}

	return ll
}
