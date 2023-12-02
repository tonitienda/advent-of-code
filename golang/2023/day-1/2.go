package main

import (
	_ "embed"
	"fmt"
	"math"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	numbersmap := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
		"1":     1,
		"2":     2,
		"3":     3,
		"4":     4,
		"5":     5,
		"6":     6,
		"7":     7,
		"8":     8,
		"9":     9,
	}

	lines := strings.Split(input, "\n")

	total := 0

	for _, line := range lines {
		firstIndex := math.MaxInt
		lastIndex := -1
		firstNumber := 0
		lastNumber := 0
		for key, value := range numbersmap {
			fIndex := strings.Index(line, key)
			lIndex := strings.LastIndex(line, key)
			if fIndex > -1 && fIndex < firstIndex {
				firstNumber = value * 10
				firstIndex = fIndex
			}

			if lIndex > lastIndex {
				lastNumber = value
				lastIndex = lIndex
			}

		}

		fmt.Printf("%d %d\n", firstNumber, lastNumber)
		total += firstNumber + lastNumber
	}

	fmt.Println(total)
}
