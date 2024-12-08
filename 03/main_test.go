package main

import (
	"fmt"
	"os"
	"testing"
)

// BenchmarkSumMuls-12          	    6468	    187908 ns/op	   51579 B/op	    4144 allocs/op
// BenchmarkSumMulsRegexp-12    	    1021	   1174053 ns/op	  118407 B/op	    1099 allocs/op

func BenchmarkSumMuls(b *testing.B) {
	raw, err := os.ReadFile("./input.txt")
	if err != nil {
		fmt.Printf("Error reading file: %v", err)
		return
	}

	for i := 0; i < b.N; i++ {
		sumMuls(string(raw))
	}
}

func BenchmarkSumMulsRegexp(b *testing.B) {
	raw, err := os.ReadFile("./input.txt")
	if err != nil {
		fmt.Printf("Error reading file: %v", err)
		return
	}

	for i := 0; i < b.N; i++ {
		sumMulsWithRegex(string(raw))
	}
}

func TestSumMuls(t *testing.T) {
	tcs := []struct {
		expr  string
		total int
	}{
		{
			expr:  "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))",
			total: 48,
		},
		{
			expr:  "mul(4*, mul(6,9!, ?(12,34)",
			total: 0,
		},
		{
			expr:  "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))",
			total: 161,
		},
		{
			expr:  "mul(44,46)",
			total: 2024,
		},
		{
			expr:  "mul ( 2 , 4 )",
			total: 0,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.expr, func(t *testing.T) {
			total := sumMuls(tc.expr)
			if total != tc.total {
				t.Errorf("Failed total, expected %d got %d for expr %s", tc.total, total, tc.expr)
			}
		})
	}
}

func TestIsValid(t *testing.T) {
	tcs := map[string]struct {
		tokens  []rune
		expStat strStat
		num1    int
		num2    int
	}{
		"invalid":      {tokens: []rune{'n'}, expStat: invalid, num1: -1, num2: -1},
		"empty":        {tokens: []rune{}, expStat: building, num1: -1, num2: -1},
		"m":            {tokens: []rune{'m'}, expStat: building, num1: -1, num2: -1},
		"mu":           {tokens: []rune{'m', 'u'}, expStat: building, num1: -1, num2: -1},
		"mul":          {tokens: []rune{'m', 'u', 'l'}, expStat: building, num1: -1, num2: -1},
		"mul(":         {tokens: []rune{'m', 'u', 'l', '('}, expStat: building, num1: -1, num2: -1},
		"mul(1":        {tokens: []rune{'m', 'u', 'l', '(', '1'}, expStat: building, num1: -1, num2: -1},
		"mul(12":       {tokens: []rune{'m', 'u', 'l', '(', '1', '2'}, expStat: building, num1: -1, num2: -1},
		"mul(123":      {tokens: []rune{'m', 'u', 'l', '(', '1', '2', '3'}, expStat: building, num1: -1, num2: -1},
		"mul(1,23":     {tokens: []rune{'m', 'u', 'l', '(', '1', ',', '2', '3'}, expStat: building, num1: -1, num2: -1},
		"mul(12,3":     {tokens: []rune{'m', 'u', 'l', '(', '1', '2', ',', '3'}, expStat: building, num1: -1, num2: -1},
		"mul(123,":     {tokens: []rune{'m', 'u', 'l', '(', '1', '2', '3', ','}, expStat: building, num1: -1, num2: -1},
		"mul(123,4":    {tokens: []rune{'m', 'u', 'l', '(', '1', '2', '3', ',', '4'}, expStat: building, num1: -1, num2: -1},
		"mul(123,45":   {tokens: []rune{'m', 'u', 'l', '(', '1', '2', '3', ',', '4', '5'}, expStat: building, num1: -1, num2: -1},
		"mul(123,456":  {tokens: []rune{'m', 'u', 'l', '(', '1', '2', '3', ',', '4', '5', '6'}, expStat: building, num1: -1, num2: -1},
		"mul(1,4)":     {tokens: []rune{'m', 'u', 'l', '(', '1', ',', '4', ')'}, expStat: finished, num1: 1, num2: 4},
		"mul(1,45)":    {tokens: []rune{'m', 'u', 'l', '(', '1', ',', '4', '5', ')'}, expStat: finished, num1: 1, num2: 45},
		"mul(1,456)":   {tokens: []rune{'m', 'u', 'l', '(', '1', ',', '4', '5', '6', ')'}, expStat: finished, num1: 1, num2: 456},
		"mul(12,4)":    {tokens: []rune{'m', 'u', 'l', '(', '1', '2', ',', '4', ')'}, expStat: finished, num1: 12, num2: 4},
		"mul(12,45)":   {tokens: []rune{'m', 'u', 'l', '(', '1', '2', ',', '4', '5', ')'}, expStat: finished, num1: 12, num2: 45},
		"mul(12,456)":  {tokens: []rune{'m', 'u', 'l', '(', '1', '2', ',', '4', '5', '6', ')'}, expStat: finished, num1: 12, num2: 456},
		"mul(123,4)":   {tokens: []rune{'m', 'u', 'l', '(', '1', '2', '3', ',', '4', ')'}, expStat: finished, num1: 123, num2: 4},
		"mul(123,45)":  {tokens: []rune{'m', 'u', 'l', '(', '1', '2', '3', ',', '4', '5', ')'}, expStat: finished, num1: 123, num2: 45},
		"mul(123,456)": {tokens: []rune{'m', 'u', 'l', '(', '1', '2', '3', ',', '4', '5', '6', ')'}, expStat: finished, num1: 123, num2: 456},
	}

	for name, tc := range tcs {
		t.Run(name, func(t *testing.T) {
			num1, num2, stat := extractNums(tc.tokens)
			if stat != tc.expStat {
				t.Errorf("Failed stat, expected %d got %d for tokens %s", tc.expStat, stat, string(tc.tokens))
			}

			if num1 != tc.num1 {
				t.Errorf("Failed num1, expected %d got %d for tokens %s", tc.num1, num1, string(tc.tokens))
			}

			if num2 != tc.num2 {
				t.Errorf("Failed num2, expected %d got %d for tokens %s", tc.num2, num2, string(tc.tokens))
			}
		})
	}
}
