package main

import (
	"advent/utils/funct"
	"advent/utils/input"
	"advent/utils/search"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func StrToInt(str string) int {
	return funct.GetValue(strconv.Atoi(str))
}

type MonkeyData struct {
	currentItems        []int
	operation           func(int) int
	getNextMonkey       func(int) int
	totalInspectedItems int
}

func genValue(old int, arg string) int {

	if arg == "old" {
		return old
	}
	return StrToInt(arg)
}

func genOperation(args []string) func(int) int {
	arg1 := strings.Trim(args[0], " ")
	op := strings.Trim(args[1], " ")
	arg2 := strings.Trim(args[2], " ")

	return func(old int) int {
		var1 := genValue(old, arg1)
		var2 := genValue(old, arg2)
		switch op {
		case "+":
			return var1 + var2
		case "-":
			return var1 - var2
		case "*":
			return var1 * var2
		case "/":
			return var1 / var2
		}

		panic("Operation not supported!")
	}
}

func deleteFirst(array []int) []int {
	newArray := []int{}

	for _, item := range array[1:] {
		newArray = append(newArray, item)
	}

	return newArray
}

func contains(node [2]int, list [][2]int) bool {
	for _, item := range list {
		if node[0] == item[0] && node[1] == item[1] {
			return true
		}
	}

	return false
}

func getCandidates(currentNode [2]int, grid [][]rune) [][2]int {
	neighbours := search.Get4Neighbours(currentNode, grid)

	currentNodeValue := grid[currentNode[0]][currentNode[1]]

	candidates := [][2]int{}

	for _, neighbour := range neighbours {
		if grid[neighbour[0]][neighbour[1]]-currentNodeValue <= 1 {
			candidates = append(candidates, neighbour)
		}
	}

	return candidates
}

func main() {
	startValue := rune('S')
	endValue := rune('E')

	endNode := [2]int{}

	grid := input.Get2DMatrix(input.GetContents(2022, 12, "input.txt"), "\n", "", func(str string) rune { return rune(str[0]) })

	startNodes := [][2]int{}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == startValue {
				grid[i][j] = rune('a')
				//startNode = [2]int{i, j}
			} else if grid[i][j] == endValue {
				grid[i][j] = rune('z')
				endNode = [2]int{i, j}
			}

			if grid[i][j] == rune('a') {
				startNodes = append(startNodes, [2]int{i, j})
			}
		}
	}

	fmt.Println("startnodes", len(startNodes))

	minSteps := int(math.Inf(1))
	for idx, startNode := range startNodes {
		shortestPath := search.AStar(startNode, endNode, grid, getCandidates)

		if shortestPath != nil && minSteps > len(shortestPath) {
			minSteps = len(shortestPath)
		}

		fmt.Println("Calculated", idx)

	}

	fmt.Println(minSteps - 1)
}
