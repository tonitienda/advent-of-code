package main

import (
	"advent/utils/funct"
	"advent/utils/input"
	"fmt"
	"strconv"
)

func StrToInt(str string) int {
	return funct.GetValue(strconv.Atoi(str))
}

func isVisibleFromTop(trees [][]int, i, j int) bool {
	height := trees[i][j]

	for k := 0; k < i; k++ {
		if trees[k][j] >= height {

			return false
		}
	}

	return true
}

func isVisibleFromBottom(trees [][]int, i, j int) bool {
	height := trees[i][j]

	for k := i + 1; k < len(trees); k++ {
		if trees[k][j] >= height {
			return false
		}
	}

	return true
}

func isVisibleFromRight(trees [][]int, i, j int) bool {
	height := trees[i][j]

	for k := j + 1; k < len(trees[i]); k++ {
		if trees[i][k] >= height {
			return false
		}
	}

	return true
}

func isVisibleFromLeft(trees [][]int, i, j int) bool {
	height := trees[i][j]

	for k := 0; k < j; k++ {
		if trees[i][k] >= height {
			return false
		}
	}

	return true
}

func isVisible(trees [][]int, i, j int) bool {
	return isVisibleFromRight(trees, i, j) ||
		isVisibleFromLeft(trees, i, j) ||
		isVisibleFromBottom(trees, i, j) ||
		isVisibleFromTop(trees, i, j)
}

func main() {
	trees := input.Get2DMatrix(input.GetContents(2022, 8, "input.txt"), "\n", "", func(str string) int { return StrToInt(str) })

	visibleTrees := 0
	for i := 0; i < len(trees); i++ {
		for j := 0; j < len(trees[i]); j++ {
			if isVisible(trees, i, j) {
				visibleTrees++
			}
		}
	}
	fmt.Println(visibleTrees)
}
