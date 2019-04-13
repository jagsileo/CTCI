package palindromePermutation

import "strings"

func isPalindromePermutation(str string) bool {
	if str == "" || len(str) == 1 {
		return true
	}

	var countVector [256]bool
	str = strings.Replace(str, " ", "", -1)
	str = strings.ToLower(str)
	strRune := []rune(str)
	var doubles int
	l := len(str)

	for i := 0; i < l; i++ {
		asciiVal := int(strRune[i])
		if countVector[asciiVal] == true {
			doubles++
			countVector[asciiVal] = false
		} else {
			countVector[asciiVal] = true
		}

	}

	if doubles == l/2 {
		return true
	}
	return false
}
