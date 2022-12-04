package main

import (
	"advent/utils/array"
	"advent/utils/funct"
	"advent/utils/input"
	"fmt"
	"strconv"
	"strings"
)

func StrToInt(str string) int {
	return funct.GetValue(strconv.Atoi(str))
}

func SplitBy(token string) func(string) []string {
	return func(str string) []string {
		return strings.Split(str, token)
	}
}

func getComponents(str string) []int {
	return array.Map(array.FlatMap(strings.Split(str, ","), SplitBy("-")), StrToInt)
}

func overlaps(coords []int) bool {
	return coords[0] <= coords[3] && coords[1] >= coords[2]
}

func countOverlaps(points [][]int) int {
	counter := 0

	for _, row := range points {
		if overlaps(row) {
			counter++
		}
	}

	return counter
}

func main() {
	data := array.Map(input.GetLines(2022, 4, "input.txt"), getComponents)

	fmt.Println((countOverlaps(data)))

}
