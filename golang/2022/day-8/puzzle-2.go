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

func getScenicScoreTop(trees [][]int, i, j int) int64 {
	height := trees[i][j]
	var score int64 = 0

	for k := i - 1; k >= 0; k-- {
		score++
		if trees[k][j] >= height {
			return score
		}
	}

	return score
}

func getScenicScoreBottom(trees [][]int, i, j int) int64 {
	height := trees[i][j]

	var score int64 = 0
	for k := i + 1; k < len(trees); k++ {
		score++
		if trees[k][j] >= height {
			return score
		}
	}

	return score
}

func getScenicScoreRight(trees [][]int, i, j int) int64 {
	height := trees[i][j]
	var score int64 = 0

	for k := j + 1; k < len(trees[i]); k++ {
		score++
		if trees[i][k] >= height {
			return score
		}
	}

	return score
}

func getScenicScoreLeft(trees [][]int, i, j int) int64 {
	height := trees[i][j]
	var score int64 = 0

	for k := j - 1; k >= 0; k-- {
		score++
		if trees[i][k] >= height {
			return score
		}
	}

	return score
}

func getScenicScore(trees [][]int, i, j int) int64 {
	return getScenicScoreRight(trees, i, j) *
		getScenicScoreLeft(trees, i, j) *
		getScenicScoreBottom(trees, i, j) *
		getScenicScoreTop(trees, i, j)
}

func main() {
	trees := input.Get2DMatrix(input.GetContents(2022, 8, "input.txt"), "\n", "", func(str string) int { return StrToInt(str) })

	var maxScore int64 = 0
	for i := 0; i < len(trees); i++ {
		for j := 0; j < len(trees[i]); j++ {
			treeScore := getScenicScore(trees, i, j)
			if maxScore < treeScore {
				maxScore = treeScore
			}
		}
	}
	fmt.Println(maxScore)
}
