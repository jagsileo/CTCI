package isUnique

import (
	"sort"
)

type sortRune []rune

func (s sortRune) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRune) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRune) Len() int {
	return len(s)
}

func isUnique(str string) bool {

	if str == "" || len(str) == 1 {
		return true
	}
	sortStrA := sortString(str)
	for i := 0; i <= len(str)-2; i++ {
		if sortStrA[i] == sortStrA[i+1] {
			return false
		}
	}

	return true
}

func sortString(s string) string {
	r := []rune(s)
	sort.Sort(sortRune(r))
	return string(r)
}
