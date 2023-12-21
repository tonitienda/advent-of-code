package y2023y21

import (
	_ "embed"
	"fmt"
	"slices"
	"strings"
)

//go:embed test.txt
var test string

//go:embed input.txt
var input string

func findStart(garden [][]rune) [2]int {

	for i, row := range garden {
		for j, c := range row {
			if c == 'S' {
				return [2]int{i, j}

			}
		}
	}

	return [2]int{}

}

var neighboursMap map[int]map[int][][2]int

func init() {
	neighboursMap = map[int]map[int][][2]int{}
}

func findReachableNeighboursInfinite(garden [][]rune, location [2]int) [][2]int {
	reachableNeighbours := [][2]int{}

	rows := len(garden)
	cols := len(garden[0])

	i := location[0]
	j := location[1]

	ni := (i % rows) + rows
	nj := (j % cols) + cols

	// TOP
	if garden[(ni-1)%rows][nj%cols] == '.' {
		reachableNeighbours = append(reachableNeighbours, [2]int{i - 1, j})
	}

	// BOTTOM
	if garden[(ni+1)%rows][nj%cols] == '.' {
		reachableNeighbours = append(reachableNeighbours, [2]int{i + 1, j})
	}

	// LEFT
	if garden[ni%rows][(nj-1)%cols] == '.' {
		reachableNeighbours = append(reachableNeighbours, [2]int{i, j - 1})
	}

	// RIGHT
	if garden[ni%rows][(nj+1)%cols] == '.' {
		reachableNeighbours = append(reachableNeighbours, [2]int{i, j + 1})
	}

	return reachableNeighbours
}

func findReachableNeighbours(garden [][]rune, location [2]int) [][2]int {
	reachableNeighbours := [][2]int{}

	i := location[0]
	j := location[1]

	// TOP
	if i > 0 && garden[i-1][j] == '.' {
		reachableNeighbours = append(reachableNeighbours, [2]int{i - 1, j})
	}

	// BOTTOM
	if i < len(garden)-1 && garden[i+1][j] == '.' {
		reachableNeighbours = append(reachableNeighbours, [2]int{i + 1, j})
	}

	// LEFT
	if j > 0 && garden[i][j-1] == '.' {
		reachableNeighbours = append(reachableNeighbours, [2]int{i, j - 1})
	}

	// RIGHT
	if j < len(garden[0])-1 && garden[i][j+1] == '.' {
		reachableNeighbours = append(reachableNeighbours, [2]int{i, j + 1})
	}

	return reachableNeighbours
}

func findReachableLocations(garden [][]rune, reached [][2]int) [][2]int {
	reachable := [][2]int{}
	for _, r := range reached {
		neighbours := findReachableNeighbours(garden, r)
		for _, n := range neighbours {
			if !slices.Contains(reachable, n) {
				reachable = append(reachable, n)
			}
		}
	}

	return reachable
}

func findReachableLocationsInfinite(garden [][]rune, reached [][2]int) [][2]int {
	reachable := [][2]int{}
	for _, r := range reached {
		neighbours := findReachableNeighboursInfinite(garden, r)
		for _, n := range neighbours {
			if !slices.Contains(reachable, n) {
				reachable = append(reachable, n)
			}
		}
	}

	return reachable
}

func Run1() {
	fmt.Println(input)
	numberOfSteps := 64

	garden := [][]rune{}

	for _, line := range strings.Split(input, "\n") {
		garden = append(garden, []rune(line))
	}

	start := findStart(garden)

	garden[start[0]][start[1]] = '.'
	reachedLocations := [][2]int{start}

	for i := 0; i < numberOfSteps; i++ {
		reachedLocations = findReachableLocations(garden, reachedLocations)
	}

	fmt.Println("reachedLocations: ", len(reachedLocations))

}

func Run2() {
	fmt.Println(test)
	numberOfSteps := 5000

	garden := [][]rune{}

	for _, line := range strings.Split(test, "\n") {
		garden = append(garden, []rune(line))
	}

	start := findStart(garden)

	garden[start[0]][start[1]] = '.'
	reachedLocations := [][2]int{start}

	for i := 0; i < numberOfSteps; i++ {
		reachedLocations = findReachableLocationsInfinite(garden, reachedLocations)
	}

	fmt.Println("reachedLocations: ", len(reachedLocations))

}
