package main

import (
	"advent/utils/array"
	"advent/utils/funct"
	"advent/utils/input"
	"fmt"
	"math"
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

	index = makeMonkeyMap(monkeys)
	rootMonkey := index["root"]
	rootMonkey.operation = "="
	index["root"] = rootMonkey

	index2 := map[string]func() int{}
	index3 := map[string]bool{}

	descriptions := map[string]string{}

	currentHumnValue := 0

	for !index3["root"] {
		for key, monkey := range index {
			if !index3[key] {
				if key == "humn" {
					index2[key] = func(key string) func() int {
						return func() int { return currentHumnValue }
					}(key)
					index3[key] = true
					descriptions["humn"] = "currentHumnValue"
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
							descriptions[key] = "(" + descriptions[monkey.m1] + "*" + descriptions[monkey.m2] + ")"

						case "/":
							index2[key] = func(f1, f2 func() int, key string) func() int {
								return func() int { return f1() / f2() }
							}(f1, f2, key)
							descriptions[key] = "(" + descriptions[monkey.m1] + "/" + descriptions[monkey.m2] + ")"

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

	fmt.Println("\n\n\n\n\n")
	f := index2["root"]
	// currentHumnValue = -305
	// fmt.Println(f())

	// fmt.Println(descriptions["root"])

	for currentHumnValue = 0; currentHumnValue < math.MaxInt; currentHumnValue++ {
		if currentHumnValue%1000000 == 0 {
			fmt.Println("Trying with", currentHumnValue)
		}

		if f() == 0 {
			fmt.Println("Humn Value", currentHumnValue)
			break
		}
	}

	secondsPerMillion := 2
	cpus := 8

	totalSeconds := (math.MaxInt / cpus / 1000000) * secondsPerMillion

	days := totalSeconds / 3600 / 24

	years := days / 365
	fmt.Println(days, "days")
	fmt.Println(years, "years")

	// fmt.Println("root", root)
	// fmt.Println("humn", humn)
	// root.operation = "="

	// fmt.Println("root", root)

	// fmt.Println("pppw", pppw)

	// //	fmt.Println("!!===>", monkey.id, monkey1, monkey2)

	// for i := 0; i < math.MaxInt; i++ {
	// 	//	if i%10000 == 0 {
	// 	fmt.Println("Trying with", i)
	// 	//	}

	// 	humn.value = i

	// 	//fmt.Println("HUMN value", humn.value)
	// 	fmt.Println("Here")
	// 	rootValue := root.calculate()

	// 	// Equality check
	// 	if rootValue == 0 {
	// 		fmt.Println("HUMN value was", humn.value)
	// 		return
	// 	}
	// }

}
