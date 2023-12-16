package y2023d15

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed test.txt
var test string

//go:embed input.txt
var input string

type lens struct {
	string
	int
}

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
	commands := strings.Split(test, ",")

	total := 0

	for _, command := range commands {
		hash := getHash(command)

		fmt.Println(command, "becomes", hash)
		total += hash
	}

	fmt.Println("Result", total)

}

func removeLabel(label string, lenses []lens) []lens {
	result := []lens{}

	for _, lens := range lenses {
		if lens.string != label {
			result = append(result, lens)
		}
	}

	return result
}

func findLabel(label string, lenses []lens) (int, bool) {
	for idx, lens := range lenses {
		if lens.string == label {
			return idx, true
		}
	}

	return 0, false
}

func Run2() {
	commands := strings.Split(input, ",")

	boxes := map[int][]lens{}

	for _, command := range commands {
		if strings.Contains(command, "=") {
			parts := strings.Split(command, "=")
			label := parts[0]
			value := parts[1]

			ivalue, err := strconv.Atoi(value)

			if err != nil {
				panic(err)
			}

			boxIdx := getHash(label)

			arr, ok := boxes[boxIdx]

			if !ok {
				arr = []lens{}
			}

			idx, found := findLabel(label, arr)

			if found {
				arr[idx] = lens{label, ivalue}
			} else {
				arr = append(arr, lens{label, ivalue})

			}

			boxes[boxIdx] = arr

			fmt.Println("BOX ", boxIdx, ":", arr)

		} else {
			parts := strings.Split(command, "-")
			label := parts[0]

			boxIdx := getHash(label)

			arr, ok := boxes[boxIdx]

			if !ok {
				arr = []lens{}
			}

			arr = removeLabel(label, arr)

			boxes[boxIdx] = arr

			fmt.Println("BOX ", boxIdx, ":", arr)
		}

	}

	sum := 0
	for i := 0; i < 256; i++ {
		for lensIdx, lens := range boxes[i] {
			focusingPower := (i + 1) * (lensIdx + 1) * lens.int
			fmt.Println(lens.string, ":", (i + 1), "box(", i, ") * ", (lensIdx + 1), "(slot) * ", lens.int, "(focal length) = ", focusingPower)
			sum += focusingPower
		}
	}

	fmt.Println("Result:", sum)
}
