package main

import (
	"advent/utils/input"
	"fmt"
	"strings"
)

func areUnique(str string, index, numchars int) bool {
	part := str[index : numchars+index]

	//fmt.Println(index, numchars)
	//fmt.Println(part)
	for i, c := range strings.Split(part, "") {
		if strings.LastIndex(part, string(c)) != i {
			return false
		}
	}

	return true
}

func main() {
	data := input.GetContents(2022, 6, "input.txt")

	NUM_CHARS := 14
	fmt.Println(data)

	for i := 0; i < len(data)-(NUM_CHARS-1); i++ {
		if areUnique(data, i, NUM_CHARS) {
			fmt.Println(i + NUM_CHARS)
			return
		}
	}
}
