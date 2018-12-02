package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("resources/Day1_1")

	if err == nil {
		fmt.Printf("Result of part1: %d\n", part1(file))
		_, _ = file.Seek(0, 0)
		fmt.Printf("Result of part2: %d\n", part2(file, make(map[int]bool), 0))
		_ = file.Close()
	} else {
		fmt.Printf("Error: %s", err.Error())
	}
}

func part1(input io.Reader) int {
	inputScanner := bufio.NewScanner(input)
	sum := 0
	for inputScanner.Scan() {
		value, err := strconv.Atoi(inputScanner.Text())
		if err == nil {
			sum += value
		}
	}
	return sum
}

func part2(input io.Reader, set map[int]bool, sum int) int {
	inputScanner := bufio.NewScanner(input)
	set[0] = true
	for inputScanner.Scan() {
		value, err := strconv.Atoi(inputScanner.Text())
		if err == nil {
			sum += value
			if !set[sum] {
				set[sum] = true
			} else {
				return sum
			}
		}
	}

	seeker, ok := input.(io.Seeker)
	if ok {
		_, _ = seeker.Seek(0, 0)
	}

	return part2(input, set, sum)
}
