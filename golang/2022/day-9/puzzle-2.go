package main

import (
	"advent/utils/array"
	"advent/utils/funct"
	"advent/utils/input"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func StrToInt(str string) int {
	return funct.GetValue(strconv.Atoi(str))
}

func within1Unit(knot1, knot2 [2]int) bool {

	// Overlaps
	if knot1[0] == knot2[0] && knot1[1] == knot2[1] {
		return true
	}

	// Same row
	if knot1[0] == knot2[0] && math.Abs(float64(knot1[1]-knot2[1])) == 1 {
		return true
	}

	// Same column
	if knot1[1] == knot2[1] && math.Abs(float64(knot1[0]-knot2[0])) == 1 {
		return true
	}

	// Diagonal
	if math.Abs(float64(knot1[1]-knot2[1])) == 1 && math.Abs(float64(knot1[0]-knot2[0])) == 1 {
		return true
	}

	return false
}

func calculateNextMove(move, knot1, knot2 [2]int) [2]int {
	if within1Unit(knot1, knot2) {
		return [2]int{0, 0}
	}

	if move[0] == 0 && move[1] == 0 {
		return [2]int{0, 0}
	}

	newPos := [2]int{}

	// Lateral Move: knot2 needs to go to the same row as knot1
	if move[0] == 0 {
		// Right move
		if move[1] > 0 {
			newPos = [2]int{knot1[0], knot1[1] - 1}
		} else {
			newPos = [2]int{knot1[0], knot1[1] + 1}
		}
	} else { // Vertical Move
		// Down move
		if move[0] > 0 {
			newPos = [2]int{knot1[0] - 1, knot1[1]}
		} else {
			newPos = [2]int{knot1[0] + 1, knot1[1]}
		}
	}

	return [2]int{newPos[0] - knot2[0], newPos[1] - knot2[1]}

}

func getMovesFromCommand(direction string, units int) [][2]int {
	moves := [][2]int{}
	moveToAdd := [2]int{}

	switch direction {
	case "R":
		moveToAdd = [2]int{0, 1}
	case "L":
		moveToAdd = [2]int{0, -1}
	case "U":
		moveToAdd = [2]int{-1, 0}
	case "D":
		moveToAdd = [2]int{1, 0}
	}
	for i := 0; i < units; i++ {
		moves = append(moves, moveToAdd)
	}

	return moves
}

func main() {
	commands := array.Map(input.GetLines(2022, 9, "input.txt"), func(str string) []string { return strings.Split(str, " ") })

	knots := make([][2]int, 10)

	for i := 0; i < 10; i++ {
		knots[i] = [2]int{0, 0}
	}

	//fmt.Println(commands)
	fmt.Println(knots)

	uniqueVisits := map[string]bool{}

	for _, command := range commands {
		fmt.Println(command)
		moves := getMovesFromCommand(command[0], StrToInt(command[1]))

		for _, move := range moves {
			//fmt.Println("move", move, "head", knots[0], "tail", knots[len(knots)-1])

			knots[0][0] += move[0]
			knots[0][1] += move[1]
			nextMove := move
			for i := 1; i < len(knots); i++ {
				nextMove = calculateNextMove(nextMove, knots[i-1], knots[i])
				if nextMove[0] == 0 && nextMove[1] == 0 {
					break
				}
				knots[i][0] += nextMove[0]
				knots[i][1] += nextMove[1]
				//fmt.Println("move", move, "knot1", knots[i-1], "knot2", knots[i], "nextMove", nextMove)

			}
			fmt.Println("tail:", knots[len(knots)-1])
			uniqueVisits[fmt.Sprintf("%d-%d", knots[len(knots)-1][0], knots[len(knots)-1][1])] = true

		}
	}

	// fmt.Println(calculateNextMove([2]int{0, 1}, [2]int{0, 0}, [2]int{0, 0}), " = ", [2]int{0, 0})
	// fmt.Println(calculateNextMove([2]int{0, 1}, [2]int{0, 1}, [2]int{0, 0}), " = ", [2]int{0, 0})
	// fmt.Println(calculateNextMove([2]int{0, 1}, [2]int{0, 2}, [2]int{0, 0}), " = ", [2]int{0, 1})
	// fmt.Println(calculateNextMove([2]int{0, 1}, [2]int{0, 3}, [2]int{0, 1}), " = ", [2]int{0, 1})

	// totalTailVisits := [][2]int{}

	// tailVisits := [][2]int{}
	// totalTailVisits = append(totalTailVisits, [2]int{0, 0})

	// knots, tailVisits = move(commands[0][0], StrToInt(commands[0][1]), knots)
	// totalTailVisits = append(totalTailVisits, tailVisits...)

	// knots, tailVisits = move(commands[1][0], StrToInt(commands[1][1]), knots)
	// totalTailVisits = append(totalTailVisits, tailVisits...)

	// for _, command := range commands[:] {
	// 	knots, tailVisits = move(command[0], StrToInt(command[1]), knots)
	// 	fmt.Println(knots)

	// 	totalTailVisits = append(totalTailVisits, tailVisits...)

	// }

	// uniqueVisits := map[string]bool{}

	// for _, visit := range totalTailVisits {
	// 	uniqueVisits[fmt.Sprintf("%d-%d", visit[0], visit[1])] = true
	// }

	// fmt.Println(totalTailVisits)
	// fmt.Println(uniqueVisits)

	// fmt.Println(len(totalTailVisits))
	fmt.Println(len(uniqueVisits))

}
