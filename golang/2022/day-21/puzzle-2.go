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
		return func(n1, n2 int) int { fmt.Printf("%d - %d\n", n2, n1); return n2 - n1 }
	case "-":
		return func(n1, n2 int) int { fmt.Printf("%d + %d\n", n1, n2); return n1 + n2 }
	case "*":
		return func(n1, n2 int) int { fmt.Printf("%d / %d\n", n1, n2); return n1 / n2 }
	case "/":
		return func(n1, n2 int) int { fmt.Printf("%d * %d\n", n1, n2); return n1 * n2 }

	default:
		panic("Operation " + monkey.operation + " not supported.")

	}
}

func getOperation(monkey Monkey) func(n1, n2 int) int {
	switch monkey.operation {
	case "+":
		return func(n1, n2 int) int { fmt.Printf("%d + %d\n", n1, n2); return n1 + n2 }
	case "-":
		return func(n1, n2 int) int { fmt.Printf("%d - %d\n", n1, n2); return n1 - n2 }
	case "*":
		return func(n1, n2 int) int { fmt.Printf("%d * %d\n", n1, n2); return n1 * n2 }
	case "/":
		return func(n1, n2 int) int { fmt.Printf("%d / %d\n", n1, n2); return n1 / n2 }

	default:
		panic("Operation " + monkey.operation + " not supported.")

	}
}

func main() {

	execType := "input"
	data := input.GetLines(2022, 21, execType+".txt")

	monkeys := array.Map(data, getMonkey)

	index := makeMonkeyMap(monkeys)

	index2 := map[string]func() int{}
	index3 := map[string]bool{}

	descriptions := map[string]string{}

	currentHumnValue := 0

	somethingSolved := true
	// Solve everything solvable
	for somethingSolved {
		somethingSolved = false

		for key, monkey := range index {
			if monkey.m1 != "humn" && key != "root" && !monkey.isSolved {
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

	for !index3["root"] {
		for key, monkey := range index {
			if !index3[key] {
				if key == "humn" {
					index2[key] = func(key string) func() int {
						return func() int { return currentHumnValue }
					}(key)
					index3[key] = true
					descriptions["humn"] = "humn"
				} else if monkey.isSolved {
					index2[key] = func(monkey Monkey) func() int { return func() int { return monkey.value } }(monkey)
					index3[key] = true

					descriptions[key] = strconv.Itoa(monkey.value)
				} else {
					// The monkey needs other monkeys data

					f1, ok1 := index2[monkey.m1]
					f2, ok2 := index2[monkey.m2]

					if ok1 && ok2 {

						switch monkey.operation {
						case "+":
							index2[key] = func(f1, f2 func() int, key string) func() int {
								return func() int { return f1() + f2() }
							}(f1, f2, key)
							descriptions[key] = "(" + descriptions[monkey.m1] + "+" + descriptions[monkey.m2] + ")"

						case "-":
							index2[key] = func(f1, f2 func() int, key string) func() int {
								return func() int { return f1() - f2() }
							}(f1, f2, key)
							descriptions[key] = "(" + descriptions[monkey.m1] + "-" + descriptions[monkey.m2] + ")"

						case "*":
							index2[key] = func(f1, f2 func() int, key string) func() int {
								return func() int { return f1() * f2() }
							}(f1, f2, key)
							descriptions[key] = descriptions[monkey.m1] + "*" + descriptions[monkey.m2]

						case "/":
							index2[key] = func(f1, f2 func() int, key string) func() int {
								return func() int { return f1() / f2() }
							}(f1, f2, key)
							descriptions[key] = descriptions[monkey.m1] + "/" + descriptions[monkey.m2]

						case "=":
							index2[key] = func(f1, f2 func() int) func() int {
								return func() int { return f1() - f2() }
							}(f1, f2)
							descriptions[key] = "(" + descriptions[monkey.m1] + "==" + descriptions[monkey.m2] + ")"

						}

						index3[key] = true
					}

				}
			}
		}
	}

	//root := index["root"]

	// node1 := index[root.m1]
	// node2 := index[root.m2]

	//result := solveEquation(0, root, index)
	// if node1.isSolved {
	// 	result = solveEquation(node1.value, node2, index)
	// } else {
	// 	result = solveEquation(node2.value, node1, index)
	// }

	// fmt.Println(result)
	// fmt.Println("\n\n\n\n\n")
	// f := index2["root"]

	// currentHumnValue = result
	// fmt.Println(f())

	// root := Monkey{id: "root", m1: "a", m2: "b", operation: "+"}
	// a := Monkey{id: "a", value: 24, isSolved: true}
	// b := Monkey{id: "b", m1: "c", m2: "d", operation: "+"}
	// c := Monkey{id: "c", value: 8, isSolved: true}
	// d := Monkey{id: "d", m1: "e", m2: "f", operation: "-"}
	// e := Monkey{id: "e", value: 20, isSolved: true}
	// f := Monkey{id: "f", m1: "humn", m2: "g", operation: "*"}
	// humn := Monkey{id: "humn", value: 0, isSolved: false}
	// g := Monkey{id: "g", value: 2, isSolved: true}

	// indextest := map[string]Monkey{
	// 	"root": root,
	// 	"a":    a,
	// 	"b":    b,
	// 	"c":    c,
	// 	"d":    d,
	// 	"e":    e,
	// 	"f":    f,
	// 	"g":    g,
	// 	"humn": humn,
	// }
	fmt.Println(descriptions["root"])
	humn := index["humn"]
	humn.isSolved = false
	index["humn"] = humn
	root := index["root"]
	m1 := index[root.m1]
	m2 := index[root.m2]
	fmt.Println(descriptions[root.m1])
	fmt.Println(descriptions[root.m2])

	if m1.isSolved {
		solve(m1.value, m2, index)
	} else if m2.isSolved {
		solve(m2.value, m1, index)
	}

}

func getNextLevelSolvedMonkey(m Monkey, index map[string]Monkey) Monkey {
	m1 := index[m.m1]
	m2 := index[m.m2]

	if m1.isSolved {
		return m1
	}

	if m2.isSolved {
		return m2
	}

	panic("None of the children of " + m.id + " are solved.")
}

func solve(value int, m Monkey, index map[string]Monkey) int {
	if m.id == "humn" {
		fmt.Println("HUMN: Value", value)
		return value
	}

	// m = 24 + x
	m1 := index[m.m1] // 24
	m2 := index[m.m2] // 8 + y
	//m21 := 8
	// m22 := y

	fmt.Println("value", value, "id", m.id, "[", m.m1, ",", m.m2, "]")
	if m1.isSolved {
		fmt.Println(m1.id, "is solved. v:", m1.value, "m op:", m.operation)
		operation := getInverseOperation(m)
		if m.operation == "-" || m.operation == "/" {
			operation = getOperation(m)
		}

		if m.operation == "*" {
			value = operation(value, m1.value)
		} else {
			value = operation(m1.value, value)
		}
		fmt.Println("New Value", value)
		return solve(value, m2, index)
	}

	if m2.isSolved {
		fmt.Println(m2.id, "is solved. v:", m2.value)
		operation := getInverseOperation(m)
		value = operation(value, m2.value)

		fmt.Println("New Value", value)
		return solve(value, m1, index)
	}

	return 0

}
