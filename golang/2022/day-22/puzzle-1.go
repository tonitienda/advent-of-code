package main

import (
	"advent/utils/array"
	"advent/utils/funct"
	"advent/utils/input"
	"fmt"
	"strconv"
	"strings"
)

type Point struct {
	row int
	col int
}

type Movement struct {
	row int
	col int
}

var Path = 0
var Wall = 1
var Nothing = -1
var Me = 99

var TurnRight = -1
var TurnLeft = -2

func StrToInt(str string) int {
	return funct.GetValue(strconv.Atoi(str))
}

func getValue(str string) int {
	switch str {
	case ".":
		return Path
	case "#":
		return Wall
	case " ":
		return Nothing
	default:
		panic("Value not supported:" + str)

	}
}

func getSymbol(val int) string {
	switch val {
	case Path:
		return "."
	case Wall:
		return "#"
	case Nothing:
		return " "
	case Me:
		return "@"
	default:
		panic("Value not supported:" + strconv.Itoa(val))

	}
}

func findLastCellByRow(row int, board [][]int) Point {
	for col := len(board[row]) - 1; col >= 0; col-- {
		if board[row][col] != Nothing {
			return Point{row: row, col: col}
		}
	}

	panic("")
}

func findLastCellByCol(col int, board [][]int) Point {
	for row := len(board) - 1; row >= 0; row-- {
		if col >= len(board[row]) {
			continue
		}
		if board[row][col] != Nothing {
			return Point{row: row, col: col}
		}
	}

	panic("...")
}

func findFirstCellByRow(row int, board [][]int) Point {
	for col, _ := range board[row] {
		if board[row][col] != Nothing {
			return Point{row: row, col: col}
		}
	}

	panic("First path could not be found in row " + strconv.Itoa(row))
}

func findFirstCellByCol(col int, board [][]int) Point {
	for row, _ := range board {
		if col >= len(board[row]) {
			continue
		}
		if board[row][col] != Nothing {
			return Point{row: row, col: col}
		}
	}

	panic("First path could not be found in col " + strconv.Itoa(col))
}

var FinalValues = map[string]int{
	"right": 0,
	"left":  2,
	"up":    3,
	"down":  1,
}

func changeMovement(current string, turn int) string {
	switch turn {
	// Clockwise
	case TurnRight:
		switch current {
		case "right":
			return "down"
		case "down":
			return "left"
		case "left":
			return "up"
		case "up":
			return "right"
		}
	case TurnLeft:
		switch current {
		case "right":
			return "up"
		case "up":
			return "left"
		case "left":
			return "down"
		case "down":
			return "right"
		}
	}

	panic("Not supported. Turn " + strconv.Itoa(turn) + ", curent " + current)
}

func processCommandLine(commandline string) []int {
	current := ""
	commands := []int{}

	chars := strings.Split(commandline, "")

	for _, c := range chars {
		if c != "L" && c != "R" {
			current += c
		} else {
			commands = append(commands, StrToInt(current))
			current = ""

			if c == "R" {
				commands = append(commands, TurnRight)
			} else {
				commands = append(commands, TurnLeft)

			}

		}
	}

	if current != "" {
		commands = append(commands, StrToInt(current))
	}
	return commands
}

func moveUp(p Point, board [][]int) Point {
	next := Point{row: p.row - 1, col: p.col}

	// If next point is out of limits or not walkable, try to wrap to the other side
	if next.row < 0 || board[next.row][next.col] == Nothing {
		next = findLastCellByCol(next.col, board)
	}

	// If the next point is a wall, return the current point
	if board[next.row][next.col] == Wall {
		return p
	}

	return next

}

func moveDown(p Point, board [][]int) Point {
	next := Point{row: p.row + 1, col: p.col}

	// If next point is out of limits or not walkable, try to wrap to the other side
	if next.row >= len(board) || next.col >= len(board[next.row]) || board[next.row][next.col] == Nothing {
		next = findFirstCellByCol(next.col, board)
	}

	// If the next point is a wall, return the current point
	if board[next.row][next.col] == Wall {
		return p
	}

	return next

}

func moveRight(p Point, board [][]int) Point {
	next := Point{row: p.row, col: p.col + 1}

	// If next point is out of limits or not walkable, try to wrap to the other side
	if next.col >= len(board[next.row]) || board[next.row][next.col] == Nothing {
		next = findFirstCellByRow(next.row, board)
	}

	// If the next point is a wall, return the current point
	if board[next.row][next.col] == Wall {
		return p
	}

	return next

}

func moveLeft(p Point, board [][]int) Point {
	next := Point{row: p.row, col: p.col - 1}

	// If next point is out of limits or not walkable, try to wrap to the other side
	if next.col < 0 || board[next.row][next.col] == Nothing {
		next = findLastCellByRow(next.row, board)
	}

	// If the next point is a wall, return the current point
	if board[next.row][next.col] == Wall {
		return p
	}

	return next

}

func move(p Point, direction string, board [][]int) Point {
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

var directionSymbol = map[string]string{
	"right": ">",
	"left":  "<",
	"up":    "^",
	"down":  "v",
}

func printBoard(board [][]int, directions map[Point]string) {
	for row, _ := range board {
		for col, _ := range board[row] {
			if board[row][col] == Me {
				fmt.Print("@")
				continue
			}

			direction, ok := directions[Point{row: row, col: col}]

			if ok {
				fmt.Print(directionSymbol[direction])
			} else {
				switch board[row][col] {
				case Wall:
					fmt.Print("#")
				case Path:
					fmt.Print(".")
				default:
					fmt.Print(" ")
				}
			}
		}
		fmt.Println()
	}
}

func main() {
	execType := "input"

	data := strings.Split(input.GetContents(2022, 22, execType+".txt"), "\n\n")

	boardDesc := data[0]
	commandline := data[1]
	//commandline = "10R1L2R16R6R13R6L41R41R9L3R33L39R30"

	board := array.Map(strings.Split(boardDesc, "\n"), func(str string) []int { return array.Map(strings.Split(str, ""), getValue) })
	commands := processCommandLine(commandline)

	//fmt.Println(commands)

	current := findFirstCellByRow(0, board)
	direction := "right"

	trail := map[Point]string{}
	for _, command := range commands {

		trail[current] = direction

		if command == TurnLeft || command == TurnRight {
			direction = changeMovement(direction, command)
		} else {
			for i := 0; i < command; i++ {
				// We can optimize this. If there is a wall we can skip some iterations
				current = move(current, direction, board)
				trail[current] = direction
			}
		}
		//fmt.Println()
	}

	//board[current.row][current.col] = Me
	//printBoard(board, trail)

	// Expected (rows index 1) row 6, col 8, direction 0 (right) 1000 * 6 + 4 * 8 + 0: 6032
	fmt.Println(current, direction)
	fmt.Println((current.row+1)*1000 + (current.col+1)*4 + FinalValues[direction])

}
