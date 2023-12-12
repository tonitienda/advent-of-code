package y2023d12

import (
	_ "embed"
	"fmt"
	"math"
	"strconv"
	"strings"
)

//go:embed test.txt
var test string

//go:embed input.txt
var input string

func accordingToSpecs(springs []rune, specs []int) bool {
	groups := []int{}
	currentGroupItems := 0

	for _, r := range springs {
		if r == '.' && currentGroupItems > 0 {
			groups = append(groups, currentGroupItems)
			currentGroupItems = 0
		} else if r == '#' {
			currentGroupItems++
		}
	}

	if currentGroupItems > 0 {
		groups = append(groups, currentGroupItems)
	}

	//fmt.Println(groups)

	if len(groups) != len(specs) {
		return false
	}

	for idx, g := range groups {
		if g != specs[idx] {
			return false
		}
	}

	return true

}

func getArrangements(row string) int {
	data := strings.Split(row, " ")

	springs := []rune(data[0])
	specs := []int{}

	for _, n := range strings.Split(data[1], ",") {
		v, err := strconv.Atoi(n)

		if err != nil {
			panic(err)
		}

		specs = append(specs, v)
	}

	unknownIndices := []int{}

	for idx, r := range springs {
		if r == '?' {
			unknownIndices = append(unknownIndices, idx)
		}
	}

	//fmt.Println(springs, specs, "=>", unknownIndices)

	//fmt.Println()
	//fmt.Println(string(springs))

	totalValidArrangements := 0
	for i := 0; i < int(math.Pow(2, float64(len(unknownIndices)))); i++ {
		for j := 0; j < len(unknownIndices); j++ {
			idx := unknownIndices[j]
			bitIndex := int(math.Pow(2, float64(j)))
			if bitIndex&i == 0 {
				springs[idx] = '#'
			} else {
				springs[idx] = '.'
			}
		}

		// fmt.Println(string(springs))
		if accordingToSpecs(springs, specs) {
			totalValidArrangements++
		}
	}

	// fmt.Println()
	return totalValidArrangements
}

func Run1() {
	//fmt.Println(test)

	rows := strings.Split(input, "\n")

	totalArrangements := 0
	for _, row := range rows {
		totalArrangements += getArrangements(row)
	}

	fmt.Println("Total arrangements:", totalArrangements)
}
