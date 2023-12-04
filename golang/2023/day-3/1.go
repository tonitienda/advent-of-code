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

func isThereSymbolNearby(rowIdx, colIdx int, symbolCoordinates map[int]map[int]bool) bool {
	return symbolCoordinates[rowIdx-1][colIdx-1] ||
		symbolCoordinates[rowIdx][colIdx-1] ||
		symbolCoordinates[rowIdx+1][colIdx-1] ||
		symbolCoordinates[rowIdx-1][colIdx] ||
		symbolCoordinates[rowIdx+1][colIdx] ||
		symbolCoordinates[rowIdx-1][colIdx+1] ||
		symbolCoordinates[rowIdx][colIdx+1] ||
		symbolCoordinates[rowIdx+1][colIdx+1]
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

	symbolCoordinates := map[int]map[int]bool{}

	symbolCoordinates[-1] = map[int]bool{}
	symbolCoordinates[len(strings.Split(input, "\n"))] = map[int]bool{}

	for rowIdx, line := range strings.Split(input, "\n") {
		symbolCoordinates[rowIdx] = map[int]bool{}
		for colIdx, cell := range []rune(line) {

			if !unicode.IsDigit(cell) && !(cell == '.') {
				symbolCoordinates[rowIdx][colIdx] = true
			}
		}
	}

	withinNumber := false
	hasSymbolNearby := false
	currentNumber := ""
	totalSum := 0

	for rowIdx, line := range strings.Split(input, "\n") {
		for colIdx, cell := range []rune(line) {

			if unicode.IsDigit(cell) {
				withinNumber = true
				hasSymbolNearby = hasSymbolNearby || isThereSymbolNearby(rowIdx, colIdx, symbolCoordinates)
				currentNumber += string(cell)
			} else {
				if withinNumber && hasSymbolNearby {
					n, err := strconv.Atoi(currentNumber)
					if err != nil {
						panic(err)
					}

					fmt.Println(n)

					totalSum += n
				}

				hasSymbolNearby = false
				currentNumber = ""
				withinNumber = false
			}
		}
	}

	fmt.Println("Total", totalSum)
}
