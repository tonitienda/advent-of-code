package common

import (
	"advent/utils/funct"
	"fmt"
	"strconv"
	"strings"
)

type Point struct {
	Row int
	Col int
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

func GetValue(str string) int {
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

func GetSymbol(val int) string {
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

var FinalValues = map[string]int{
	"right": 0,
	"left":  2,
	"up":    3,
	"down":  1,
}

func ChangeDirection(current string, turn int) string {
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

func ProcessCommandLine(commandline string) []int {
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

var DirectionSymbol = map[string]string{
	"right": ">",
	"left":  "<",
	"up":    "^",
	"down":  "v",
}

func PrintBoard(board [][]int, directions map[Point]string) {
	for row, _ := range board {
		for col, _ := range board[row] {
			if board[row][col] == Me {
				fmt.Print("@")
				continue
			}

			direction, ok := directions[Point{Row: row, Col: col}]

			if ok {
				fmt.Print(DirectionSymbol[direction])
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
