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

func snafuToDec(str string) int {
	comps := strings.Split(str, "")

	number := 0
	mult := 1
	for i := len(comps) - 1; i >= 0; i-- {
		if comps[i] == "=" {
			number += mult * -2
		} else if comps[i] == "-" {
			number += mult * -1
		} else {
			number += StrToInt(comps[i]) * mult
		}

		mult *= 5
	}

	return number
}

func decToSnafu(num int) string {
	base := 5
	negs := []string{"-", "="}
	str := ""

	if num == 0 {
		return "0"
	}

	for num > 0 {
		mod := num % base
		num = num / base

		idx := base - mod

		if idx <= len(negs) {
			str = negs[idx-1] + str
			num++
		} else {
			str = strconv.Itoa(mod) + str
		}
	}

	return str

}

func main() {
	execType := "input"
	data := array.Map(input.GetLines(2022, 25, execType+".txt"), snafuToDec)

	fmt.Println(data)
	// Tests
	fmt.Println(snafuToDec("1=-0-2"), "==", 1747)
	fmt.Println(snafuToDec("12111"), "==", 906)
	fmt.Println(snafuToDec("2=0="), "==", 198)
	fmt.Println(snafuToDec("21"), "==", 11)
	fmt.Println(snafuToDec("2=01"), "==", 201)
	fmt.Println(snafuToDec("111"), "==", 31)
	fmt.Println(snafuToDec("20012"), "==", 1257)
	fmt.Println(snafuToDec("112"), "==", 32)
	fmt.Println(snafuToDec("1=-1="), "==", 353)
	fmt.Println(snafuToDec("1-12"), "==", 107)
	fmt.Println(snafuToDec("12"), "==", 7)
	fmt.Println(snafuToDec("1="), "==", 3, "==", decToSnafu(3))
	fmt.Println(snafuToDec("122"), "==", 37, "==", decToSnafu(37))

	sum := 0

	for _, num := range data {
		sum += num
	}

	fmt.Println(decToSnafu(sum))

}
