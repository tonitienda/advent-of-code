package main

import (
	"advent/utils/funct"
	"advent/utils/input"
	"fmt"
	"strconv"
	"strings"
)

// Shapes start in the column #2

func BinaryToInt(str string) int {
	return int(funct.GetValue(strconv.ParseInt(str, 2, 32)))
}

type Shape struct {
	structure []int
	width     int
	height    int
	top       int
	left      int
}

var HLine = Shape{
	structure: []int{BinaryToInt("0011110")},
	width:     4,
	height:    1,
}

// 0001000 => 8
// 0011100 => 24
// 0001000 =>
var Plus = Shape{
	structure: []int{
		BinaryToInt("0001000"),
		BinaryToInt("0011100"),
		BinaryToInt("0001000"),
	},
	width:  3,
	height: 3,
}

var LShape = Shape{
	structure: []int{
		BinaryToInt("0000100"),
		BinaryToInt("0000100"),
		BinaryToInt("0011100"),
	},
	width:  3,
	height: 3,
}

var VLine = Shape{
	structure: []int{
		BinaryToInt("0010000"),
		BinaryToInt("0010000"),
		BinaryToInt("0010000"),
		BinaryToInt("0010000"),
	},
	width:  1,
	height: 4,
}

var Square = Shape{
	structure: []int{
		BinaryToInt("0011000"),
		BinaryToInt("0011000"),
	},
	width:  2,
	height: 2,
}

var Shapes = []Shape{HLine, Plus, LShape, VLine, Square}

func printBoard(board []int) {
	for i := len(board) - 1; i >= 0; i-- {
		binary := fmt.Sprintf("%07s", strconv.FormatInt(int64(board[i]), 2))
		fmt.Println(strings.ReplaceAll(strings.ReplaceAll(binary, "0", "."), "1", "#"))
	}

}

// TODO - For now we ignore the pieces moving
func findCollission(shape []int, board []int) int {
	for boardRow := len(board) - 1; boardRow >= len(shape)-1; boardRow-- {
		for i := 0; i < len(shape); i++ {
			if board[boardRow-i]&shape[i] > 0 {
				// Collission, so we return the previous row
				return boardRow + 1
			}
		}
	}
	return len(shape) - 1
}

func main() {

	execType := "test"
	data := strings.Split(input.GetContents(2022, 17, execType+".txt"), "")

	fmt.Println(data)

	board := []int{}
	shapeCount := 0
	lastRowWithFallenRocks := 0

	for i := 0; i < 2022; i++ {
		// Falling means going from the end of the board to the beginning

		currentShape := Shapes[shapeCount%5]
		freeRows := len(board) - lastRowWithFallenRocks
		requiredFreeRows := 3 + len(currentShape.structure)

		fmt.Println("requiredFreeRows", requiredFreeRows, "freeRows", freeRows)
		for i := 0; i < requiredFreeRows-freeRows; i++ {
			board = append(board, 0)
		}

		//printBoard(board)

		row := findCollission(currentShape.structure, board)
		lastRowWithFallenRocks = row + len(currentShape.structure) - 1
		fmt.Println("Collission at:", row)
		for i := 0; i < len(currentShape.structure); i++ {
			board[row-i] |= currentShape.structure[i]
		}

		// printBoard(board)
		// fmt.Println()
		shapeCount++

	}

	printBoard(board)
	fmt.Println(shapeCount, "shapes")
	fmt.Println(lastRowWithFallenRocks, "lastRowWithFallenRocks")

}
