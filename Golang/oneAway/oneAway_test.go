package oneAway

import "testing"

func TestOneAway(t *testing.T) {
	strA := "pale"
	strB := "pal"

	if !isOneAway(strA, strB) {
		t.Error("Test Failed for pale, pal")
	}

	strB = "bake"
	if isOneAway(strA, strB) {
		t.Error("Test Failed for pale, bake")
	}

	strB = "bale"
	if !isOneAway(strA, strB) {
		t.Error("Test Failed for pale, bale")
	}
}
