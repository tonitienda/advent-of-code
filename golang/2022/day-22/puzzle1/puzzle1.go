package puzzle1

import (
	"advent/2022/day-22/common"
	"advent/utils/array"
	"advent/utils/input"
	"fmt"
	"strconv"
	"strings"
)

func findLastCellByRow(row int, board [][]int) common.Point {
	for col := len(board[row]) - 1; col >= 0; col-- {
		if board[row][col] != common.Nothing {
			return common.Point{Row: row, Col: col}
		}
	}

	panic("")
}

func findLastCellByCol(col int, board [][]int) common.Point {
	for row := len(board) - 1; row >= 0; row-- {
		if col >= len(board[row]) {
			continue
		}
		if board[row][col] != common.Nothing {
			return common.Point{Row: row, Col: col}
		}
	}

	panic("...")
}

func findFirstCellByRow(row int, board [][]int) common.Point {
	for col, _ := range board[row] {
		if board[row][col] != common.Nothing {
			return common.Point{Row: row, Col: col}
		}
	}

	panic("First path could not be found in row " + strconv.Itoa(row))
}

func findFirstCellByCol(col int, board [][]int) common.Point {
	for row, _ := range board {
		if col >= len(board[row]) {
			continue
		}
		if board[row][col] != common.Nothing {
			return common.Point{Row: row, Col: col}
		}
	}

	panic("First path could not be found in col " + strconv.Itoa(col))
}

func moveUp(p common.Point, board [][]int) common.Point {
	next := common.Point{Row: p.Row - 1, Col: p.Col}

	// If next point is out of limits or not walkable, try to wrap to the other side
	if next.Row < 0 || board[next.Row][next.Col] == common.Nothing {
		next = findLastCellByCol(next.Col, board)
	}

	// If the next point is a wall, return the current point
	if board[next.Row][next.Col] == common.Wall {
		return p
	}

	return next

}

func moveDown(p common.Point, board [][]int) common.Point {
	next := common.Point{Row: p.Row + 1, Col: p.Col}

	// If next point is out of limits or not walkable, try to wrap to the other side
	if next.Row >= len(board) || next.Col >= len(board[next.Row]) || board[next.Row][next.Col] == common.Nothing {
		next = findFirstCellByCol(next.Col, board)
	}

	// If the next point is a wall, return the current point
	if board[next.Row][next.Col] == common.Wall {
		return p
	}

	return next

}

func moveRight(p common.Point, board [][]int) common.Point {
	next := common.Point{Row: p.Row, Col: p.Col + 1}

	// If next point is out of limits or not walkable, try to wrap to the other side
	if next.Col >= len(board[next.Row]) || board[next.Row][next.Col] == common.Nothing {
		next = findFirstCellByRow(next.Row, board)
	}

	// If the next point is a wall, return the current point
	if board[next.Row][next.Col] == common.Wall {
		return p
	}

	return next

}

func moveLeft(p common.Point, board [][]int) common.Point {
	next := common.Point{Row: p.Row, Col: p.Col - 1}

	// If next point is out of limits or not walkable, try to wrap to the other side
	if next.Col < 0 || board[next.Row][next.Col] == common.Nothing {
		next = findLastCellByRow(next.Row, board)
	}

	// If the next point is a wall, return the current point
	if board[next.Row][next.Col] == common.Wall {
		return p
	}

	return next

}

func move(p common.Point, direction string, board [][]int) common.Point {
	switch direction {
	case "up":
		return moveUp(p, board)
	case "right":
		return moveRight(p, board)
	case "down":
		return moveDown(p, board)
	case "left":
		return moveLeft(p, board)
	}

	panic("Direction not supported")

}

func Run() {
	execType := "input"

	data := strings.Split(input.GetContents(2022, 22, execType+".txt"), "\n\n")

	boardDesc := data[0]
	commandline := data[1]
	//commandline = "10R1L2R16R6R13R6L41R41R9L3R33L39R30"

	board := array.Map(strings.Split(boardDesc, "\n"), func(str string) []int { return array.Map(strings.Split(str, ""), common.GetValue) })
	commands := common.ProcessCommandLine(commandline)

	//fmt.Println(commands)

	current := findFirstCellByRow(0, board)
	direction := "right"

	trail := map[common.Point]string{}
	for _, command := range commands {

		trail[current] = direction

		if command == common.TurnLeft || command == common.TurnRight {
			direction = common.ChangeDirection(direction, command)
		} else {
			for i := 0; i < command; i++ {
				// We can optimize this. If there is a wall we can skip some iterations
				current = move(current, direction, board)
				trail[current] = direction
			}
		}
		//fmt.Println()
	}

	//board[current.Row][current.Col] = Me
	//printBoard(board, trail)

	// Expected (rows index 1) row 6, col 8, direction 0 (right) 1000 * 6 + 4 * 8 + 0: 6032
	fmt.Println(current, direction)
	fmt.Println((current.Row+1)*1000 + (current.Col+1)*4 + common.FinalValues[direction])

}
