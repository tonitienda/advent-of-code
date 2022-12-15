package main

import (
	"advent/utils/funct"
	"advent/utils/input"
	"fmt"
	"strconv"
	"strings"
)

func StrToInt(str string) int {
	return funct.GetValue(strconv.Atoi(str))
}

func parseLine(str string) [][]int {
	numbers := [][]int{}
	idx := 0

	tokens := strings.Split(str, ",")

	//fmt.Println(str, tokens)

	for _, token := range tokens {
		//fmt.Println(token)
		if idx > len(numbers)-1 {
			numbers = append(numbers, []int{})
		}

		if strings.Index(token, "[") == 0 {
			if len(numbers[idx]) > 0 {
				idx++
				if idx > len(numbers)-1 {
					numbers = append(numbers, []int{})
				}
			}
			num := strings.ReplaceAll(strings.ReplaceAll(token, "[", ""), "]", "")
			if num != "" {
				numbers[idx] = append(numbers[idx], StrToInt(num))
			}

			if strings.Index(token, "]") > -1 {
				idx++
			}
		} else if strings.Index(token, "]") > -1 {
			num := strings.ReplaceAll(token, "]", "")
			//fmt.Println("token", token, "num", num)
			numbers[idx] = append(numbers[idx], StrToInt(num))
			idx++
		} else if token != "" {
			//fmt.Println("token2", token)
			numbers[idx] = append(numbers[idx], StrToInt(token))
		}
	}

	return numbers
}

func areCorrectOrder(str1, str2 string) bool {

	contents1 := parseLine(str1)
	contents2 := parseLine(str2)

	// fmt.Println((contents1))
	// fmt.Println((contents2))
	// fmt.Println()
	fmt.Println("Compare", str1, "vs", str2)
	fmt.Println("Compare", contents1, "vs", contents2)

	if len(contents1) > len(contents2) {
		return false
	}
	for idx, list := range contents1 {
		if len(list) == 0 {
			fmt.Println("Left side ran out of items, so inputs are in the right order")
			return true
		}
		for idx2, item := range list {
			//fmt.Printf("[%d][%d] => %d == %d\n", idx, idx2, item, contents2[idx][idx2])

			if len(contents2[idx]) <= idx2 {
				fmt.Println("\tRight side ran out of items, so inputs are not in the right order")
				return false
			}

			fmt.Println("\tCompare", item, "vs", contents2[idx][idx2])
			if item < contents2[idx][idx2] {
				fmt.Println("\tLeft side is smaller, so inputs are in the right order")

				return true
			}

			if item > contents2[idx][idx2] {
				fmt.Println("\tRight side is smaller, so inputs are not in the right order")

				return false
			}

		}

	}

	return true
}

func main() {

	list := input.Get2DMatrix(input.GetContents(2022, 13, "test.txt"), "\n\n", "\n", func(str string) string { return str })

	rightOrder := []int{}
	for idx, pair := range list {
		if areCorrectOrder(pair[0], pair[1]) {
			rightOrder = append(rightOrder, idx+1)
		}
		fmt.Println()
	}

	fmt.Println(rightOrder)

	sum := 0
	for _, num := range rightOrder {
		sum += num
	}

	fmt.Println(sum)
}
