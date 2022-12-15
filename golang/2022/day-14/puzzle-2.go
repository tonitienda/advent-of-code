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

func getPoints(str string) [][2]int {
	points := [][2]int{}

	coords := strings.Split(str, " -> ")

	for _, coord := range coords {
		nums := array.Map(strings.Split(coord, ","), StrToInt)

		points = append(points, [2]int{nums[0], nums[1]})
	}

	return points

}

func findCollissionPoint(posX, posY int, lines [][][2]int, stones [][2]int) ([2]int, bool) {

	i := posY
	for {
		if isBlocked(posX, i, lines, stones) {
			return [2]int{posX, i}, true
		}
		i++
	}

	//return [2]int{}, false
}

func isBlocked(posX, posY int, lines [][][2]int, stones [][2]int) bool {

	return isBlockedByStone(posX, posY, stones) || isBlockedByLines(posX, posY, lines)
}

func isBlockedByStone(posX, posY int, stones [][2]int) bool {
	// Start with stones
	for _, stone := range stones {
		if posX == stone[0] && posY == stone[1] {
			//fmt.Println("Blocked by a stone at ", stone)
			return true
		}
	}

	return false
}

func isBlockedByLines(posX, posY int, lines [][][2]int) bool {

	for _, lineGroup := range lines {
		for i := 0; i < len(lineGroup)-1; i++ {
			segmentPoint1 := lineGroup[i]
			segmentPoint2 := lineGroup[i+1]

			// Vertical line
			if segmentPoint1[0] == segmentPoint2[0] {
				if segmentPoint1[0] == posX && (segmentPoint1[1] >= posY && segmentPoint2[1] <= posY || segmentPoint1[1] <= posY && segmentPoint2[1] >= posY) {
					return true
				}
			} else if segmentPoint1[1] == segmentPoint2[1] { // Horizontal line
				if segmentPoint1[1] == posY && (segmentPoint1[0] >= posX && segmentPoint2[0] <= posX || segmentPoint1[0] <= posX && segmentPoint2[0] >= posX) {
					//fmt.Println(posX, ",", posY, ": Collided with ", segmentPoint1[0], ",", segmentPoint1[1], "|", segmentPoint2[0], ",", segmentPoint2[1])
					return true
				}
			} else {
				panic("Diagonal lines not supported")
			}
		}
	}

	return false
}

func findFinalPosition(lines [][][2]int, stones [][2]int) ([2]int, bool) {
	// Until we find the final position of the stone
	fallingY := 0
	fallingX := 500

	for {
		//fmt.Println()
		//fmt.Println("Trying with ", fallingX, fallingY, stones)

		collisionPoint, ok := findCollissionPoint(fallingX, fallingY, lines, stones)

		if !ok {
			panic("Collission not found")
		}

		// If the stone stops at the entrance
		if collisionPoint[0] == 500 && collisionPoint[1] == 0 {
			//fmt.Println("The stone did not collide", stones, fallingX)
			return [2]int{}, false
		}

		//fmt.Println("Stone rest:", collisionPoint)

		// If it is not blocked from the left we fall from there
		if !isBlocked(fallingX-1, collisionPoint[1], lines, stones) {
			fallingY = collisionPoint[1]
			fallingX--
			continue
		}

		// If it is not blocked from the right we fall from there
		if !isBlocked(fallingX+1, collisionPoint[1], lines, stones) {
			fallingY = collisionPoint[1]
			fallingX++
			continue
		}

		// if it is blocked in both, return the position of the stone.
		// Which is one Y above the collission point

		stonePos := [2]int{collisionPoint[0], collisionPoint[1] - 1}

		return stonePos, true

	}
}

func printGrid(lines [][][2]int, stones [][2]int) {

	for i := 0; i < 50; i++ {
		gridLine := ""

		for j := 480; j < 520; j++ {
			if isBlockedByLines(j, i, lines) {
				gridLine += "#"
			} else if isBlockedByStone(j, i, stones) {
				gridLine += "o"
			} else {
				gridLine += "."
			}
		}
		fmt.Println(gridLine)
	}
}

func main() {

	data := input.GetLines(2022, 14, "input.txt")
	lines := array.Map(data, getPoints)

	maxY := 0

	for _, line := range lines {
		for _, point := range line {
			if maxY < point[1] {
				maxY = point[1]
			}
		}
	}

	maxY = maxY + 2

	bottomLine := [][2]int{[2]int{int(math.Inf(-1)), int(maxY)}, [2]int{int(math.Inf(1)), int(maxY)}}
	lines = append(lines, bottomLine)

	fallenStones := [][2]int{}

	// For each stone
	for {

		stonePos, ok := findFinalPosition(lines, fallenStones)

		if ok {
			fallenStones = append(fallenStones, stonePos)
			//fmt.Println("Stone added:", fallenStones)
			fmt.Println(len(fallenStones))

		} else {
			break
		}

	}
	//fmt.Println("Printing grid")
	//fmt.Println(lines)
	//fmt.Println(fallenStones)
	//fmt.Println(len(fallenStones))

	printGrid(lines, fallenStones)
	fmt.Println(len(fallenStones))

}
