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
		printRow(i, board[i])
	}

}

func binary(n int) string {
	return fmt.Sprintf("%07s", strconv.FormatInt(int64(n), 2))
}

func printRow(idx, row int) {

	fmt.Println(idx, "\t", strings.ReplaceAll(strings.ReplaceAll(binary(row), "0", "."), "1", "#"))
}

// 7 bits
var maxNumber = 127

func moveRight(shape []int) ([]int, bool) {
	newShape := []int{}
	for i := 0; i < len(shape); i++ {
		// If the number is even, it means there is a 1
		// in position 0 so we cannot divide further
		if shape[i]%2 == 1 {
			return shape, false
		}
		n := shape[i] >> 1

		newShape = append(newShape, n)
	}

	return newShape, true
}

func moveLeft(shape []int) ([]int, bool) {
	newShape := []int{}
	for i := 0; i < len(shape); i++ {
		n := shape[i] << 1

		// If the shape goes beyond limits, return the original shape
		if n > maxNumber {
			return shape, false
		}

		newShape = append(newShape, n)
	}

	return newShape, true
}

func shapeCollides(board, shape []int, row int) bool {
	for i := len(shape) - 1; i >= 0; i-- {

		if board[row-i]&shape[i] > 0 {
			return true
		}
	}
	return false
}

func move(board, shape []int, movement string, row int) ([]int, bool) {

	var newShape []int
	var couldMove bool

	if movement == ">" {
		//fmt.Println("Jet of gas pushes rock right:")
		newShape, couldMove = moveRight(shape)
	} else if movement == "<" {
		//fmt.Println("Jet of gas pushes rock left:")
		newShape, couldMove = moveLeft(shape)
	}

	if !couldMove {
		return shape, couldMove
	}

	collides := shapeCollides(board, newShape, row)

	if collides {
		return shape, false
	}

	return newShape, true
}

// TODO - For now we ignore the pieces moving
func placeShape(shape []int, board []int, movements []string, movementIdx int) (int, []int, int) {
	boardHeight := len(board)
	shapeHeight := len(shape)
	row := boardHeight - 1
	//fmt.Println("The rock begins falling from", row)

	for {

		// Move horizontally
		shape, _ = move(board, shape, movements[movementIdx], row)
		movementIdx = (movementIdx + 1) % len(movements)

		// printBoard(addShape(board, shape, row))
		// fmt.Println()

		// fmt.Println("Rock falls 1 unit:")
		row--

		if shapeCollides(board, shape, row) {
			return row + 1, shape, movementIdx
		}

		if row < shapeHeight {
			shape, _ := move(board, shape, movements[movementIdx], row)
			movementIdx = (movementIdx + 1) % len(movements)

			return row, shape, movementIdx
		}
	}
}

func addShape(board, shape []int, row int) []int {
	newBoard := []int{}

	for i := 0; i < len(board); i++ {
		newBoard = append(newBoard, board[i])
	}
	// Update the board with the shape in rest
	for i := 0; i < len(shape); i++ {
		newBoard[row-i] |= shape[i]
	}

	return newBoard
}

func main() {

	execType := "test"
	movements := strings.Split(input.GetContents(2022, 17, execType+".txt"), "")

	fmt.Println(movements)

	board := []int{}
	shapeIdx := 0
	lastRowWithFallenRocks := 0
	movementIdx := 0

	for shapeIdx = 0; shapeIdx < 1000000; shapeIdx++ {
		// Falling means going from the end of the board to the beginning
		currentShape := Shapes[shapeIdx%5]

		for i := 0; i < len(board); i++ {
			if board[i] == 0 {
				lastRowWithFallenRocks = i
				break
			}
		}
		requiredRows := 3 + len(currentShape.structure) + lastRowWithFallenRocks
		diff := requiredRows - len(board)

		// fmt.Println("len(board)", len(board), "lastRowWithFallenRocks", lastRowWithFallenRocks, "len(currentShape.structure)", len(currentShape.structure), "requiredRows:", requiredRows, "diff", diff)

		// printBoard(board)
		// fmt.Println()

		if diff > 0 {
			for i := 0; i < diff; i++ {
				board = append(board, 0)
			}
		} else {
			//fmt.Println("old board len", len(board))

			board = board[:len(board)+diff]
			//fmt.Println("new board len", len(board), "requiredLength", requiredRows)

		}

		if len(board) != requiredRows {
			panic("Rows should match")
		}

		// printBoard(addShape(board, currentShape.structure, len(board)-1))
		// fmt.Println()

		row, shape, newMovementIdx := placeShape(currentShape.structure, board, movements, movementIdx)
		movementIdx = newMovementIdx
		// lastRowWithFallenRocks = row
		// Update the board with the shape in rest
		board = addShape(board, shape, row)

		if board[row] == 7 {
			fmt.Println("Row is complete")
			return

		}
		// printBoard(board)
		// fmt.Println()
		// if shapeIdx >= 9 {
		// 	break
		// }

	}

	for i := 0; i < len(board); i++ {
		if board[i] == 0 {
			lastRowWithFallenRocks = i
			break
		}
	}

	// printBoard(board)
	// fmt.Println()
	//printBoard(board[len(board)-100:])
	//fmt.Println(len(board))
	fmt.Println(shapeIdx, "shapes")
	fmt.Println(lastRowWithFallenRocks, "lastRowWithFallenRocks")

}
