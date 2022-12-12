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

func within1Unit(head, tail [2]int) bool {

	// Overlaps
	if head[0] == tail[0] && head[1] == tail[1] {
		return true
	}

	// Same row
	if head[0] == tail[0] && math.Abs(float64(head[1]-tail[1])) == 1 {
		return true
	}

	// Same column
	if head[1] == tail[1] && math.Abs(float64(head[0]-tail[0])) == 1 {
		return true
	}

	// Diagonal
	if math.Abs(float64(head[1]-tail[1])) == 1 && math.Abs(float64(head[0]-tail[0])) == 1 {
		return true
	}

	return false
}
func applyCommand(direction string, units int, head, tail [2]int) ([2]int, [2]int, [][2]int) {
	switch direction {
	case "R":
		return moveR(units, head, tail)
	case "L":
		return moveL(units, head, tail)
	case "U":
		return moveU(units, head, tail)
	case "D":
		return moveD(units, head, tail)
	}

	visits := [][2]int{}

	return head, tail, visits
}

func moveR(units int, head, tail [2]int) ([2]int, [2]int, [][2]int) {
	visits := [][2]int{}

	for i := 0; i < units; i++ {
		head[1]++

		if !within1Unit(head, tail) {
			tail[1] = head[1] - 1
			tail[0] = head[0]

			newTail := [2]int{}
			copy(newTail[:], tail[:])
			visits = append(visits, newTail)
		}

	}
	return head, tail, visits
}

func moveU(units int, head, tail [2]int) ([2]int, [2]int, [][2]int) {
	visits := [][2]int{}

	for i := 0; i < units; i++ {
		head[0]--

		if !within1Unit(head, tail) {
			tail[0] = head[0] + 1
			tail[1] = head[1]

			newTail := [2]int{}
			copy(newTail[:], tail[:])
			visits = append(visits, newTail)
		}
	}

	return head, tail, visits
}

func moveL(units int, head, tail [2]int) ([2]int, [2]int, [][2]int) {
	visits := [][2]int{}

	for i := 0; i < units; i++ {
		head[1]--

		if !within1Unit(head, tail) {
			tail[1] = head[1] + 1
			tail[0] = head[0]

			newTail := [2]int{}
			copy(newTail[:], tail[:])
			visits = append(visits, newTail)
		}
	}

	return head, tail, visits
}

func moveD(units int, head, tail [2]int) ([2]int, [2]int, [][2]int) {
	visits := [][2]int{}

	for i := 0; i < units; i++ {
		head[0]++

		if !within1Unit(head, tail) {
			tail[0] = head[0] - 1
			tail[1] = head[1]

			newTail := [2]int{}
			copy(newTail[:], tail[:])
			visits = append(visits, newTail)
		}

	}

	return head, tail, visits
}

func main() {
	commands := array.Map(input.GetLines(2022, 9, "input.txt"), func(str string) []string { return strings.Split(str, " ") })

	head := [2]int{0, 0}
	tail := [2]int{0, 0}
	totalTailVisits := [][2]int{}

	tailVisits := [][2]int{}
	totalTailVisits = append(totalTailVisits, [2]int{0, 0})

	for _, command := range commands[:] {
		head, tail, tailVisits = applyCommand(command[0], StrToInt(command[1]), head, tail)

		totalTailVisits = append(totalTailVisits, tailVisits...)

	}

	uniqueVisits := map[string]bool{}

	for _, visit := range totalTailVisits {
		uniqueVisits[fmt.Sprintf("%d-%d", visit[0], visit[1])] = true
	}

	//fmt.Println(uniqueVisits)
	fmt.Println(len(uniqueVisits))

}
