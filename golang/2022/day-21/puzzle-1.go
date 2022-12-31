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

type Monkey struct {
	id        string
	value     int
	isSolved  bool
	m1        string
	operation string
	m2        string
	m1r       *Monkey
	m2r       *Monkey
}

func getMonkey(line string) Monkey {

	data := strings.Split(line, ":")

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

func makeMonkeyMap(monkeys []Monkey) map[string]Monkey {
	index := map[string]Monkey{}

	for _, monkey := range monkeys {
		index[monkey.id] = monkey
	}

	return index
}

func main() {

	execType := "test"
	data := input.GetLines(2022, 21, execType+".txt")

	monkeys := array.Map(data, getMonkey)
	index := makeMonkeyMap(monkeys)

	index2 := map[string]func() int{}
	index3 := map[string]bool{}

	for !index3["root"] {
		for key, monkey := range index {
			if !index3[key] {
				if monkey.isSolved {
					index2[key] = func(monkey Monkey) func() int { return func() int { return monkey.value } }(monkey)
					index3[key] = true
				} else {
					// The monkey needs other monkeys data
					f1, ok1 := index2[monkey.m1]
					f2, ok2 := index2[monkey.m2]

					if ok1 && ok2 {
						switch monkey.operation {
						case "+":
							index2[key] = func(f1, f2 func() int, m1, m2 string) func() int {
								return func() int { return f1() + f2() }
							}(f1, f2, monkey.m1, monkey.m2)
						case "-":
							index2[key] = func(f1, f2 func() int, m1, m2 string) func() int {
								return func() int { return f1() - f2() }
							}(f1, f2, monkey.m1, monkey.m2)
						case "*":
							index2[key] = func(f1, f2 func() int, m1, m2 string) func() int {
								return func() int { return f1() * f2() }
							}(f1, f2, monkey.m1, monkey.m2)
						case "/":
							index2[key] = func(f1, f2 func() int, m1, m2 string) func() int {
								return func() int { return f1() / f2() }
							}(f1, f2, monkey.m1, monkey.m2)
						case "=":
							index2[key] = func(f1, f2 func() int) func() int {
								return func() int { return f1() - f2() }
							}(f1, f2)
						}

						index3[key] = true
					}
				}
			}
		}

	}
	f := index2["root"]
	fmt.Println(f())

}
