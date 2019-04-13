package stringCompression

import (
	"testing"
)

func TestCompressString(t *testing.T) {
	strA := "aabcccccaaa"

	if compressString(strA) != "a2b1c5a3" {
		t.Error("Test Failed for aabcccccaaa")
	}

	strA = "abc"

	if compressString(strA) != "abc" {
		t.Error("Test Failed for abc")
	}
}
