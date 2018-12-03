package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
)

func main() {
	file, err := os.Open("resources/Day3_1")

	if err == nil {
		fmt.Printf("Result of part1: %d\n", part1(file))
		_, _ = file.Seek(0, 0)
		fmt.Printf("Result of part2: %s\n", part2(file))
		_ = file.Close()
	} else {
		fmt.Printf("Error: %s", err.Error())
	}
}

type point struct {
	x int
	y int
}

func part1(input io.Reader) int {
	inputScanner := bufio.NewScanner(input)
	points := make(map[point]int)
	overlap := 0
	lineRegexp := regexp.MustCompile(" @ |,|: |x")
	for inputScanner.Scan() {
		splitLine := lineRegexp.Split(inputScanner.Text(), -1)
		xPoint, err := strconv.Atoi(splitLine[1])
		if err != nil {
			panic("Unable to parse X coordinate")
		}
		yPoint, err := strconv.Atoi(splitLine[2])
		if err != nil {
			panic("Unable to parse Y coordinate")
		}
		xDistance, err := strconv.Atoi(splitLine[3])
		if err != nil {
			panic("Unable to parse X distance")
		}
		yDistance, err := strconv.Atoi(splitLine[4])
		if err != nil {
			panic("Unable to parse Y distance")
		}
		for i := xPoint; i < xDistance+xPoint; i++ {
			for j := yPoint; j < yDistance+yPoint; j++ {
				if points[point{i, j}] == 1 {
					overlap++
				}
				points[point{i, j}]++
			}
		}
	}

	return overlap
}

func part2(input io.Reader) string {
	inputScanner := bufio.NewScanner(input)
	points := make(map[point][]string)
	claims := make(map[string]bool)
	lineRegexp := regexp.MustCompile(" @ |,|: |x")
	for inputScanner.Scan() {
		splitLine := lineRegexp.Split(inputScanner.Text(), -1)
		claims[splitLine[0]] = true
		xPoint, err := strconv.Atoi(splitLine[1])
		if err != nil {
			panic("Unable to parse X coordinate")
		}
		yPoint, err := strconv.Atoi(splitLine[2])
		if err != nil {
			panic("Unable to parse Y coordinate")
		}
		xDistance, err := strconv.Atoi(splitLine[3])
		if err != nil {
			panic("Unable to parse X distance")
		}
		yDistance, err := strconv.Atoi(splitLine[4])
		if err != nil {
			panic("Unable to parse Y distance")
		}
		for i := xPoint; i < xDistance+xPoint; i++ {
			for j := yPoint; j < yDistance+yPoint; j++ {
				points[point{i, j}] = append(points[point{i, j}], splitLine[0])
				if (len(points[point{i, j}]) > 1) {
					for _, claim := range points[point{i, j}] {
						claims[claim] = false
					}
				}
			}
		}
	}

	for key, value := range claims {
		if value {
			return key
		}
	}

	panic("Everything overlapped")
}
