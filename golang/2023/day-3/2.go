package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

//go:embed input.txt
var input string

func hasCoordinates(rowIdx, colIdx int, gearsCoordinates map[int]map[int][]int) bool {
	if _, ok := gearsCoordinates[rowIdx]; ok {
		if _, ok = gearsCoordinates[rowIdx][colIdx]; ok {
			return true
		}
	}

	return false
}

func getGearsNearby(rowIdx, colIdx int, gearsCoordinates map[int]map[int][]int) [][2]int {
	gearsNearby := [][2]int{}

	navigation := [][2]int{
		{rowIdx - 1, colIdx - 1},
		{rowIdx, colIdx - 1},
		{rowIdx + 1, colIdx - 1},
		{rowIdx - 1, colIdx},
		{rowIdx + 1, colIdx},
		{rowIdx - 1, colIdx + 1},
		{rowIdx, colIdx + 1},
		{rowIdx + 1, colIdx + 1},
	}

	for _, nav := range navigation {
		if hasCoordinates(nav[0], nav[1], gearsCoordinates) {
			gearsNearby = append(gearsNearby, [2]int{nav[0], nav[1]})
		}
	}

	return gearsNearby

}

func hasCoords(gearsNearby [][2]int, coords [2]int) bool {
	for _, gear := range gearsNearby {
		if gear == coords {
			return true
		}
	}

	return false
}

func main() {
	// 	test := `467..114..
	// ...*......
	// ..35..633.
	// ......#...
	// 617*......
	// .....+.58.
	// ..592.....
	// ......755.
	// ...$.*....
	// .664.598..`

	gearsCoordinates := map[int]map[int][]int{}

	for rowIdx, line := range strings.Split(input, "\n") {
		gearsCoordinates[rowIdx] = map[int][]int{}
		for colIdx, cell := range []rune(line) {

			if cell == '*' {
				gearsCoordinates[rowIdx][colIdx] = []int{}
			}
		}
	}

	withinNumber := false
	currentNumber := ""
	gearsNearby := [][2]int{}

	for rowIdx, line := range strings.Split(input, "\n") {
		for colIdx, cell := range []rune(line) {

			if unicode.IsDigit(cell) {
				withinNumber = true

				gearsNearbyCandidates := getGearsNearby(rowIdx, colIdx, gearsCoordinates)

				for _, gearCoords := range gearsNearbyCandidates {
					if !hasCoords(gearsNearby, gearCoords) {
						gearsNearby = append(gearsNearby, gearCoords)
					}
				}

				currentNumber += string(cell)
				// fmt.Println("currentNumber", currentNumber)

			} else {
				if withinNumber {
					n, err := strconv.Atoi(currentNumber)
					if err != nil {
						panic(err)
					}

					//fmt.Println("number", n, "gearsNearby", gearsNearby)

					for _, gearNearby := range gearsNearby {
						gearsCoordinates[gearNearby[0]][gearNearby[1]] = append(gearsCoordinates[gearNearby[0]][gearNearby[1]], n)
					}
				}

				gearsNearby = [][2]int{}

				currentNumber = ""
				withinNumber = false
			}
		}
	}

	total := 0
	for rowIdx, row := range gearsCoordinates {
		for colIdx, nums := range row {
			if len(nums) > 0 {
				fmt.Printf("[%d, %d] => [%v]\n", rowIdx, colIdx, nums)

				if len(nums) == 2 {
					total += nums[0] * nums[1]
				}
			}
		}
	}

	fmt.Println("Total:", total)

}
