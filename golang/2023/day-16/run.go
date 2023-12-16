package y2023d16

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

const Vertical = '|'
const Horizontal = '|'
const Slash = '/'
const Backslash = '\\'

const MovingUp = 1
const MovingLeft = 2
const MovingRight = 3
const MovingDown = 4

func getMatrix(text string) [][]rune {
	matrix := [][]rune{}

	for _, line := range strings.Split(text, "\n") {
		matrix = append(matrix, []rune(line))
	}

	return matrix
}

type thing struct {
	i         int
	j         int
	direction int
}

func tryMoveRight(position thing, lastRow, lastCol int) *thing {
	if position.j == lastCol {
		// We cannot move further right
		return nil
	}

	return &thing{
		i:         position.i,
		j:         position.j + 1,
		direction: MovingRight,
	}

}

func tryMoveLeft(position thing, lastRow, lastCol int) *thing {
	if position.j == 0 {
		// We cannot move further left
		return nil
	}

	return &thing{
		i:         position.i,
		j:         position.j - 1,
		direction: MovingLeft,
	}

}

func tryMoveUp(position thing, lastRow, lastCol int) *thing {
	if position.i == 0 {
		// We cannot move further up
		return nil
	}

	return &thing{
		i:         position.i - 1,
		j:         position.j,
		direction: MovingUp,
	}
}

func tryMoveDown(position thing, lastRow, lastCol int) *thing {
	if position.i == lastRow {
		// We cannot move further down
		return nil
	}

	return &thing{
		i:         position.i + 1,
		j:         position.j,
		direction: MovingDown,
	}
}

func getNextPositions(position thing, matrix [][]rune) []thing {

	currentCell := matrix[position.i][position.j]
	lastRow := len(matrix) - 1
	lastCol := len(matrix[position.i]) - 1

	if currentCell == '.' {
		switch position.direction {
		case MovingRight:
			nextMove := tryMoveRight(position, lastRow, lastCol)

			if nextMove != nil {
				return []thing{*nextMove}
			}
		case MovingLeft:
			nextMove := tryMoveLeft(position, lastRow, lastCol)

			if nextMove != nil {
				return []thing{*nextMove}
			}
		case MovingUp:
			nextMove := tryMoveUp(position, lastRow, lastCol)

			if nextMove != nil {
				return []thing{*nextMove}
			}

		case MovingDown:
			nextMove := tryMoveDown(position, lastRow, lastCol)

			if nextMove != nil {
				return []thing{*nextMove}
			}
		}
	}

	if currentCell == '/' {
		switch position.direction {
		case MovingRight:
			nextMove := tryMoveUp(position, lastRow, lastCol)

			if nextMove != nil {
				return []thing{*nextMove}
			}
		case MovingLeft:
			nextMove := tryMoveDown(position, lastRow, lastCol)

			if nextMove != nil {
				return []thing{*nextMove}
			}
		case MovingUp:
			nextMove := tryMoveRight(position, lastRow, lastCol)

			if nextMove != nil {
				return []thing{*nextMove}
			}

		case MovingDown:
			nextMove := tryMoveLeft(position, lastRow, lastCol)

			if nextMove != nil {
				return []thing{*nextMove}
			}
		}
	}

	if currentCell == '\\' {
		switch position.direction {
		case MovingRight:
			nextMove := tryMoveDown(position, lastRow, lastCol)

			if nextMove != nil {
				return []thing{*nextMove}
			}
		case MovingLeft:
			nextMove := tryMoveUp(position, lastRow, lastCol)

			if nextMove != nil {
				return []thing{*nextMove}
			}
		case MovingUp:
			nextMove := tryMoveLeft(position, lastRow, lastCol)

			if nextMove != nil {
				return []thing{*nextMove}
			}

		case MovingDown:
			nextMove := tryMoveRight(position, lastRow, lastCol)

			if nextMove != nil {
				return []thing{*nextMove}
			}
		}
	}

	if currentCell == '|' {
		switch position.direction {
		case MovingRight:
			nextMoves := []thing{}

			n1 := tryMoveUp(position, lastRow, lastCol)

			if n1 != nil {
				nextMoves = append(nextMoves, *n1)
			}

			n2 := tryMoveDown(position, lastRow, lastCol)

			if n2 != nil {
				nextMoves = append(nextMoves, *n2)
			}
			return nextMoves

		case MovingLeft:
			nextMoves := []thing{}

			n1 := tryMoveUp(position, lastRow, lastCol)

			if n1 != nil {
				nextMoves = append(nextMoves, *n1)
			}

			n2 := tryMoveDown(position, lastRow, lastCol)

			if n2 != nil {
				nextMoves = append(nextMoves, *n2)
			}
			return nextMoves

		case MovingUp:
			nextMove := tryMoveUp(position, lastRow, lastCol)

			if nextMove != nil {
				return []thing{*nextMove}
			}

		case MovingDown:
			nextMove := tryMoveDown(position, lastRow, lastCol)

			if nextMove != nil {
				return []thing{*nextMove}
			}
		}
	}

	if currentCell == '-' {
		switch position.direction {
		case MovingRight:
			nextMove := tryMoveRight(position, lastRow, lastCol)

			if nextMove != nil {
				return []thing{*nextMove}
			}
			return []thing{}
		case MovingLeft:
			nextMove := tryMoveLeft(position, lastRow, lastCol)

			if nextMove != nil {
				return []thing{*nextMove}
			}

			return []thing{}
		case MovingUp:
			nextMoves := []thing{}

			n1 := tryMoveRight(position, lastRow, lastCol)

			if n1 != nil {
				nextMoves = append(nextMoves, *n1)
			}

			n2 := tryMoveLeft(position, lastRow, lastCol)

			if n2 != nil {
				nextMoves = append(nextMoves, *n2)
			}
			return nextMoves

		case MovingDown:
			nextMoves := []thing{}

			n1 := tryMoveRight(position, lastRow, lastCol)

			if n1 != nil {
				nextMoves = append(nextMoves, *n1)
			}

			n2 := tryMoveLeft(position, lastRow, lastCol)

			if n2 != nil {
				nextMoves = append(nextMoves, *n2)
			}
			return nextMoves
		}

	}

	return []thing{}

}

