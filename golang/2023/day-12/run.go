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

func accordingToSpecs2(springs []rune, specs []int) bool {
	//fmt.Println(string(springs), specs)
	itemsInCurrentGroup := 0
	currentGroupIdx := 0
	expectedNoGroups := len(specs)
	totalChars := len(springs)

	for idx, r := range springs {
		//	fmt.Println("Evaluating ", idx, ":", string(r))
		if r == '?' {
			// TODO - Here we are returning true very naively.
			// We should veryfiy that the position is still valid.
			// For example, for each "group" in specs, we need at least 2 character
			// numCharsRequired = len(specs)*2 -1
			// Maybe we have reached a ?, the the remaining chars (including ?)
			// is not enough to form the required groups.

			//fmt.Println("\t? => true")
			remainingChars := totalChars - idx
			remainingGroups := expectedNoGroups - currentGroupIdx

			// TODO - check if this formula is correct
			// TODO - We are still permissive here not following the *
			// but we can cut out some possibilities
			return remainingChars >= remainingGroups
		}

		if r == '#' {
			itemsInCurrentGroup++

			// More groups than expected
			// We could check this only once, when the first # of the group is found
			if currentGroupIdx > expectedNoGroups-1 {
				//	fmt.Println("\tcurrentGroupIdx > expectedNoGroups-1 => false")

				return false
			}

			if itemsInCurrentGroup > specs[currentGroupIdx] {
				//	fmt.Println("\titemsInCurrentGroup > specs[currentGroupIdx] => false")
				return false
			}

			continue
		}

		if r == '.' && itemsInCurrentGroup > 0 {
			if itemsInCurrentGroup != specs[currentGroupIdx] {
				//	fmt.Println("\titemsInCurrentGroup != specs[currentGroupIdx] => false")
				return false
			}
			currentGroupIdx++
			itemsInCurrentGroup = 0

		}
	}

	// No more items remaining means that all groups where checked.
	if itemsInCurrentGroup == 0 {
		// We have a currentGroupIdx++ to prepare for the next group
		// but there are no items in that last group, so we need to discard one ++
		if currentGroupIdx == expectedNoGroups {
			//	fmt.Println("\titemsInCurrentGroup == 0 => true")
			return true
		} else {
			//fmt.Println("\titemsInCurrentGroup == 0 => false (no groups does not match)")
			return false
		}
	}

	if currentGroupIdx != expectedNoGroups-1 {
		//fmt.Println("\tcurrentGroupIdx != expectedNoGroups-1 => false")

		return false
	}

	lastItemOk := itemsInCurrentGroup == specs[currentGroupIdx]

	// if itemsInCurrentGroup == 0 {
	// 	fmt.Println("\tlastItemOk => ", lastItemOk)
	// }
	return lastItemOk

}

func getArrangements(springs []rune, specs []int, unknownIndices []int) int {
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
		if accordingToSpecs2(springs, specs) {
			// fmt.Println("According to specs\n")
			totalValidArrangements++
		} else {
			// fmt.Println("Not according to specs\n")
		}
	}

	// fmt.Println()
	return totalValidArrangements

}

func getArrangements2(springs []rune, specs []int, unknownIndices []int) int {
	if !accordingToSpecs2(springs, specs) {
		//fmt.Println("\t", string(springs), false)
		return 0
	}

	// No more unknown positions
	if len(unknownIndices) == 0 {
		//fmt.Println("Correct: ", string(springs), ":", specs)
		return 1
	}

	//totalValidArrangements := 0

	//for _, idx := range unknownIndices {
	idx := unknownIndices[0]
	springs[idx] = '.'
	n1 := getArrangements2(springs, specs, unknownIndices[1:])
	//fmt.Println("n1", string(springs), ":", specs, "=>", n1)

	springs[idx] = '#'
	n2 := getArrangements2(springs, specs, unknownIndices[1:])
	//fmt.Println("n2", string(springs), ":", specs, "=>", n2)

	springs[idx] = '?'

	return n1 + n2

	// }

	//return totalValidArrangements

}

func getArrangementsByRow(row string) int {
	springs, specs := getSpringsAndSpecs(row)

	unknownIndices := []int{}

	for idx, r := range springs {
		if r == '?' {
			unknownIndices = append(unknownIndices, idx)
		}
	}

	return getArrangements2(springs, specs, unknownIndices)
}

func getSpringsAndSpecs(row string) ([]rune, []int) {
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

	return springs, specs

}

func getArrangementsByLongRow(row string) int {
	springs, specs := getSpringsAndSpecs(row)

	longSprings := []rune{}
	longSprings = append(longSprings, springs...)
	longSprings = append(longSprings, '?')
	longSprings = append(longSprings, springs...)
	longSprings = append(longSprings, '?')
	longSprings = append(longSprings, springs...)
	longSprings = append(longSprings, '?')
	longSprings = append(longSprings, springs...)
	longSprings = append(longSprings, '?')
	longSprings = append(longSprings, springs...)

	longSpecs := []int{}
	longSpecs = append(longSpecs, specs...)
	longSpecs = append(longSpecs, specs...)
	longSpecs = append(longSpecs, specs...)
	longSpecs = append(longSpecs, specs...)
	longSpecs = append(longSpecs, specs...)

	unknownIndices := []int{}

	for idx, r := range longSprings {
		if r == '?' {
			unknownIndices = append(unknownIndices, idx)
		}
	}

	//fmt.Println("\t", string(longSprings), longSpecs)

	return getArrangements2(longSprings, longSpecs, unknownIndices)
}

func Run1() {
	//fmt.Println(test)

	rows := strings.Split(input, "\n")

	totalArrangements := 0
	for _, row := range rows {
		arrangements := getArrangementsByRow(row)
		//fmt.Println(row, "=>", arrangements)
		totalArrangements += arrangements
	}

	fmt.Println("Total arrangements:", totalArrangements)
}

func Run2() {
	//fmt.Println(test)

	rows := strings.Split(input, "\n")

	totalArrangements := 0
	for idx, row := range rows {
		fmt.Println("Evaluating", idx+1, "/", len(rows))
		arrangements := getArrangementsByLongRow(row)
		fmt.Println("\t=>", arrangements)
		totalArrangements += arrangements
	}

	fmt.Println("Total arrangements:", totalArrangements)
}
