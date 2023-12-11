package y2023d10

import (
	_ "embed"
	"fmt"
	"math"
	"strings"
)

//go:embed test.txt
var test string

//go:embed input.txt
var input string

const Up = '|'
const Down = '|'
const Left = '-'
const Right = '-'
const DownRight = 'L'
const LeftUp = 'L'
const DownLeft = 'J'
const RightUp = 'J'
const UpLeft = '7'
const RightDown = '7'
const UpRight = 'F'
const LeftDown = 'F'
const Ground = '.'
const StartPoint = 'S'

const MovingUp = 1
const MovingLeft = 2
const MovingRight = 3
const MovingDown = 4

func getNextCell(ci, cj int, table [][]rune, direction int) (int, int, int, bool) {

	if ci < 0 || ci >= len(table) || cj < 0 || cj >= len(table[0]) {
		return ci, cj, 0, false
	}

	currentCell := table[ci][cj]

	if currentCell == Ground {
		return ci, ci, 0, false
	}

	switch direction {
	case MovingDown:
		switch currentCell {
		case Down:
			return ci + 1, cj, MovingDown, true
		case DownLeft:
			return ci, cj - 1, MovingLeft, true
		case DownRight:
			return ci, cj + 1, MovingRight, true
		}
		return ci, cj, 0, false
	case MovingUp:
		switch currentCell {
		case Up:
			return ci - 1, cj, MovingUp, true
		case UpLeft:
			return ci, cj - 1, MovingLeft, true
		case UpRight:
			return ci, cj + 1, MovingRight, true
		}
		return ci, cj, 0, false
	case MovingLeft:
		switch currentCell {
		case Left:
			return ci, cj - 1, MovingLeft, true
		case LeftDown:
			return ci + 1, cj, MovingDown, true
		case LeftUp:
			return ci - 1, cj, MovingUp, true
		}
		return ci, cj, 0, false
	case MovingRight:
		switch currentCell {
		case Right:
			return ci, cj + 1, MovingRight, true
		case RightDown:
			return ci + 1, cj, MovingDown, true
		case RightUp:
			return ci - 1, cj, MovingUp, true
		}
		return ci, cj, 0, false
	}

	return ci, ci, 0, false
}

func getStartingNode(table [][]rune) (int, int) {

	for i, row := range table {
		for j, col := range row {
			if col == StartPoint {
				return i, j
			}
		}
	}

	panic("Starting node not found in table")

}

func loop(starti, startj int, direction int, table [][]rune) (int, [][]int, bool) {

	ci, cj := starti, startj
	distances := make([][]int, len(table))

	// I know I am starting from the next step, not the Starting Point
	steps := 1

	for i, row := range table {
		distances[i] = make([]int, len(row))
	}

	ok := false
	for {
		steps++
		ci, cj, direction, ok = getNextCell(ci, cj, table, direction)

		if !ok {
			return 0, distances, false
		}

		if table[ci][cj] == StartPoint {
			return 0, distances, true
		}

		distances[ci][cj] = steps

	}
}

func getVertices(starti, startj int, direction int, table [][]rune) ([][2]int, bool) {

	ci, cj := starti, startj

	path := [][2]int{}

	// Checking the files S node is a corner in both cases
	// To make this generic we should make sure it does not behave
	// as | or -
	path = append(path, [2]int{starti, startj})

	ok := false
	for {
		ci, cj, direction, ok = getNextCell(ci, cj, table, direction)

		if !ok {
			return path, false
		}

		if table[ci][cj] == StartPoint {
			// We need the polygon to be closed, so we add the starting point again
			path = append(path, [2]int{ci, cj})

			return path, true
		}

		if table[ci][cj] != '.' && table[ci][cj] != '-' && table[ci][cj] != '|' {
			path = append(path, [2]int{ci, cj})
		}

	}
}

func Run1() {

	// Parse
	lines := strings.Split(input, "\n")
	table := [][]rune{}

	for _, line := range lines {
		table = append(table, []rune(line))
	}

	// Find starting node
	si, sj := getStartingNode(table)

	fmt.Println("start:", si, sj)

	maxUp, distancesUp, okUp := loop(si-1, sj, MovingUp, table)
	//fmt.Println("UP", maxUp, okUp, distancesUp)
	fmt.Println("UP", maxUp, okUp)

	maxDown, _, okDown := loop(si+1, sj, MovingDown, table)
	//fmt.Println("Down", maxDown, okDown, distancesDown)
	fmt.Println("Down", maxDown, okDown)

	maxRight, _, okRight := loop(si, sj+1, MovingRight, table)
	//fmt.Println("Right", maxRight, okRight, distancesRight)
	fmt.Println("Right", maxRight, okRight)

	maxLeft, distancesLeft, okLeft := loop(si, sj-1, MovingLeft, table)
	//fmt.Println("Left", maxLeft, okLeft, distancesLeft)
	fmt.Println("Left", maxLeft, okLeft)

	fmt.Println()
	// for _, row := range table {
	// 	for _, val := range row {
	// 		fmt.Printf("%c\t", val)
	// 	}
	// 	fmt.Println()
	// }

	fmt.Println()
	maxDistance := 0

	for i, row := range table {
		for j, _ := range row {
			minDistance := int(math.Min(float64(distancesLeft[i][j]), float64(distancesUp[i][j])))
			maxDistance = int(math.Max(float64(maxDistance), float64(minDistance)))

			// fmt.Printf("%d\t", minDistance)
		}
		//fmt.Println()
	}

	fmt.Println("max", maxDistance)

}

func calculateArea(path [][2]int) int {

	area := 0

	return area
}

func Run2() {
	lines := strings.Split(test, "\n")
	table := [][]rune{}

	for _, line := range lines {
		table = append(table, []rune(line))
	}

	// Find starting node
	si, sj := getStartingNode(table)

	fmt.Println("start:", si, sj)

	pathUp, okUp := getVertices(si-1, sj, MovingUp, table)
	//fmt.Println("UP", maxUp, okUp, distancesUp)
	fmt.Println("UP", pathUp, okUp)

	pathDown, okDown := getVertices(si+1, sj, MovingDown, table)
	//fmt.Println("Down", maxDown, okDown, distancesDown)
	fmt.Println("Down", pathDown, okDown)

	pathRight, okRight := getVertices(si, sj+1, MovingRight, table)
	//fmt.Println("Right", maxRight, okRight, distancesRight)
	fmt.Println("Right", pathRight, okRight)

	pathLeft, okLeft := getVertices(si, sj-1, MovingLeft, table)
	fmt.Println("PathLeft", pathLeft, okLeft)

	area := 0
	if okUp {
		area = calculateArea(pathUp)

	} else if okDown {
		area = calculateArea(pathDown)

	} else if okLeft {
		area = calculateArea(pathLeft)

	} else if okRight {
		area = calculateArea(pathRight)

	}

	fmt.Println("Area:", area)

}
