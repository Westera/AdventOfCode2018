package main

import (
	"strings"
	"testing"
)

func TestOnlyPositive(t *testing.T) {
	result := part1(strings.NewReader("+1\n+1\n+1\n"))
	if result != 3 {
		t.Errorf("Expected %d but was %d", 3, result)
	}
}

func TestOnlyNegative(t *testing.T) {
	result := part1(strings.NewReader("-1\n-2\n-3"))
	if result != -6 {
		t.Errorf("Expected %d but was %d", -6, result)
	}
}

func TestMixed(t *testing.T) {
	result := part1(strings.NewReader("+1\n+1\n-2"))
	if result != 0 {
		t.Errorf("Expected %d but was %d", 0, result)
	}
}

func TestFirst0(t *testing.T) {
	result := part2(strings.NewReader("+1\n-1"), make(map[int]bool), 0)
	if result != 0 {
		t.Errorf("Expected %d but was %d", 0, result)
	}
}

func TestFirst10(t *testing.T) {
	result := part2(strings.NewReader("+3\n+3\n+4\n-2\n-4"), make(map[int]bool), 0)
	if result != 10 {
		t.Errorf("Expected %d but was %d", 10, result)
	}
}

func TestFirst5(t *testing.T) {
	result := part2(strings.NewReader("-6\n+3\n+8\n+5\n-6"), make(map[int]bool), 0)
	if result != 5 {
		t.Errorf("Expected %d but was %d", 5, result)
	}
}

func TestFirst14(t *testing.T) {
	result := part2(strings.NewReader("+7\n+7\n2\n-7\n-4"), make(map[int]bool), 0)
	if result != 14 {
		t.Errorf("Expected %d but was %d", 14, result)
	}
}
