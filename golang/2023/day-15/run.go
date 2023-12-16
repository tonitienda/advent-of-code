package y2023d15

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed test.txt
var test string

//go:embed input.txt
var input string

func getHash(val string) int {
	hash := 0
	for _, r := range val {
		hash += int(r)
		hash *= 17
		hash %= 256
	}

	return hash
}

func Run1() {
	commands := strings.Split(input, ",")

	total := 0

	for _, command := range commands {
		hash := getHash(command)

		fmt.Println(command, "becomes", hash)
		total += hash
	}

	fmt.Println("Result", total)

}
