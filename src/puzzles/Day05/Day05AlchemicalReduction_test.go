package main

import (
	"strings"
	"testing"
)

func TestOfSimple(t *testing.T) {
	result := part1(strings.NewReader("aA"))
	if result != 0 {
		t.Errorf("Expected %d but was %d", 0, result)
	}
}

func TestOfSimpleChaine(t *testing.T) {
	result := part1(strings.NewReader("aBbA"))
	if result != 0 {
		t.Errorf("Expected %d but was %d", 0, result)
	}
}

func TestOfNoPairs(t *testing.T) {
	result := part1(strings.NewReader("aBAb"))
	if result != 4 {
		t.Errorf("Expected %d but was %d", 4, result)
	}
}

func TestOfNoPairsButSameCharacters(t *testing.T) {
	result := part1(strings.NewReader("aaBBAAbb"))
	if result != 8 {
		t.Errorf("Expected %d but was %d", 8, result)
	}
}

func TestOfLongExample(t *testing.T) {
	result := part1(strings.NewReader("dabAcCaCBAcCcaDA"))
	if result != 10 {
		t.Errorf("Expected %d but was %d", 10, result)
	}
}

func TestOfLongExamplePart2(t *testing.T) {
	result := part2(strings.NewReader("dabAcCaCBAcCcaDA"))
	if result != 4 {
		t.Errorf("Expected %d but was %d", 4, result)
	}
}
