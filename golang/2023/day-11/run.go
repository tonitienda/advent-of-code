package y2023d11

import (
	_ "embed"
	"fmt"
	"math"
	"strings"
)

//go:embed test.txt
var test string

//go:embed input.txt
var input string

func findGalaxies(text string) ([][2]int, int, int) {
	galaxies := [][2]int{}
	rows := strings.Split(text, "\n")
	columns := strings.Split(rows[0], "")

	for i, row := range rows {
		for j, val := range []rune(row) {
			if val == '#' {
				galaxies = append(galaxies, [2]int{i, j})
			}
		}
	}

	return galaxies, len(rows), len(columns)

}

func anyGalaxyInColumn(galaxies [][2]int, column int) bool {
	for _, galaxy := range galaxies {
		if galaxy[1] == column {
			return true
		}
	}

	return false
}

func anyGalaxyInRow(galaxies [][2]int, row int) bool {
	for _, galaxy := range galaxies {
		if galaxy[0] == row {
			return true
		}
	}

	return false
}

func expandUniverse(galaxies [][2]int, maxHeight, maxWidth int, units int) ([][2]int, int, int) {
	// Expand rows
	widthExpansion := map[int]int{}

	expanseW := 0
	for j := 0; j < maxWidth; j++ {
		widthExpansion[j] = expanseW
		if !anyGalaxyInColumn(galaxies, j) {
			expanseW = expanseW + units - 1
		}
	}

	heightExpansion := map[int]int{}

	expanseH := 0
	for i := 0; i < maxHeight; i++ {
		heightExpansion[i] = expanseH
		if !anyGalaxyInRow(galaxies, i) {
			expanseH = expanseH + units - 1
		}
	}

	expandedGalaxies := [][2]int{}

	for _, galaxy := range galaxies {
		expandedGalaxies = append(expandedGalaxies, [2]int{
			galaxy[0] + heightExpansion[galaxy[0]],
			galaxy[1] + widthExpansion[galaxy[1]],
		})
	}

	return expandedGalaxies, maxHeight + heightExpansion[maxHeight-1], maxWidth + widthExpansion[maxWidth-1]
}

func calculateSquareDistance(g1, g2 [2]int) int64 {

	return int64(math.Max(float64(g1[0]), float64(g2[0])) - math.Min(float64(g1[0]), float64(g2[0])) +
		math.Max(float64(g1[1]), float64(g2[1])) - math.Min(float64(g1[1]), float64(g2[1])))
}

func Run1() {

	fmt.Println(input)

	galaxies, maxHeight, maxWidth := findGalaxies(input)

	expandedGalaxies, expandedHeight, expandedWidth := expandUniverse(galaxies, maxHeight, maxWidth, 1)

	fmt.Println(galaxies, maxHeight, maxWidth)

	fmt.Println(expandedGalaxies, expandedHeight, expandedWidth)

	totalDistance := int64(0)

	totalGalaxies := len(expandedGalaxies)

	fmt.Println("totalGalaxies", totalGalaxies)
	for idx1, g1 := range expandedGalaxies {
		for idx2, g2 := range expandedGalaxies {
			if idx1 != idx2 {
				distance := calculateSquareDistance(g1, g2)

				fmt.Printf("Distance %d - %d = %d\n", idx1+1, idx2+1, distance)

				totalDistance += distance
			}

		}
	}

	fmt.Println("Total distance", totalDistance/2)

}

func Run2() {

	galaxies, maxHeight, maxWidth := findGalaxies(input)

	expandedGalaxies, expandedHeight, expandedWidth := expandUniverse(galaxies, maxHeight, maxWidth, 1000000)

	fmt.Println(galaxies, maxHeight, maxWidth)

	fmt.Println(expandedGalaxies, expandedHeight, expandedWidth)

	totalDistance := int64(0)

	totalGalaxies := len(expandedGalaxies)

	fmt.Println("totalGalaxies", totalGalaxies)
	for idx1, g1 := range expandedGalaxies {
		for idx2, g2 := range expandedGalaxies {
			if idx1 != idx2 {
				distance := calculateSquareDistance(g1, g2)

				fmt.Printf("Distance %d - %d = %d\n", idx1+1, idx2+1, distance)

				totalDistance += distance
			}

		}
	}

	fmt.Println("Total distance", totalDistance/2)

}
