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
		printRow(board[i])
	}

}

func printRow(row int) {
	binary := fmt.Sprintf("%07s", strconv.FormatInt(int64(row), 2))
	fmt.Println(strings.ReplaceAll(strings.ReplaceAll(binary, "0", "."), "1", "#"))
}

// 7 bits
var maxNumber = 127

func moveRight(shape []int) []int {
	newShape := []int{}
	for i := 0; i < len(shape); i++ {
		// If the number is even, it means there is a 1
		// in position 0 so we cannot divide further
		if shape[i]%2 == 1 {
			return shape
		}
		n := shape[i] >> 1

		newShape = append(newShape, n)
	}

	return newShape
}

func moveLeft(shape []int) []int {
	newShape := []int{}
	for i := 0; i < len(shape); i++ {
		n := shape[i] << 1

		// If the shape goes beyond limits, return the original shape
		if n > maxNumber {
			return shape
		}

		newShape = append(newShape, n)
	}

	return newShape
}

func shapeCollides(board, shape []int, row int) bool {
	for i := 0; i < len(shape); i++ {
		if board[row-i]&shape[i] > 0 {
			// Collission, so we return the previous row
			fmt.Println("Collission at row", row)
			return true
		}
	}
	return false
}

func move(board, shape []int, movement string, row int) []int {
	newShape := shape

	if movement == ">" {
		fmt.Println("Jet of gas pushes rock right:")
		newShape = moveRight(shape)
	} else if movement == "<" {
		fmt.Println("Jet of gas pushes rock left:")

		newShape = moveLeft(shape)
	}

	collides := shapeCollides(board, newShape, row)

	if collides {
		return shape
	}

	return newShape

}

// TODO - For now we ignore the pieces moving
func placeShape(shape []int, board []int, movements []string, movementIdx int) (int, []int, int) {
	fmt.Println("The rock begins falling:", movementIdx)
	printBoard(shape)

	for boardRow := len(board) - 1; boardRow >= len(shape)-1; boardRow-- {
		fmt.Println("Row:", boardRow)
		// Move shape based on movements
		movement := movements[movementIdx]

		shape = move(board, shape, movement, boardRow)
		movementIdx = (movementIdx + 1) % len(movements)

		if shapeCollides(board, shape, boardRow) {
			// Try to move again and see if it does not collide
			shape2 := move(board, shape, movements[movementIdx], boardRow-1)
			collidesAgain := shapeCollides(board, shape2, boardRow)

			if collidesAgain {
				return boardRow + 1, shape, movementIdx
			} else {
				movementIdx = (movementIdx + 1) % len(movements)
				shape = shape2
			}
		}

	}
	fmt.Println("No Collission. Returning ", len(shape)-1)

	return len(shape) - 1, shape, movementIdx
}

func main() {

	execType := "test"
	movements := strings.Split(input.GetContents(2022, 17, execType+".txt"), "")

	fmt.Println(movements)

	board := []int{}
	shapeCount := 0
	lastRowWithFallenRocks := 0
	movementIdx := 0

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

		row, shape, newMovementIdx := placeShape(currentShape.structure, board, movements, movementIdx)
		movementIdx = newMovementIdx
		lastRowWithFallenRocks = row + len(currentShape.structure) - 1
		fmt.Println("Collission at:", row)
		for i := 0; i < len(shape); i++ {
			board[row-i] |= shape[i]
		}

		printBoard(board)
		fmt.Println()
		shapeCount++
		if shapeCount > 2 {
			break
		}

	}

	printBoard(board)
	fmt.Println(shapeCount, "shapes")
	fmt.Println(lastRowWithFallenRocks, "lastRowWithFallenRocks")

}
