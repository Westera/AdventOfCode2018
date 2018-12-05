package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("resources/Day5_1")

	if err == nil {
		fmt.Printf("Result of part1: %d\n", part1(file))
		_, _ = file.Seek(0, 0)
		fmt.Printf("Result of part2: %d\n", part2(file))
		_ = file.Close()
	} else {
		fmt.Printf("Error: %s", err.Error())
	}
}

func part1(input io.Reader) int {
	inputReader := bufio.NewReader(input)

	var result []rune

	for {
		if c1, _, err := inputReader.ReadRune(); err != nil {
			if err == io.EOF {
				break
			} else {
				panic(err)
			}
		} else {
			if len(result) > 0 && (result[len(result)-1] < 97 && result[len(result)-1]+32 == c1 || result[len(result)-1] > 96 && result[len(result)-1]-32 == c1) {
				result = result[:len(result)-1]
			} else {
				result = append(result, c1)
			}
		}
	}

	return len(result)
}

func part2(input io.Reader) int {
	var result []rune
	smallest := -1

	for ignored := rune('A'); ignored <= rune('Z'); ignored++ {
		inputReader := bufio.NewReader(input)
		for {
			if c1, _, err := inputReader.ReadRune(); err != nil {
				if err == io.EOF {
					break
				} else {
					panic(err)
				}
			} else {
				if len(result) > 0 && (result[len(result)-1] < 97 && result[len(result)-1]+32 == c1 || result[len(result)-1] > 96 && result[len(result)-1]-32 == c1) {
					result = result[:len(result)-1]
				} else if c1 != ignored && c1-32 != ignored {
					result = append(result, c1)
				}
			}
		}
		if smallest > len(result) || smallest == -1 {
			smallest = len(result)
		}
		result = result[:0]
		seeker, ok := input.(io.Seeker)
		if ok {
			_, _ = seeker.Seek(0, 0)
		}
	}

	return smallest
}
