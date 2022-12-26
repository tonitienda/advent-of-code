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

func findLastCellByRow(row int, board [][]int) Point {
	for col := len(board[row]) - 1; col <= 0; col-- {
		if board[row][col] == Path {
			return Point{row: row, col: col}
		}
	}

	panic("")
}

func findLastCellByCol(col int, board [][]int) Point {
	for row := len(board) - 1; row <= 0; row-- {
		if board[row][col] == Path {
			return Point{row: row, col: col}
		}
	}

	panic("")
}

func findFirstCellByRow(row int, board [][]int) Point {
	for col, val := range board[row] {
		if val == Path {
			return Point{row: row, col: col}
		}
	}

	panic("First path could not be found in row " + strconv.Itoa(row))
}

func findFirstCellByCol(col int, board [][]int) Point {
	for row, _ := range board {
		if board[row][col] == Path {
			return Point{row: row, col: col}
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
		next = findLastCellByCol(next.col, board)
	} else if next.col < 0 {
		next = findLastCellByRow(next.row, board)
	} else if next.row >= len(board) {
		next = findFirstCellByCol(next.row, board)
	} else if next.col >= len(board[next.row]) {
		next = findFirstCellByRow(next.row, board)
	}

	if board[next.row][next.col] == Nothing {
		switch direction {
		case "right":
			next = findFirstCellByRow(next.row, board)
		case "down":
			next = findFirstCellByCol(next.col, board)
		case "left":
			next = findLastCellByRow(next.row, board)
		case "up":
			next = findLastCellByCol(next.col, board)
		}
	}

	if board[next.row][next.col] == Wall {
		return p
	}

	return next

}

func wrapMovement(point Point, direction string, board [][]int) Point {
	if direction == "right" {
		return findFirstCellByRow(point.row, board)
	}

	if direction == "down" {
		return findFirstCellByCol(point.col, board)
	}

	if direction == "left" {
		return findLastCellByRow(point.row, board)
	}

	if direction == "up" {
		return findLastCellByCol(point.col, board)
	}

	panic("")
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

	current := findFirstCellByRow(0, board)
	direction := "right"

	for _, command := range commands {

		if command == TurnLeft || command == TurnRight {
			direction = changeMovement(direction, command)
		} else {
			for i := 0; i < command; i++ {
				// We can optimize this. If there is a wall we can skip some iterations
				current = move(current, direction, board)
			}
		}
	}

	fmt.Println(current, direction)

}
