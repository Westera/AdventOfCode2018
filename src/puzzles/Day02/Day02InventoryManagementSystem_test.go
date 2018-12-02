package main

import (
	"strings"
	"testing"
)

func TestChecksumExample(t *testing.T) {
	result := part1(strings.NewReader("abcdef\nbababc\nabbcde\nabcccd\naabcdd\nabcdee\nababab"))
	if result != 12 {
		t.Errorf("Expected %d but was %d", 12, result)
	}
}

func TestFindingTheBoxesExample(t *testing.T) {
	result := part2(strings.NewReader("abcde\nfghij\nklmno\npqrst\nfguij\naxcye\nwvxyz"))
	if result != "fgij" {
		t.Errorf("Expected %s but was %s", "fgij", result)
	}
}
