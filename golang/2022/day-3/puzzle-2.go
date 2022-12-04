package main

import (
	"advent/utils/input"
	"fmt"
	"strings"
)

func getPriority(r rune) int {
	if r >= 'a' && r <= 'z' {
		value := int(r - 'a' + 1)
		return value
	}

	if r >= 'A' && r <= 'Z' {
		value := int(r - 'A' + 27)
		return value

	}

	return 0
}

var Empty struct{}

func calculatePriority(str []string) int {
	priority := 0

	alreadyTaken := map[rune]struct{}{}

	for _, r := range str[0] {
		index := strings.IndexRune(str[1], r)

		if _, ok := alreadyTaken[r]; !ok && index > -1 {
			alreadyTaken[r] = Empty
			index = strings.IndexRune(str[2], r)

			if index > -1 {
				priority += getPriority(r)
			}
		}

	}

	return priority
}

func main() {
	data := input.GetLines(2022, 3, "input.txt")
	priority := 0
	for i := 0; i < len(data)/3; i++ {
		lines := data[i*3 : i*3+3]
		priority += calculatePriority(lines)
	}

	fmt.Println(priority)

}
