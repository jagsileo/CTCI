package stackFindMin

type sampleStack struct {
	item int
}

func stackFindMin(arr []int) int {
	ss := []sampleStack{}
	topPtr := 0
	maxLen := len(arr) - 1
	min := arr[0]
	for _, num := range arr {
		push(num, &ss, &topPtr, maxLen, &min)
	}

	return min
}

func push(i int, ssPtr *[]sampleStack, topPtr *int, maxLen int, minPtr *int) {
	if *topPtr > maxLen {
		return
	}
	print(*topPtr)
	print("\n")
	ss := *ssPtr
	ss = append(ss, sampleStack{item: i})
	*topPtr = *topPtr + 1
	if *minPtr > i {
		*minPtr = i
	}
}
