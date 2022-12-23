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
	for i := 0; i < len(shape); i++ {
		if board[row-i]&shape[i] > 0 {
			// Collission, so we return the previous row
			fmt.Println("Collission at row", row)
			return true
		}
	}
	return false
}

func move(board, shape []int, movement string, row int) ([]int, bool) {

	if movement == ">" {
		fmt.Println("Jet of gas pushes rock right:")
		return moveRight(shape)
	} else if movement == "<" {
		fmt.Println("Jet of gas pushes rock left:")
		return moveLeft(shape)
	}

	panic(movement + " not supported")
}

// TODO - For now we ignore the pieces moving
func placeShape(shape []int, board []int, movements []string, movementIdx int) (int, []int, int) {
	boardHeight := len(board)
	shapeHeight := len(shape)
	fmt.Println("The rock begins falling from", boardHeight-1)

	for row := boardHeight - 1; row >= shapeHeight; row-- {

		fmt.Println("Move down", row)
		printBoard(addShape(board, shape, row))
		fmt.Println()

		// If the next position collides, we try to move and see if we can go on
		// if not, return
		nextPositionCollides := shapeCollides(board, shape, row-1)

		if nextPositionCollides {
			movedShape, couldMove := move(board, shape, movements[movementIdx], row)

			if !couldMove {
				fmt.Println("Collided at ", row, "returning", row)
				return row, shape, movementIdx
			}
			nextNextPositionCollides := shapeCollides(board, movedShape, row-1)

			if nextNextPositionCollides {
				fmt.Println("Collided at ", row, "returning", row)
				movementIdx = (movementIdx + 1) % len(movements)

				return row, movedShape, movementIdx
			}

			shape = movedShape
			movementIdx = (movementIdx + 1) % len(movements)
		}

		// Get next movement
		movement := movements[movementIdx]

		// Try to move piece horizontally
		movedShape, couldMove := move(board, shape, movement, row)

		// If the shape could move (did not go out of bounds)
		// and the new location does not collide with other shapes
		// The shape is moved
		newPositionCollides := shapeCollides(board, movedShape, row)
		if couldMove && !newPositionCollides {
			shape = movedShape
		}

		movementIdx = (movementIdx + 1) % len(movements)

		// // If the shape is going to collide in the next step down
		// if shapeCollides(board, shape, row-1) {
		// 	newShape, canMoveAgain := move(board, shape, movements[movementIdx], row)

		// 	if canMoveAgain {
		// 		fmt.Println("Shape collides but can move one more time")
		// 		shape = newShape
		// 		movementIdx = (movementIdx + 1) % len(movements)

		// 		continue
		// 	}

		// 	fmt.Println("Shape collides and cannot move again")
		// 	return row, shape, movementIdx

		// }

	}

	bottomRow := len(shape) - 1
	newShape, _ := move(board, shape, movements[movementIdx], bottomRow)
	fmt.Println("No Collission. Returning ", len(shape)-1)

	return bottomRow, newShape, movementIdx + 1
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
	shapeCount := 0
	lastRowWithFallenRocks := 0
	movementIdx := 0

	for shapeIdx := 0; shapeIdx < 2022; shapeIdx++ {
		// Falling means going from the end of the board to the beginning

		currentShape := Shapes[shapeIdx%5]
		freeRows := len(board) - lastRowWithFallenRocks
		requiredFreeRows := 3 + len(currentShape.structure)

		fmt.Println("lastRowWithFallenRocks", lastRowWithFallenRocks, "Freerows:", freeRows, "RequiredFreeRows:", requiredFreeRows)

		for i := 0; i < requiredFreeRows-freeRows; i++ {
			board = append(board, 0)
		}

		row, shape, newMovementIdx := placeShape(currentShape.structure, board, movements, movementIdx)
		movementIdx = newMovementIdx
		lastRowWithFallenRocks = row + len(shape)
		// Update the board with the shape in rest
		board = addShape(board, shape, row)

		printBoard(board)
		fmt.Println()
		if shapeIdx >= 2 {
			break
		}

	}

	printBoard(board)
	fmt.Println(shapeCount, "shapes")
	fmt.Println(lastRowWithFallenRocks, "lastRowWithFallenRocks")

}
