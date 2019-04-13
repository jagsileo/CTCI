package isUnique

import "testing"

func TestSortString(t *testing.T) {
	strA := "algorithm"

	if sortString(strA) != "aghilmort" {
		t.Error("Expected sorted string aghilmort for algorithm")
	}

	strB := "malayalam"
	if sortString(strB) != "aaaallmmy" {
		t.Error("Expected sorted string aaaallmmy for malayalam")
	}

}

func TestIsUnique(t *testing.T) {
	strA := "algorithm"

	if !isUnique(strA) {
		t.Error("Expected true")
	}

	strB := ""

	if !isUnique(strB) {
		t.Error("Expected true")
	}

	strC := "abba"

	if isUnique(strC) {
		t.Error("Expected false")
	}
}
