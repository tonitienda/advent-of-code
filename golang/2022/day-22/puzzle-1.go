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

func findLastCellByRow(row int, board [][]int) (Point, bool) {
	for col := len(board[row]) - 1; col <= 0; col-- {
		if board[row][col] == Wall {
			return Point{row: row, col: col}, false
		}
		if board[row][col] == Path {
			return Point{row: row, col: col}, true
		}
	}

	panic("")
}

func findLastCellByCol(col int, board [][]int) (Point, bool) {
	for row := len(board) - 1; row <= 0; row-- {
		if board[row][col] == Wall {
			return Point{row: row, col: col}, false
		}
		if board[row][col] == Path {
			return Point{row: row, col: col}, true
		}
	}

	panic("")
}

func findFirstCellByRow(row int, board [][]int) (Point, bool) {
	for col, val := range board[row] {
		if val == Wall {
			return Point{row: row, col: col}, false
		}
		if val == Path {
			return Point{row: row, col: col}, true
		}
	}

	panic("First path could not be found in row " + strconv.Itoa(row))
}

func findFirstCellByCol(col int, board [][]int) (Point, bool) {
	for row, _ := range board {
		if board[row][col] == Wall {
			return Point{row: row, col: col}, false
		}
		if board[row][col] == Path {
			return Point{row: row, col: col}, true
		}
	}

	panic("First path could not be found in col " + strconv.Itoa(col))
}

var Movements = map[string]Movement{
	"right": {row: 0, col: 1},
	"left":  {row: 0, col: -1},
	"up":    {row: -1, col: 0},
	"down":  {row: 1, col: 0},
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
	return commands
}

func move(p Point, direction string, board [][]int) Point {

	m := Movements[direction]
	next := Point{
		row: p.row + m.row,
		col: p.col + m.col,
	}

	if next.row < 0 {
		next, _ = findLastCellByCol(next.col, board)
	} else if next.col < 0 {
		next, _ = findLastCellByRow(next.row, board)

	} else if next.row >= len(board) {
		next, _ = findFirstCellByCol(next.col, board)

	} else if next.col >= len(board[next.row]) {
		next, _ = findFirstCellByRow(next.row, board)

	}

	fmt.Println("current", p, "next", next)
	if board[next.row][next.col] == Nothing {
		switch direction {
		case "right":
			next2, ok := findFirstCellByRow(next.row, board)
			if ok {
				next = next2
			} else {
				return p
			}

		case "down":
			next2, ok := findFirstCellByCol(next.col, board)
			if ok {
				next = next2
			} else {
				return p
			}
		case "left":
			next2, ok := findLastCellByRow(next.row, board)
			if ok {
				next = next2
			} else {
				return p
			}

		case "up":
			next2, ok := findLastCellByCol(next.col, board)
			if ok {
				next = next2
			} else {
				return p
			}

		}
	}

	if board[next.row][next.col] == Wall {
		fmt.Println("Next cell is a wall. Returning", p)
		return p
	}
	fmt.Println("\tReturning", next)
	return next

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
	execType := "test"

	data := strings.Split(input.GetContents(2022, 22, execType+".txt"), "\n\n")

	boardDesc := data[0]
	commandline := data[1]

	board := array.Map(strings.Split(boardDesc, "\n"), func(str string) []int { return array.Map(strings.Split(str, ""), getValue) })
	commands := processCommandLine(commandline)

	fmt.Println(board)
	fmt.Println(commands)

	current, _ := findFirstCellByRow(0, board)
	direction := "right"

	trail := map[Point]string{}
	printBoard(board, trail)
	for _, command := range commands {

		trail[current] = direction

		fmt.Println(current)
		fmt.Println("command", command)
		if command == TurnLeft || command == TurnRight {
			direction = changeMovement(direction, command)
			fmt.Println("direction", direction)

		} else {
			fmt.Println("moving", command, "units")

			for i := 0; i < command; i++ {
				// We can optimize this. If there is a wall we can skip some iterations
				current = move(current, direction, board)
				trail[current] = direction
			}
		}
		fmt.Println()
	}

	printBoard(board, trail)

	// Expected (rows index 1) row 6, col 8, direction 0 (right) 1000 * 6 + 4 * 8 + 0: 6032
	fmt.Println(current, direction)
	fmt.Println((current.row+1)*1000 + (current.col+1)*4 + FinalValues[direction])

}
