package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

// https://adventofcode.com/2024/day/1
func main() {
	raw, err := os.ReadFile("./input.txt")
	if err != nil {
		fmt.Printf("Error reading file: %v", err)
		return
	}

	lines := strings.Split(string(raw), "\n")

	list1 := make([]int, 0, len(lines))
	list2 := make([]int, 0, len(lines))

	for pos, line := range lines {
		vals := strings.Split(line, "   ")

		val1, err := strconv.Atoi(vals[0])
		if err != nil {
			fmt.Printf("Error parsing val 1 '%s' (pos %d): %v", vals[0], pos, err)
			return
		}

		list1 = append(list1, val1)

		val2, err := strconv.Atoi(vals[1])
		if err != nil {
			fmt.Printf("Error parsing val 1 '%s' (pos %d): %v", vals[0], pos, err)
			return
		}

		list2 = append(list2, val2)
	}

	slices.Sort(list1)
	slices.Sort(list2)

	fmt.Printf("The total distance between lists is %d \n", distance(list1, list2))
	fmt.Printf("The similarity score between lists is %d \n", similarity(list1, list2))
}

func distance(list1, list2 []int) int {
	sum := 0
	for i := 0; i < len(list1); i++ {
		val := list1[i] - list2[i]
		if val < 0 {
			val *= -1
		}

		sum += val
	}

	return sum
}

func similarity(list1, list2 []int) int {
	sim := 0
	freq := make(map[int]int, len(list2))

	for _, val := range list2 {
		freq[val] = freq[val] + 1
	}

	for _, val := range list1 {
		sim += val * freq[val]
		freq[val] = 0 // clean up to avoid duplicate entries
	}

	return sim
}
