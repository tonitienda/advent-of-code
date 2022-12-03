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

func main() {
	data := input.GetContents(2022, 1, "input.txt")

	elvesFoods := strings.Split(data, "\n\n")
	elvesFoods2 := array.Map(elvesFoods, func(s string) []int { return array.Map(strings.Split(s, "\n"), StrToInt) })
	elvesFoodsTotals := array.Map(elvesFoods2, func(nums []int) int { return array.Sum(nums) })

	fmt.Println(array.Max(elvesFoodsTotals))

}
