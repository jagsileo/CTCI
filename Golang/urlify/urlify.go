package urlify

func urlify(str string, trueLen int) string {
	strRune := []rune(str)
	spaceRune := []rune(" ")

	idx := len(str) - 1

	for i := trueLen - 1; i >= 0; i-- {
		if strRune[i] != spaceRune[0] {
			strRune[idx] = strRune[i]
			idx--
		} else {
			strRune[idx] = '0'
			strRune[idx-1] = '2'
			strRune[idx-2] = '%'
			idx = idx - 3

		}

	}

	return string(strRune)
}
