package main

import (
	_ "embed"
	"fmt"
	"math"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

const test = `Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green`

func main() {

	total := 0

	for _, line := range strings.Split(input, "\n") {
		parts := strings.Split(line, ":")

		gameNumber, err := strconv.Atoi(strings.Replace(parts[0], "Game ", "", 1))

		if err != nil {
			panic(err)
		}

		games := strings.Split(parts[1], ";")
		minCubes := map[string]int{}
		minCubes["red"] = 1
		minCubes["blue"] = 1
		minCubes["green"] = 1

		for _, game := range games {

			gameParts := strings.Split(game, ",")

			for _, gamePart := range gameParts {
				//fmt.Printf("Game: %d, gamePart %s\n", gameNumber, gamePart)

				numColor := strings.Split(strings.Trim(gamePart, " "), " ")
				//fmt.Println("numColor", numColor)
				color := numColor[1]

				num, err := strconv.Atoi(numColor[0])

				if err != nil {
					panic(err)
				}

				minCubes[color] = int(math.Max(float64(minCubes[color]), float64(num)))

			}

		}
		//fmt.Println("Game", gameNumber, ":", minCubes)
		totalGame := minCubes["red"] * minCubes["blue"] * minCubes["green"]

		//fmt.Println(gameNumber, "=>", totalGame)

		total += totalGame

	}

	fmt.Println("total", total)

}
