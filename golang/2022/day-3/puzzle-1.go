package main

import (
	"advent/utils/array"
	"advent/utils/input"
	"fmt"
	"strings"
)

func getPriority(r rune) int {
	if r >= 'a' && r <= 'z' {
		return int(r - 'a' + 1)
	}

	if r >= 'A' && r <= 'Z' {
		return int(r - 'A' + 27)
	}

	return 0
}

var Empty struct{}

func calculatePriorities(str string) int {
	priority := 0
	part1 := str[0 : len(str)/2]
	part2 := str[len(str)/2:]

	alreadyTaken := map[rune]struct{}{}

	for _, r := range part1 {
		index := strings.IndexRune(part2, r)

		if _, ok := alreadyTaken[r]; !ok && index > -1 {
			alreadyTaken[r] = Empty
			priority += getPriority(r)
		}

	}

	return priority
}

func main() {
	data := input.GetLines(2022, 3, "input.txt")

	priorities := array.Map(data, calculatePriorities)
	fmt.Println(array.Sum(priorities))

}
