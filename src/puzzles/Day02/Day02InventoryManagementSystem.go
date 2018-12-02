package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("resources/Day2_1")

	if err == nil {
		fmt.Printf("Result of part1: %d\n", part1(file))
		_, _ = file.Seek(0, 0)
		fmt.Printf("Result of part2: %s\n", part2(file))
		_ = file.Close()
	} else {
		fmt.Printf("Error: %s", err.Error())
	}
}

func part1(input io.Reader) int {
	inputScanner := bufio.NewScanner(input)
	amountOfTwos := 0
	amountOfThrees := 0
	for inputScanner.Scan() {
		line := inputScanner.Text()
		found2 := false
		found3 := false
		for _, char := range line {
			numberOfChar := strings.Count(line, string(char))
			if numberOfChar == 2 && !found2 {
				amountOfTwos++
				found2 = true
			} else if numberOfChar == 3 && !found3 {
				amountOfThrees++
				found3 = true
			}
			line = strings.Replace(line, string(char), "", -1)
		}
	}
	return amountOfTwos * amountOfThrees
}

func part2(input io.Reader) string {
	inputScanner := bufio.NewScanner(input)
	var inputs []string
	for inputScanner.Scan() {
		line := inputScanner.Text()
		for _, prevInput := range inputs {
			nrOfDiffs := 0
			var potentialResult bytes.Buffer
			for index := range line {
				if line[index] != prevInput[index] {
					nrOfDiffs++
					if nrOfDiffs > 1 {
						break
					}
				} else {
					potentialResult.WriteString(string(line[index]))
				}

			}
			if nrOfDiffs == 1 {
				return potentialResult.String()
			}
		}
		inputs = append(inputs, line)
	}
	panic("No solution to problem")
}
