package main

import (
	"advent/utils/array"
	"advent/utils/funct"
	"advent/utils/input"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func StrToInt(str string) int {
	return funct.GetValue(strconv.Atoi(str))
}

type MonkeyData struct {
	currentItems        []int
	operation           func(int) int
	getNextMonkey       func(int) int
	totalInspectedItems int
}

func genValue(old int, arg string) int {

	if arg == "old" {
		return old
	}
	return StrToInt(arg)
}

func genOperation(args []string) func(int) int {
	arg1 := strings.Trim(args[0], " ")
	op := strings.Trim(args[1], " ")
	arg2 := strings.Trim(args[2], " ")

	return func(old int) int {
		var1 := genValue(old, arg1)
		var2 := genValue(old, arg2)
		switch op {
		case "+":
			return var1 + var2
		case "-":
			return var1 - var2
		case "*":
			return var1 * var2
		case "/":
			return var1 / var2
		}

		panic("Operation not supported!")
	}
}

func deleteFirst(array []int) []int {
	newArray := []int{}

	for _, item := range array[1:] {
		newArray = append(newArray, item)
	}

	return newArray
}

func main() {
	monkeysData := strings.Split(input.GetContents(2022, 11, "input.txt"), "\n\n")
	monkeys := []MonkeyData{}

	for _, monkeyData := range monkeysData {
		parts := strings.Split(monkeyData, "\n")
		startItems := array.Map(strings.Split(strings.ReplaceAll(parts[1], ",", ""), " ")[4:], StrToInt)
		operation := genOperation(strings.Split(strings.ReplaceAll(parts[2], "  Operation: new = ", ""), " "))

		test := StrToInt(strings.Trim(strings.ReplaceAll(parts[3], "  Test: divisible by ", ""), " "))
		trueMonkey := StrToInt(strings.Trim(strings.ReplaceAll(parts[4], "    If true: throw to monkey ", ""), " "))
		falseMonkey := StrToInt(strings.Trim(strings.ReplaceAll(parts[5], "    If false: throw to monkey ", ""), " "))

		monkey := MonkeyData{
			currentItems: startItems,
			operation:    operation,
			getNextMonkey: func(value int) int {
				if value%test == 0 {
					return trueMonkey
				} else {
					return falseMonkey
				}
			},
		}

		monkeys = append(monkeys, monkey)
	}

	for times := 0; times < 20; times++ {
		for i := 0; i < len(monkeys); i++ {
			monkey := monkeys[i]

			fmt.Printf("Monkey %d:\n", i)
			currentItems := monkey.currentItems

			for _, item := range currentItems {

				fmt.Printf("\tMonkey inspects an item with a worry level of %d.\n", item)
				newWorry := monkey.operation(item)
				fmt.Printf("\t\tWorry level is operated to %d.\n", newWorry)
				boredWorry := newWorry / 3
				fmt.Printf("\t\tMonkey gets bored with item. Worry level is divided by 3 to %d.\n", boredWorry)

				nextMonkey := monkey.getNextMonkey(boredWorry)
				fmt.Printf("\t\tItem with worry level %d is thrown to monkey %d.\n", boredWorry, nextMonkey)

				monkeys[nextMonkey].currentItems = append(monkeys[nextMonkey].currentItems, boredWorry)
			}

			monkeys[i].totalInspectedItems += len(monkeys[i].currentItems)
			monkeys[i].currentItems = []int{}
		}
	}

	fmt.Println()
	for i, monkey := range monkeys {
		fmt.Println("Monkey", i, ":", monkey.totalInspectedItems)
	}
	fmt.Println()
	var totalInspectedItems []int = array.Map(monkeys, func(m MonkeyData) int { return m.totalInspectedItems })
	sort.Ints(totalInspectedItems)

	fmt.Println(totalInspectedItems[len(totalInspectedItems)-1] * totalInspectedItems[len(totalInspectedItems)-2])
}
