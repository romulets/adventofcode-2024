package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	raw, err := os.ReadFile("./input.txt")
	if err != nil {
		fmt.Printf("Error reading file: %v", err)
		return
	}

	total := sumMuls(string(raw))

	fmt.Printf("Total %d\n", total)
}

var regex = regexp.MustCompile(`(mul\([0-9]{1,3},[0-9]{1,3}\))|(do\(\))|(don't\(\))`)

func sumMulsWithRegex(tokens string) int {
	allInputs := regex.FindAllString(tokens, -1)

	circuitOn := true
	total := 0
	for _, input := range allInputs {
		switch input {
		case "don't()":
			circuitOn = false
			continue
		case "do()":
			circuitOn = true
			continue
		}

		if !circuitOn {
			continue
		}

		dirtyNums := strings.Split(input, ",")
		num1, _ := strconv.Atoi(dirtyNums[0][4:])
		num2, _ := strconv.Atoi(dirtyNums[1][:len(dirtyNums[1])-1])
		total += num1 * num2

	}

	return total
}

func sumMuls(tokens string) int {
	expr := make([]rune, 12)
	ptr := 0
	total := 0

	circuitOn := true
	for _, token := range tokens {
		expr[ptr] = token
		currentExpr := expr[:ptr+1]

		flip, flipStat := extractSwitches(currentExpr)
		if flipStat == building {
			ptr++
			continue
		} else if flipStat == finished {
			ptr = 0
			circuitOn = convertFlip(flip, circuitOn)
			continue
		}

		if !circuitOn {
			ptr = 0
			continue
		}

		num1, num2, stat := extractNums(currentExpr)

		switch stat {
		case building:
			ptr++
		case finished:
			ptr = 0
			total += (num1 * num2)
		default:
			ptr = 0
		}
	}

	return total
}

func convertFlip(flip flipSwitch, circuitOn bool) bool {
	switch flip {
	case flipOn:
		return true
	case flipOff:
		return false
	}
	return circuitOn
}

type strStat int

const (
	finished strStat = 3
	building strStat = 2
	invalid  strStat = 1
)

type flipSwitch int

const (
	flipOn   flipSwitch = 3
	flipOff  flipSwitch = 2
	flipKeep flipSwitch = 1
)

var (
	flipOnPattern  = []rune{'d', 'o', '(', ')'}
	flipOffPattern = []rune{'d', 'o', 'n', '\'', 't', '(', ')'}
)

func extractSwitches(tokens []rune) (flipSwitch, strStat) {
	if on, onStat := extractSwitch(tokens, flipOnPattern, flipOn); onStat != invalid {
		return on, onStat
	}
	return extractSwitch(tokens, flipOffPattern, flipOff)
}

func extractSwitch(tokens []rune, pattern []rune, flip flipSwitch) (flipSwitch, strStat) {
	if len(tokens) > len(pattern) {
		return flipKeep, invalid
	}

	for pos, token := range tokens {
		if pattern[pos] != token {
			return flipKeep, invalid
		}
	}

	if len(tokens) == len(pattern) {
		return flip, finished
	}

	return flipKeep, building
}

func extractNums(tokens []rune) (int, int, strStat) {
	l := len(tokens)
	switch {
	case l == 0:
		return -1, -1, building
	case l == 1:
		return checkExactPos(0, 'm', tokens)
	case l == 2:
		return checkExactPos(1, 'u', tokens)
	case l == 3:
		return checkExactPos(2, 'l', tokens)
	case l == 4:
		return checkExactPos(3, '(', tokens)
	case l >= 5 && l <= 12:
		return checkNumbersEndExpr(tokens[4:])
	}

	return -1, -1, invalid
}

func checkExactPos(pos int, expectedRune rune, expr []rune) (int, int, strStat) {
	if expr[pos] == expectedRune {
		return -1, -1, building
	} else {
		return -1, -1, invalid
	}
}

func checkNumbersEndExpr(tokens []rune) (int, int, strStat) {
	strNum1 := strings.Builder{}
	strNum2 := strings.Builder{}
	foundSep := false
	foundEnd := false

	for pos, token := range tokens {
		if pos == 0 {
			if isNum(token) {
				strNum1.WriteRune(token)
				continue
			}

			return -1, -1, invalid
		}

		if !foundSep && pos > 0 && pos < 3 {
			if isNum(token) {
				strNum1.WriteRune(token)
				continue
			}

			if token == ',' {
				foundSep = true
				continue
			}

			return -1, -1, invalid
		}

		// edge cut out
		if !foundSep && pos == 3 && token == ',' {
			foundSep = true
			continue
		}

		if foundSep {
			if isNum(token) {
				strNum2.WriteRune(token)
				continue
			}

			if token == ')' {
				foundEnd = true
				break
			}

			return -1, -1, invalid
		}

		return -1, -1, invalid
	}

	if strNum1.Len() > 3 || strNum2.Len() > 3 {
		return -1, -1, invalid
	}

	if foundEnd {
		num1, _ := strconv.Atoi(strNum1.String())
		num2, _ := strconv.Atoi(strNum2.String())
		return num1, num2, finished
	}

	return -1, -1, building
}

func isNum(n rune) bool {
	return n >= '0' && n <= '9'
}
