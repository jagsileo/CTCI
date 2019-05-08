package stackFindMin

import "testing"

func TestStackFindMin(t *testing.T) {
	arr := []int{2, 7, 1, 3, 5}
	minVal := stackFindMin(arr)
	if minVal != 1 {
		t.Error("Test Failed #1")
	}

	arr = []int{1, 2, 3, 4, 5}
	minVal = stackFindMin(arr)
	if minVal != 1 {
		t.Error("Test Failed #2")
	}

}
