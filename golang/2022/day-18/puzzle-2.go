package main

import (
	"advent/utils/array"
	"advent/utils/funct"
	"advent/utils/input"
	"fmt"
	"strconv"
	"strings"
)

func StrToInt(str string) int {
	return funct.GetValue(strconv.Atoi(str))
}

func isEqual(c1, c2 []int) bool {
	return c1[0] == c2[0] && c1[1] == c2[1] && c1[2] == c2[2]
}

func isContact(c1, c2 []int) bool {

	if c1[0] == c2[0] && c1[1] == c2[1] && (c1[2]-c2[2] == -1 || c1[2]-c2[2] == 1) {
		return true
	}
	if c1[0] == c2[0] && c1[2] == c2[2] && (c1[1]-c2[1] == -1 || c1[1]-c2[1] == 1) {
		return true
	}

	if c1[2] == c2[2] && c1[1] == c2[1] && (c1[0]-c2[0] == -1 || c1[0]-c2[0] == 1) {
		return true
	}

	return false
}

func contacts(cube []int, cubes [][]int) int {

	contacts := 0
	for _, cube2 := range cubes {
		if !isEqual(cube, cube2) && isContact(cube, cube2) {
			contacts++
		}
	}

	return contacts

}

func main() {

	execType := "test"
	cubes := array.Map(input.GetLines(2022, 18, execType+".txt"), func(str string) []int { return array.Map(strings.Split(str, ","), StrToInt) })

	cubesContacts := array.Map(cubes, func(cube []int) int { return contacts(cube, cubes) })

	fmt.Println(cubes)

	fmt.Println(cubesContacts)

	freeSides := array.Map(cubesContacts, func(contacts int) int { return 6 - contacts })
	fmt.Println(cubesContacts)
	fmt.Println(freeSides)

	total := 0
	for _, s := range freeSides {
		total += s
	}

	fmt.Println("Total", total)
}
