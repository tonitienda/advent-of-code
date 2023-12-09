package y2023d9

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed test.txt
var test string

//go:embed input.txt
var input string

func getDiffValues(values []int) ([]int, bool) {
	diffValues := []int{}
	done := true

	for i := 0; i < len(values)-1; i++ {
		diff := values[i+1] - values[i]

		if diff != 0 {
			done = false
		}

		diffValues = append(diffValues, diff)
	}

	return diffValues, done

}

func resolveNextValue(values []int) int {
	lastValuesSum := 0

	for {
		fmt.Println(values)
		lastValuesSum += values[len(values)-1]

		diff, done := getDiffValues(values)

		if done {
			break
		}

		values = diff
	}

	return lastValuesSum
}

func resolvePreviousValue(values []int) int {
	firstValues := []int{}

	for {
		firstValues = append(firstValues, values[0])

		diff, done := getDiffValues(values)

		if done {
			break
		}

		values = diff
	}

	result := 0
	for i := len(firstValues) - 1; i >= 0; i-- {
		result = firstValues[i] - result
	}
	return result
}

func Run(resolve func([]int) int) {
	lines := strings.Split(input, "\n")

	result := 0
	for _, line := range lines {
		valuesstr := strings.Split(line, " ")

		values := []int{}

		for _, str := range valuesstr {
			val, err := strconv.Atoi(str)

			if err != nil {
				panic(err)
			}

			values = append(values, val)
		}

		result += resolve(values)
	}
	fmt.Println("Result:", result)

}

func Run1() {
	Run(resolveNextValue)
}

func Run2() {
	Run(resolvePreviousValue)
}
