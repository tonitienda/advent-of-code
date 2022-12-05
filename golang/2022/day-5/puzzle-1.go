package main

import (
	"advent/utils/array"
	"advent/utils/funct"
	"advent/utils/input"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func getStacks(str string) [][]string {

	stacks := make([][]string, 9)

	lines := strings.Split(str, "\n")
	for _, line := range lines[:len(lines)-1] {
		for j := 0; j <= len(line)/4; j++ {
			block := line[j*4 : int(math.Min(float64(j*4+4), float64(len(line)-1)))]

			block2 := cleanBlock(block)
			if block2 != "" {
				stacks[j] = append(stacks[j], block2)
			}
		}
	}

	return stacks
}

func cleanBlock(str string) string {
	return strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(str, "[", ""), "]", ""), " ", "")
}

func StrToInt(str string) int {
	return funct.GetValue(strconv.Atoi(str))
}

func toMovement(str string) []int {
	components := strings.Split(str, " ")

	return []int{StrToInt(components[1]), StrToInt(components[3]), StrToInt(components[5])}
}

func getMovements(str string) [][]int {
	lines := strings.Split(str, "\n")

	return array.Map(lines, toMovement)
}

// TODO - Copying the array. Pass by reference to improve performance
func applyMovement(stacks [][]string, movement []int) [][]string {
	q := movement[0]
	from := movement[1] - 1
	to := movement[2] - 1

	for i := 0; i < q; i++ {
		if len(stacks[from]) > 0 {
			take := stacks[from][0]
			stacks[to] = append([]string{take}, stacks[to]...)
			stacks[from] = stacks[from][1:]
		}
	}

	return stacks
}

func main() {
	data := strings.Split(input.GetContents(2022, 5, "input.txt"), "\n\n")

	stacks := getStacks(data[0])
	movs := getMovements(data[1])

	for _, mov := range movs {
		stacks = applyMovement(stacks, mov)

	}

	result := ""
	for _, s := range stacks {
		if len(s) > 0 {
			result += s[0]
		}
	}

	fmt.Println(result)
}