func Run(matrix [][]rune, pendingPositions []thing) int {
	//fmt.Println(input)

	energizedCells := map[int]map[int]map[int]int{}

	for len(pendingPositions) > 0 {
		//fmt.Println("Pending positions:", len(pendingPositions))
		position := pendingPositions[0]
		pendingPositions = pendingPositions[1:]

		if _, ok := energizedCells[position.i]; !ok {
			energizedCells[position.i] = map[int]map[int]int{}
		}

		if _, ok := energizedCells[position.i][position.j]; !ok {
			energizedCells[position.i][position.j] = map[int]int{}
		}

		energizedCells[position.i][position.j][position.direction]++

		nextPositions := getNextPositions(position, matrix)

		for _, nextPosition := range nextPositions {
			if _, ok := energizedCells[nextPosition.i]; !ok {
				pendingPositions = append(pendingPositions, nextPosition)
				continue
			}
			if _, ok := energizedCells[nextPosition.i][nextPosition.j]; !ok {
				pendingPositions = append(pendingPositions, nextPosition)
				continue
			}
			if _, ok := energizedCells[nextPosition.i][nextPosition.j][nextPosition.direction]; !ok {
				pendingPositions = append(pendingPositions, nextPosition)
				continue
			}

		}

	}

	totalCells := 0
	for _, v := range energizedCells {
		totalCells += len(v)
	}

	return totalCells
}

func Run1() {

	initialPosition := thing{
		i:         0,
		j:         0,
		direction: MovingRight,
	}

	matrix := getMatrix(input)
	pendingPositions := []thing{initialPosition}

	result := Run(matrix, pendingPositions)

	fmt.Println("Result", result)

}

func Run2() {

	matrix := getMatrix(input)
	possibleBeamStart := []thing{}

	numRows := len(matrix)
	numCols := len(matrix[0])

	// Adding starting points from first and last row
	for j := 0; j < numCols; j++ {
		possibleBeamStart = append(possibleBeamStart, thing{
			i:         0,
			j:         j,
			direction: MovingDown,
		})
		possibleBeamStart = append(possibleBeamStart, thing{
			i:         numRows - 1,
			j:         j,
			direction: MovingUp,
		})
	}

	for i := 0; i < numRows; i++ {
		possibleBeamStart = append(possibleBeamStart, thing{
			i:         i,
			j:         0,
			direction: MovingRight,
		})
		possibleBeamStart = append(possibleBeamStart, thing{
			i:         i,
			j:         numCols - 1,
			direction: MovingLeft,
		})
	}
	result := 0

	for _, pos := range possibleBeamStart {
		result = int(math.Max(float64(result), float64(Run(matrix, []thing{pos}))))
	}

	fmt.Println("Result:", result)
}
