package stringRotation

import (
	"testing"
)

func TestIsStringRotation(t *testing.T) {
	strA := "waterbottle"
	strB := "erbottlewat"

	if !isStringRotation(strA, strB) {
		t.Error("Test Failed for waterbottle -> erbottlewat")
	}

	strB = "erbottelatw"
	if isStringRotation(strA, strB) {
		t.Error("Test Failed for waterbottle -> erbottlewat")
	}
}
