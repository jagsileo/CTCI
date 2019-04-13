package oneAway

func isOneAway(strA string, strB string) bool {
	var charArr [256]int
	lenA := len(strA)
	strRuneA := []rune(strA)

	lenB := len(strB)
	strRuneB := []rune(strB)

	for i := 0; i < lenA; i++ {
		charArr[int(strRuneA[i])] = 1
	}

	for i := 0; i < lenB; i++ {
		charArr[int(strRuneB[i])] = 0
	}

	total := 0
	for i := 0; i < 256; i++ {
		total += charArr[i]
	}

	if lenA-total < 3 {
		return false
	}

	return true
}
