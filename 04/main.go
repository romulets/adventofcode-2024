package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	raw, err := os.ReadFile("./input.txt")
	if err != nil {
		fmt.Printf("Error reading file: %v", err)
		return
	}

	fmt.Println(findXmas(string(raw)))
	fmt.Println(findXmasVector(string(raw)))
}

const windowSize = 4

func findXmas(tokens string) int {
	total := 0
	lines := strings.Split(tokens, "\n")

	total += findHorizontal(lines)
	total += findVertical(lines)
	total += findDiagonalRight(lines)
	total += findDiagonalLeft(lines)

	return total
}

func findHorizontal(lines []string) int {
	total := 0
	for i := 0; i < len(lines); i++ {
		start := 0
		end := windowSize

		for end <= len(lines[i]) {

			if isXmas(lines[i][start:end]) {
				total++
			}
			start++
			end++
		}
	}
	return total
}

func findVertical(lines []string) int {
	total := 0

	for j := 0; j < len(lines[0]); j++ {
		start := 0
		tokens := make([]rune, 4)

		for start+windowSize <= len(lines) {
			for i := 0; i < windowSize; i++ {
				tokens[i] = rune(lines[i+start][j])
			}

			if isXmas(string(tokens)) {
				total++
			}
			start++
		}
	}
	return total
}

func findDiagonalRight(lines []string) int {
	total := 0
	tokens := make([]rune, windowSize)

	for j := 0; j < len(lines[0]); j++ {
		for i := 0; i < len(lines); i++ {

			hasValid := true
			for n := 0; n < windowSize; n++ {
				if i+n >= len(lines) || j+n >= len(lines[0]) {
					hasValid = false
					break
				}
				tokens[n] = rune(lines[i+n][j+n])
			}

			if !hasValid {
				break
			}

			if isXmas(string(tokens)) {
				total++
			}
		}
	}

	return total
}

func findDiagonalLeft(lines []string) int {
	total := 0
	tokens := make([]rune, windowSize)

	for j := len(lines[0]) - 1; j >= 0; j-- {
		for i := 0; i < len(lines); i++ {

			hasValid := true
			for n := 0; n < windowSize; n++ {
				if i+n >= len(lines) || j-n < 0 {
					hasValid = false
					break
				}
				tokens[n] = rune(lines[i+n][j-n])
			}

			if !hasValid {
				break
			}

			if isXmas(string(tokens)) {
				total++
			}
		}
	}

	return total
}

func isXmas(tokens string) bool {
	return tokens == "XMAS" || tokens == "SAMX"
}

func findXmasVector(tokens string) int {
	total := 0
	lines := strings.Split(tokens, "\n")

	xmas1 := make([]rune, 3)
	xmas2 := make([]rune, 3)

	vect := make([][]rune, 3)
	for n := 0; n < 3; n++ {
		vect[n] = make([]rune, 3)
	}

	for i := 0; i <= len(lines); i++ {
		for j := 0; j <= len(lines[0]); j++ {

			valid := true
			for n1 := 0; n1 < 3; n1++ {
				for n2 := 0; n2 < 3; n2++ {
					if n1+i >= len(lines) || n2+j >= len(lines[0]) {
						valid = false
						break
					}
					vect[n1][n2] = rune(lines[n1+i][n2+j])
				}
			}

			if !valid {
				break
			}

			xmas1[0] = vect[0][0]
			xmas1[1] = vect[1][1]
			xmas1[2] = vect[2][2]

			xmas2[0] = vect[0][2]
			xmas2[1] = vect[1][1]
			xmas2[2] = vect[2][0]

			if isMas(string(xmas1)) && isMas(string(xmas2)) {
				total++
			}

		}
	}

	return total
}

func isMas(tokens string) bool {
	return tokens == "MAS" || tokens == "SAM"
}
