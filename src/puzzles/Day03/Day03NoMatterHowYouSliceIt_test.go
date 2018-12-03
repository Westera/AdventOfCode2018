package main

import (
	"strings"
	"testing"
)

func TestOfExample(t *testing.T) {
	result := part1(strings.NewReader("#1 @ 1,3: 4x4\n#2 @ 3,1: 4x4\n#3 @ 5,5: 2x2"))
	if result != 4 {
		t.Errorf("Expected %d but was %d", 4, result)
	}
}

func TestOfExampleAgain(t *testing.T) {
	result := part2(strings.NewReader("#1 @ 1,3: 4x4\n#2 @ 3,1: 4x4\n#3 @ 5,5: 2x2"))
	if result != "#3" {
		t.Errorf("Expected %s but was %s", "#3", result)
	}
}
