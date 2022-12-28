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

/*

root: pppw + sjmn
dbpl: 5
-- cczh: sllz + lgvd
zczc: 2
-- ptdq: humn - dvpt
-- dvpt: 3
lfqf: 4
humn: 5
ljgn: 2
sjmn: drzm * dbpl
-- sllz: 4
pppw: cczh / lfqf
 -- lgvd: ljgn * ptdq
drzm: hmdt - zczc
hmdt: 32

humn = ptdq + dvpt
humn = (lgvd / ljgn) + 3
humn = (cczh - sllz) / 2) + 3
humn = ((pppw * lfqf) - 4)) / 2



*/

type Monkey struct {
	id        string
	value     int
	isSolved  bool
	m1        string
	operation string
	m2        string
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

func getInverseOperation(monkey Monkey) func(n1, n2 int) int {
	switch monkey.operation {
	case "+":
		return func(n1, n2 int) int { return n1 - n2 }
	case "-":
		return func(n1, n2 int) int { return n1 + n2 }
	case "*":
		return func(n1, n2 int) int { return n1 / n2 }
	case "/":
		return func(n1, n2 int) int { return n1 * n2 }

	default:
		panic("Operation " + monkey.operation + " not supported.")

	}
}

func solveEquation(value int, monkey Monkey, index map[string]Monkey) int {
	fmt.Println("Trying to solve ", monkey.id)

	if monkey.id == "humn" {
		fmt.Println(value)
		return value
	}

	if monkey.isSolved {
		return monkey.value
	}

	node1 := index[monkey.m1]
	node2 := index[monkey.m2]

	operation := getInverseOperation(monkey)

	if node1.isSolved {
		return solveEquation(operation(value, node1.value), node2, index)
	} else {
		return solveEquation(operation(value, node2.value), node1, index)
	}
}

func main() {

	execType := "test"
	data := input.GetLines(2022, 21, execType+".txt")

	monkeys := array.Map(data, getMonkey)

	index := makeMonkeyMap(monkeys)

	index = makeMonkeyMap(monkeys)
	rootMonkey := index["root"]
	rootMonkey.operation = "="
	index["root"] = rootMonkey

	humnMonkey := index["humn"]
	humnMonkey.isSolved = false
	index["humn"] = humnMonkey

	// index2 := map[string]func() int{}
	// index3 := map[string]bool{}

	// descriptions := map[string]string{}

	// currentHumnValue := 0
	somethingSolved := true

	// Solve everything solvable
	for somethingSolved {
		somethingSolved = false

		for key, monkey := range index {
			if key != "humn" && key != "root" && !monkey.isSolved {
				m1 := index[monkey.m1]
				m2 := index[monkey.m2]

				if m1.isSolved && m2.isSolved {
					switch monkey.operation {
					case "+":
						monkey.value = m1.value + m2.value
					case "-":
						monkey.value = m1.value - m2.value
					case "*":
						monkey.value = m1.value * m2.value
					case "/":
						monkey.value = m1.value / m2.value

					default:
						panic("Operation " + monkey.operation + " not supported.")

					}

					monkey.isSolved = true
					index[key] = monkey
					somethingSolved = true
				}
			}
		}
	}

	root := index["root"]

	node1 := index[root.m1]
	node2 := index[root.m2]

	if node1.isSolved {
		fmt.Println(solveEquation(node1.value, node2, index))
	} else {
		fmt.Println(solveEquation(node2.value, node1, index))
	}

}
