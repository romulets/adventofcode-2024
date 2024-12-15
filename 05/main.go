package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	raw, err := os.ReadFile("./input.txt")
	if err != nil {
		fmt.Printf("Error reading file: %v", err)
		return
	}

	fmt.Println(sumOfPages(string(raw)))
}

func sumOfPages(lines string) (int, int) {
	totalCorrect := 0
	totalIncorrect := 0
	ordering := make(map[int][]int, strings.Count(lines, "|"))

	orderingSwitch := true
	for _, line := range strings.Split(lines, "\n") {
		if line == "" {
			orderingSwitch = false
			continue
		}

		if orderingSwitch {
			extractOrdering(line, ordering)
			continue
		}

		pages, valid := extractValidLine(line, ordering)

		if valid {
			totalCorrect += pages[int(math.Floor(float64(len(pages))/float64(2)))]
		} else {
			totalIncorrect += sortAndFindMiddle(pages, ordering)
		}
	}

	return totalCorrect, totalIncorrect
}

func extractOrdering(line string, ordering map[int][]int) {
	tokens := strings.Split(line, "|")
	before, _ := strconv.Atoi(tokens[0])
	after, _ := strconv.Atoi(tokens[1])

	listAfter, ok := ordering[before]
	if !ok {
		listAfter = make([]int, 0, 20)
	}

	listAfter = append(listAfter, after)
	ordering[before] = listAfter
}

func extractValidLine(line string, ordering map[int][]int) ([]int, bool) {
	rawPages := strings.Split(line, ",")
	pages := make([]int, 0, len(rawPages))
	valid := true

	for _, page := range rawPages {
		pNum, _ := strconv.Atoi(page)

		for _, mustBeBefore := range ordering[pNum] {
			if slices.Contains(pages, mustBeBefore) {
				valid = false
			}
		}

		pages = append(pages, pNum)
	}

	return pages, valid
}

func sortAndFindMiddle(pages []int, ordering map[int][]int) int {
	slices.SortFunc(pages, func(n1 int, n2 int) int {
		if _, found := ordering[n1]; !found {
			return 1
		}

		if slices.Contains(ordering[n1], n2) {
			return -1
		} else {
			return 1
		}
	})

	return pages[int(math.Floor(float64(len(pages))/float64(2)))]
}
