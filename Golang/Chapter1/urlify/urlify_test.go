package urlify

import (
	"log"
	"testing"
)

func TestUrlify(t *testing.T) {
	str := "Mr John Smith    "
	trueLen := 13

	urlified := urlify(str, trueLen)
	log.Println(urlified)
	if urlified != "Mr%20John%20Smith" {
		t.Error("Test Failed for Mr John Smith")
	}

	str = "I am going to become a googler soon              "

	urlified = urlify(str, 35)
	log.Println(urlified)
	if urlified != "I%20am%20going%20to%20become%20a%20googler%20soon" {
		t.Error("Test Failed for I am going to be a googler soon")
	}
}
