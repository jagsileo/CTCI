package stringRotation

import (
	"log"
)

func isStringRotation(strA string, strB string) bool {
	lenA := len(strA)
	lenB := len(strB)

	if lenA != lenB {
		return false
	}

	strQueue := []rune(strA)
	counter := 0
	for counter != lenA && string(strQueue) != strB {
		counter++
		char := dequeue(strQueue)
		strQueue = enqueue(strQueue, char)

	}

	if counter == lenA {
		return false
	}
	return true
}

func dequeue(strRune []rune) rune {
	buf := strRune[0]
	l := len(strRune)
	for i := 0; i < l-1; i++ {
		strRune[i] = strRune[i+1]
	}
	strRune[l-1] = buf
	log.Println(string(strRune))
	return buf
}

func enqueue(strRune []rune, char rune) []rune {

	strRune = append(strRune, char)
	return strRune
}
