package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	raw, err := os.ReadFile("./input.txt")
	if err != nil {
		fmt.Printf("Error reading file: %v", err)
		return
	}

	lines := strings.Split(string(raw), "\n")
	reports := make([][]int, 0, len(lines))

	// convert
	for _, line := range lines {
		levels := strings.Split(line, " ")
		iLevels := make([]int, 0, len(levels))

		for _, level := range levels {
			iLevel, err := strconv.Atoi(level)
			if err != nil {
				fmt.Printf("Could convert int %s: %v", level, err)
				return
			}
			iLevels = append(iLevels, iLevel)
		}

		reports = append(reports, iLevels)
	}

	safeReps := 0
	safeExclOneReps := 0
	for _, levels := range reports {
		if isSafe(levels) {
			safeReps++
			safeExclOneReps++
		} else if isSafeExcludingOne(levels) {
			safeExclOneReps++
		}
	}

	fmt.Printf("Safe reports %d\n", safeReps)
	fmt.Printf("Safe reports excluding one %d", safeExclOneReps)
}

func isSafe(levels []int) bool {
	lastLevel := -1
	direction := ""

	for _, level := range levels {
		if lastLevel == -1 {
			lastLevel = level
			continue
		}

		if direction == "" {
			if lastLevel > level {
				direction = "desc"
			} else {
				direction = "asc"
			}
		}

		diff := lastLevel - level

		if direction == "asc" {
			if diff > -1 || diff < -3 { // the difference must be negative because of ascending direction
				// fmt.Printf("%s\n%d - %d = %d (dir=%s)\n\n", line, lastLevel, level, diff, direction)
				return false
			}
		}

		if direction == "desc" {
			if diff < 1 || diff > 3 { // the difference must be positive because of descending direction
				// fmt.Printf("%s\n%d - %d = %d (dir=%s)\n\n", line, lastLevel, level, diff, direction)
				return false
			}
		}

		lastLevel = level
	}

	return true
}

func isSafeExcludingOne(levels []int) bool {
	for i := 0; i < len(levels); i++ {
		levelsExclOne := make([]int, 0, len(levels)-1)

		for idx, level := range levels {
			if idx == i {
				continue
			}

			levelsExclOne = append(levelsExclOne, level)
		}

		if isSafe(levelsExclOne) {
			return true
		}
	}

	return false
}
