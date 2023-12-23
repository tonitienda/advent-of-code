package y2023d22

import (
	_ "embed"
	"fmt"
	"slices"
	"sort"
	"strconv"
	"strings"
)

//go:embed test.txt
var test string

//go:embed input.txt
var input string

type coords [3]int

type brick [2]coords

func getCoords(str string) coords {
	parts := strings.Split(str, ",")

	c1, err := strconv.Atoi(parts[0])

	if err != nil {
		panic(err)
	}

	c2, err := strconv.Atoi(parts[1])

	if err != nil {
		panic(err)
	}

	c3, err := strconv.Atoi(parts[2])

	if err != nil {
		panic(err)
	}

	return coords{c1, c2, c3}

}

func getBricks(text string) []brick {
	bricks := []brick{}

	for _, line := range strings.Split(text, "\n") {
		brickStr := strings.Split(line, "~")

		bricks = append(bricks, brick{
			getCoords(brickStr[0]), getCoords(brickStr[1]),
		})

	}

	return bricks
}

func decompose(b brick) (int, int, int, int, int, int) {
	return b[0][0], b[0][1], b[0][2], b[1][0], b[1][1], b[1][2]
}

func isSupported(b brick, bricks []brick) bool {

	// On the floor
	if b[0][2] == 1 {
		return true
	}

	// See if another brick supports it
	for _, b2 := range bricks {
		if supports(b2, b) {
			return true
		}
	}

	return false
}

func supports(brick1, brick2 brick) bool {
	// if Upper Z of brick1 is not in contact with loweZ of brick 2
	// they are not in contact
	x11, y11, _, x12, y12, z12 := decompose(brick1)
	x21, y21, z21, x22, y22, _ := decompose(brick2)

	//fmt.Println("Evaluating ", brick1, "and", brick2)

	//	fmt.Println("z12", z12, "z21 -1", z21-1)

	if z12 != z21-1 {
		return false
	}
	//fmt.Println("zs are correct")

	if x12 >= x21 && x11 <= x22 &&
		y12 >= y21 && y11 <= y22 {
		//fmt.Println("brick1 supports brick2")

		return true
	}

	return false
}

func getSupportAndSupportedBy(text string) ([]brick, map[int][]int, map[int][]int) {

	bricks := getBricks(text)

	sort.Slice(bricks, func(i, j int) bool {
		return bricks[i][0][2] < bricks[j][0][2]
	})

	restedBricks := []brick{}

	for _, b := range bricks {
		for !isSupported(b, restedBricks) {
			// Fall one cell
			b[0][2]--
			b[1][2]--
		}
		restedBricks = append(restedBricks, b)
	}

	support := map[int][]int{}
	isSupportedBy := map[int][]int{}

	for i := 0; i < len(restedBricks); i++ {
		for j := 0; j < len(restedBricks); j++ {
			if i != j {
				//fmt.Println(string('A'+rune(i)), "supports", string('A'+rune(j)), "?")

				if supports(restedBricks[i], restedBricks[j]) {
					//	fmt.Println("\tYes")

					support[i] = append(support[i], j)
					isSupportedBy[j] = append(isSupportedBy[j], i)
				} else {
					//	fmt.Println("\tNo")
				}

			}

		}
	}

	return bricks, support, isSupportedBy

}

func Run1() {

	bricks, support, supportedBy := getSupportAndSupportedBy(input)
	fmt.Println(support)

	fmt.Println(supportedBy)

	cannotBeRemoved := map[int]bool{}

	for _, v := range supportedBy {
		// If the brick is supported by only one,
		// that one cannot be removed
		if len(v) == 1 {
			cannotBeRemoved[v[0]] = true
		}
	}

	fmt.Println("Cannot be removed:", len(cannotBeRemoved))

	// for k, _ := range cannotBeRemoved {
	// 	fmt.Println(string('A' + rune(k)))
	// }

	fmt.Println("Can be removed:", len(bricks)-len(cannotBeRemoved))

}

func allRemoved(removed []int, existing []int) bool {
	for _, e := range existing {
		if !slices.Contains(removed, e) {
			return false
		}
	}

	return true
}

func Run2() {

	bricks, support, supportedBy := getSupportAndSupportedBy(input)
	fmt.Println(support)

	fmt.Println(supportedBy)

	fallingBricks := map[int][]int{}

	for idx := range bricks {
		// If I remove a brick, how many get without support
		removedBricks := []int{idx}
		brickWasRemoved := true

		for brickWasRemoved {
			brickWasRemoved = false
			for key, value := range supportedBy {
				if slices.Index(removedBricks, key) == -1 && allRemoved(removedBricks, value) {
					removedBricks = append(removedBricks, key)
					fallingBricks[idx] = append(fallingBricks[idx], key)
					brickWasRemoved = true
				}
			}
		}

	}

	fmt.Println("fallingBricks", fallingBricks)

	total := 0

	for _, v := range fallingBricks {
		total += len(v)
	}

	fmt.Println("total:", total)
	// for k, _ := range cannotBeRemoved {
	// 	fmt.Println(string('A' + rune(k)))
	// }

}

/*

A -> [B, C]
B -> [D, E]
C -> [D, E]
D -> [F]
E -> [F]
F -> [G]


A -> []
B -> [A]
C -> [A]
D -> [B,C]
E -> [B,C]
F -> [D,E]
G -> [F]

*/
