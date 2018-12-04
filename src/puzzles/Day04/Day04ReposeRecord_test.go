package main

import (
	"strings"
	"testing"
)

func TestOfExample(t *testing.T) {
	result := part1(strings.NewReader("[1518-11-01 00:00] Guard #10 begins shift\n" +
		"[1518-11-01 00:05] falls asleep\n" +
		"[1518-11-01 00:25] wakes up\n" +
		"[1518-11-01 00:30] falls asleep\n" +
		"[1518-11-01 00:55] wakes up\n" +
		"[1518-11-01 23:58] Guard #99 begins shift\n" +
		"[1518-11-02 00:40] falls asleep\n" +
		"[1518-11-02 00:50] wakes up\n" +
		"[1518-11-03 00:05] Guard #10 begins shift\n" +
		"[1518-11-03 00:24] falls asleep\n" +
		"[1518-11-03 00:29] wakes up\n" +
		"[1518-11-04 00:02] Guard #99 begins shift\n" +
		"[1518-11-04 00:36] falls asleep\n" +
		"[1518-11-04 00:46] wakes up\n" +
		"[1518-11-05 00:03] Guard #99 begins shift\n" +
		"[1518-11-05 00:45] falls asleep\n" +
		"[1518-11-05 00:55] wakes up"))
	if result != 240 {
		t.Errorf("Expected %d but was %d", 240, result)
	}
}

func TestOfExampleAgain(t *testing.T) {
	result := part2(strings.NewReader("[1518-11-01 00:00] Guard #10 begins shift\n" +
		"[1518-11-01 00:05] falls asleep\n" +
		"[1518-11-01 00:25] wakes up\n" +
		"[1518-11-01 00:30] falls asleep\n" +
		"[1518-11-01 00:55] wakes up\n" +
		"[1518-11-01 23:58] Guard #99 begins shift\n" +
		"[1518-11-02 00:40] falls asleep\n" +
		"[1518-11-02 00:50] wakes up\n" +
		"[1518-11-03 00:05] Guard #10 begins shift\n" +
		"[1518-11-03 00:24] falls asleep\n" +
		"[1518-11-03 00:29] wakes up\n" +
		"[1518-11-04 00:02] Guard #99 begins shift\n" +
		"[1518-11-04 00:36] falls asleep\n" +
		"[1518-11-04 00:46] wakes up\n" +
		"[1518-11-05 00:03] Guard #99 begins shift\n" +
		"[1518-11-05 00:45] falls asleep\n" +
		"[1518-11-05 00:55] wakes up"))
	if result != 4455 {
		t.Errorf("Expected %d but was %d", 4455, result)
	}
}
