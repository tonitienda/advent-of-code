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

func ls(currentDir []string, commands []string, pc int, fileSizes map[string]int) (map[string]int, int) {
	numLines := len(commands)
	pc++

	for pc < numLines {
		components := strings.Split(commands[pc], " ")

		// When we reach the next command we return
		if components[0] == "$" {
			return fileSizes, pc
		}

		// When we need to ignore the directories. We only want to cound
		// the file sizes the next command we return
		if components[0] != "dir" {
			size := StrToInt(components[0])
			depth := len(currentDir)
			i := 0
			for i < depth {
				key := strings.Join(currentDir[:depth-i], "_")
				fileSizes[key] += size
				i++
			}
		}
		pc++
	}

	return fileSizes, pc
}

func cd(currentDir []string, dest string, pc int) ([]string, int) {

	switch dest {
	case "/":
		currentDir = []string{"/"}
	case "..":
		currentDir = currentDir[:len(currentDir)-1]
	default:
		currentDir = append(currentDir, dest)
	}

	return currentDir, pc + 1
}

func calculateFolderSizes(commands []string) map[string]int {
	fileSizes := map[string]int{}
	currentDir := []string{"/"}

	pc := 0
	numLines := len(commands)

	for pc < numLines {
		comps := strings.Split(commands[pc], " ")

		switch comps[0] {
		case "$":
			switch comps[1] {
			case "ls":
				fileSizes, pc = ls(currentDir, commands, pc, fileSizes)
			case "cd":
				currentDir, pc = cd(currentDir, comps[2], pc)

			}
		}
	}
	return fileSizes
}

func calculateResult(fileSizes map[string]int) int32 {
	var total int32 = 0
	for _, value := range fileSizes {
		if value <= 100000 {
			total += int32(value)
		}
	}
	return total
}

func main() {
	commands := input.GetLines(2022, 7, "input.txt")

	fileSizes := calculateFolderSizes(commands)

	fmt.Println(calculateResult(fileSizes))

}
