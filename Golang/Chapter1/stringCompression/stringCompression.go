package stringCompression

import (
	"strconv"
	"strings"
)

func compressString(str string) string {
	var b strings.Builder
	strRune := rune(str[0])
	currStrCount := 1
	l := len(str)

	for i := 1; i < l; i++ {

		if rune(str[i]) == strRune {
			currStrCount++
		} else {
			b.WriteString(string(strRune))
			b.WriteString(strconv.Itoa(currStrCount))
			strRune = rune(str[i])
			currStrCount = 1
		}
	}
	b.WriteString(string(strRune))
	b.WriteString(strconv.Itoa(currStrCount))
	cStr := b.String()
	if len(cStr) < l {
		return cStr
	}
	return str
}
