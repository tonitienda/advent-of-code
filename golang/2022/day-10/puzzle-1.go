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

func main() {
	commands := array.Map(input.GetLines(2022, 10, "input.txt"), func(str string) []string { return strings.Split(str, " ") })

	fmt.Println(commands)

	x := 1
	acc := 0
	cycle := 0
	for _, command := range commands {
		fmt.Println(command, "(", cycle, ")")

		if cycle > 220 {
			break
		}
		cycle++

		if cycle%40 == 20 {
			fmt.Println(cycle, "adding", x*cycle)
			acc += x * cycle
		}

		if command[0] == "addx" {
			x += StrToInt(command[1])
			fmt.Println("added", StrToInt(command[1]), "x", x)
			cycle++

			if cycle%40 == 20 {
				fmt.Println(cycle, "adding", x*cycle)

				acc += x * cycle
			}
		}
		fmt.Println("x:", x, "cycle", cycle, "=", x*cycle)
	}

	fmt.Println(acc)
}
