package palindromePermutation

import "testing"

func TestIsPalindromePermutation(t *testing.T) {
	str := "Tact Coa"

	if !isPalindromePermutation(str) {
		t.Error("Test Failed for Tact Coa")
	}

	str = "maya lalam"

	if !isPalindromePermutation(str) {
		t.Error("Test Failed for maya lalam")
	}
}
