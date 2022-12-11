package main

import (
	"advent/utils/array"
	"advent/utils/funct"
	"advent/utils/input"
	"fmt"
	"strconv"
	"strings"
)

func StrToInt(str string) int {
	return funct.GetValue(strconv.Atoi(str))
}

func printScreen(screen [40][40]string) {
	for i := 0; i < 40; i++ {
		for j := 0; j < 40; j++ {
			fmt.Print(screen[i][j])
		}
		fmt.Println()

	}
}

func main() {
	commands := array.Map(input.GetLines(2022, 10, "input.txt"), func(str string) []string { return strings.Split(str, " ") })

	fmt.Println(commands)
	screen := [40][40]string{}

	// for i := 0; i < 40; i++ {
	// 	for j := 0; j < 40; j++ {
	// 		screen[i][j] = "."
	// 	}

	// }

	x := 1
	cycle := 0
	for _, command := range commands {
		fmt.Println("begin executing", command)

		row := cycle / 40
		col := cycle % 40

		fmt.Println("CRT draws pixel in position", col)

		if cycle%40 >= x-1 && cycle%40 <= x+1 {
			screen[row][col] = "#"
		} else {
			screen[row][col] = "."

		}

		fmt.Println("Current CRT Draw", screen[row])

		cycle++

		if command[0] == "addx" {

			row := cycle / 40
			col := cycle % 40

			fmt.Println("CRT draws pixel in position", col)

			if cycle%40 >= x-1 && cycle%40 <= x+1 {
				screen[row][col] = "#"
			} else {
				screen[row][col] = "."

			}

			fmt.Println("Current CRT Draw", screen[row])

			x += StrToInt(command[1])

			fmt.Println("finish executing", command, "(Register X is now", x)
			cycle++

		}
	}

	printScreen(screen)
}
