package main

import (
	"advent/utils/array"
	"advent/utils/input"
	"fmt"
	"strings"
)

const Rock_1 = "A"
const Paper_1 = "B"
const Scissors_1 = "C"

const Loose = "X"
const Draw = "Y"
const Win = "Z"

const Rock_Points = 1
const Paper_Points = 2
const Scissor_Points = 3

const Loose_Points = 0
const Draw_Points = 3
const Win_Points = 6

func main() {
	data := input.GetContents(2022, 2, "input.txt")

	points := map[string]map[string]int{
		Rock_1: {
			Draw:  Rock_Points + Draw_Points,
			Win:   Paper_Points + Win_Points,
			Loose: Scissor_Points + Loose_Points,
		},
		Paper_1: {
			Loose: Rock_Points + Loose_Points,
			Draw:  Paper_Points + Draw_Points,
			Win:   Scissor_Points + Win_Points,
		},
		Scissors_1: {
			Win:   Rock_Points + Win_Points,
			Loose: Paper_Points + Loose_Points,
			Draw:  Scissor_Points + Draw_Points,
		},
	}

	parsedData := array.Map(strings.Split(data, "\n"), func(s string) []string { return strings.Split(s, " ") })

	gamePoints := array.Map(parsedData, func(s []string) int { return points[s[0]][s[1]] })

	fmt.Println((array.Sum(gamePoints)))

}
