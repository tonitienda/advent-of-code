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

func contains(x1, x2, y1, y2 int) bool {
	return x1 <= y1 && x2 >= y2
}

func overlaps(coords []int) bool {
	return contains(coords[0], coords[1], coords[2], coords[3]) ||
		contains(coords[2], coords[3], coords[0], coords[1])
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
