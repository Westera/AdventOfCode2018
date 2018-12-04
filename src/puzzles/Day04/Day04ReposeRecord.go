package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("resources/Day4_1")

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
	inputScanner := bufio.NewScanner(input)
	var schedule []string

	for inputScanner.Scan() {
		line := inputScanner.Text()
		schedule = append(schedule, line)
	}

	sort.Strings(schedule)

	guardsSleepSum := make(map[string]int)
	guardsSleepyMinutes := make(map[string]map[int]int)

	var sleepyGuard string
	sleepiestMinute := make(map[string]int)

	lineRegexp := regexp.MustCompile(":|] | begins shift| #")

	var fallsAsleep int
	var err error
	var activeGuard string

	for _, entry := range schedule {

		minuteGuardAction := lineRegexp.Split(entry, -1)

		if strings.Contains(minuteGuardAction[2], "Guard") {
			activeGuard = minuteGuardAction[3]
		} else if strings.Contains(minuteGuardAction[2], "falls asleep") {
			fallsAsleep, err = strconv.Atoi(minuteGuardAction[1])
			if err != nil {
				panic("Not a valid minute value")
			}
		} else {
			wakesUp, err := strconv.Atoi(minuteGuardAction[1])
			if err != nil {
				panic("Not a valid minute value")
			}
			for i := fallsAsleep; i < wakesUp; i++ {
				guardsSleepSum[activeGuard]++
				if guardsSleepyMinutes[activeGuard] == nil {
					guardsSleepyMinutes[activeGuard] = make(map[int]int)
				}
				guardsSleepyMinutes[activeGuard][i]++
				if guardsSleepSum[sleepyGuard] < guardsSleepSum[activeGuard] {
					sleepyGuard = activeGuard
				}
				if guardsSleepyMinutes[activeGuard][sleepiestMinute[activeGuard]] < guardsSleepyMinutes[activeGuard][i] {
					sleepiestMinute[activeGuard] = i
				}
			}
		}
	}

	guardID, _ := strconv.Atoi(sleepyGuard)

	return guardID * sleepiestMinute[sleepyGuard]
}

func part2(input io.Reader) int {
	inputScanner := bufio.NewScanner(input)
	var schedule []string

	for inputScanner.Scan() {
		line := inputScanner.Text()
		schedule = append(schedule, line)
	}

	sort.Strings(schedule)

	guardsSleepyMinutes := make(map[string]map[int]int)

	sleepiestMinute := make(map[string]int)
	var sleepiestGuard string

	lineRegexp := regexp.MustCompile(":|] | begins shift| #")

	var fallsAsleep int
	var err error
	var activeGuard string

	for _, entry := range schedule {

		minuteGuardAction := lineRegexp.Split(entry, -1)

		if strings.Contains(minuteGuardAction[2], "Guard") {
			activeGuard = minuteGuardAction[3]
		} else if strings.Contains(minuteGuardAction[2], "falls asleep") {
			fallsAsleep, err = strconv.Atoi(minuteGuardAction[1])
			if err != nil {
				panic("Not a valid minute value")
			}
		} else {
			wakesUp, err := strconv.Atoi(minuteGuardAction[1])
			if err != nil {
				panic("Not a valid minute value")
			}
			for i := fallsAsleep; i < wakesUp; i++ {
				if guardsSleepyMinutes[activeGuard] == nil {
					guardsSleepyMinutes[activeGuard] = make(map[int]int)
				}
				guardsSleepyMinutes[activeGuard][i]++
				if guardsSleepyMinutes[activeGuard][sleepiestMinute[activeGuard]] < guardsSleepyMinutes[activeGuard][i] {
					sleepiestMinute[activeGuard] = i
					if guardsSleepyMinutes[sleepiestGuard][sleepiestMinute[sleepiestGuard]] < guardsSleepyMinutes[activeGuard][sleepiestMinute[activeGuard]] {
						sleepiestGuard = activeGuard
					}
				}
			}
		}
	}

	guardID, _ := strconv.Atoi(sleepiestGuard)

	return guardID * sleepiestMinute[sleepiestGuard]
}
