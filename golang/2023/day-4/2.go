package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func countFindings(array1, array2 []int) int {
	total := 0

	for _, a1 := range array1 {
		for _, a2 := range array2 {
			if a1 == a2 {
				total++
				break
			}
		}
	}

	return total
}

func toNumbers(line string) []int {
	numbers := []int{}
	for _, part := range strings.Split(line, " ") {
		if part != "" {
			val, err := strconv.Atoi(strings.Trim(part, " "))

			if err != nil {
				panic(err)
			}

			numbers = append(numbers, val)
		}
	}

	return numbers
}

func main() {

	// 	test := `Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
	// Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
	// Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
	// Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
	// Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
	// Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11`

	games := strings.Split(input, "\n")
	scratchcardTotalCopies := make([]int, len(games))

	for i := 0; i < len(scratchcardTotalCopies); i++ {
		scratchcardTotalCopies[i] = 1
	}

	for idx, game := range games {
		gameSplit := strings.Split(game, "|")

		cardNumbers := toNumbers(strings.Split(gameSplit[0], ":")[1])
		winingNumbers := toNumbers(gameSplit[1])

		fmt.Println(cardNumbers, "|", winingNumbers)

		count := countFindings(cardNumbers, winingNumbers)

		//fmt.Println("Found", count, "coincidences")

		if count > 0 {
			//fmt.Println("scratchcardTotalCopies", scratchcardTotalCopies, "count", count)

			//totalPoints += int(math.Pow(2, float64(count-1)))
			for i := idx + 1; i < len(scratchcardTotalCopies) && i <= count+idx; i++ {
				scratchcardTotalCopies[i] += scratchcardTotalCopies[idx]
			}

			//fmt.Println("scratchcardTotalCopies", scratchcardTotalCopies)
		}

	}

	//fmt.Println("scratchcardTotalCopies", scratchcardTotalCopies)

	total := 0

	for _, val := range scratchcardTotalCopies {
		total += val
	}

	fmt.Println("Total:", total)
}
