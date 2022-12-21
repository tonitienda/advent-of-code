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

func add(mId1, mId2 string, index map[string]Monkey) int {
	return index[mId1].calculate(index) + index[mId2].calculate(index)
}

func subract(mId1, mId2 string, index map[string]Monkey) int {
	return index[mId1].calculate(index) - index[mId2].calculate(index)
}

func divide(mId1, mId2 string, index map[string]Monkey) int {
	return index[mId1].calculate(index) / index[mId2].calculate(index)
}

func multiply(mId1, mId2 string, index map[string]Monkey) int {
	return index[mId1].calculate(index) * index[mId2].calculate(index)
}

type Monkey struct {
	id        string
	value     int
	isSolved  bool
	m1        string
	operation string
	m2        string
}

func getMonkey(line string) Monkey {

	fmt.Println(line)

	data := strings.Split(line, ":")

	fmt.Println(data)
	monkeyId := data[0]

	data2 := strings.Split(strings.Trim(data[1], " "), " ")

	value := 0
	isSolved := false

	m1 := ""
	operation := ""
	m2 := ""
	// var calculate func(map[string]Monkey) (int, bool)

	if len(data2) == 1 {
		value = StrToInt(data2[0])
		isSolved = true
	} else {
		m1 = data2[0]
		operation = data2[1]
		m2 = data2[2]
	}

	return Monkey{
		id:        monkeyId,
		isSolved:  isSolved,
		value:     value,
		m1:        m1,
		operation: operation,
		m2:        m2,
	}
}

func (monkey Monkey) calculate(index map[string]Monkey) int {

	if monkey.isSolved {
		return monkey.value
	}

	switch monkey.operation {
	case "+":
		return add(monkey.m1, monkey.m2, index)
	case "-":
		return subract(monkey.m1, monkey.m2, index)
	case "*":
		return multiply(monkey.m1, monkey.m2, index)
	case "/":
		return divide(monkey.m1, monkey.m2, index)
	}

	panic("Not sure what to do")
}

func makeMonkeyMap(monkeys []Monkey) map[string]Monkey {
	index := map[string]Monkey{}

	for _, monkey := range monkeys {
		index[monkey.id] = monkey
	}

	return index
}

func getMonkeyValue(monkeyId string, index map[string]Monkey) int {

	monkey := index["root"]

	return monkey.calculate(index)

}

func main() {

	execType := "input"
	data := input.GetLines(2022, 21, execType+".txt")

	monkeys := array.Map(data, getMonkey)

	//fmt.Println(monkeys)

	index := makeMonkeyMap(monkeys)

	fmt.Println(getMonkeyValue("root", index))

}
