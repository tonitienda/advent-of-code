package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

//go:embed test.txt
var test string

func main() {

	fmt.Println(input)

	data := strings.ReplaceAll(input, "      ", "")
	data = strings.ReplaceAll(data, "     ", "")
	data = strings.ReplaceAll(data, "    ", "")
	data = strings.ReplaceAll(data, "   ", "")
	data = strings.ReplaceAll(data, "  ", "")

	lines := strings.Split(data, "\n")
	times := strings.Split(strings.Trim(strings.ReplaceAll(lines[0], "Time:", ""), " "), " ")
	distances := strings.Split(strings.Trim(strings.ReplaceAll(lines[1], "Distance:", ""), " "), " ")

	fmt.Println(times)
	fmt.Println(distances)

	races := [][2]int{}

	for idx, time := range times {

		t, err := strconv.Atoi(time)

		if err != nil {
			panic(err)
		}

		d, err := strconv.Atoi(distances[idx])

		if err != nil {
			panic(err)
		}

		races = append(races, [2]int{t, d})

	}

	fmt.Println(races)

	winingStrategies := 1

	for _, race := range races {
		//fmt.Println("Time:", race[0])
		currentTimeWiningStrategy := 0
		for t := 0; t < race[0]; t++ {
			totalDistance := (race[0] - t) * t

			//fmt.Printf("%d -> %d\n", t, totalDistance)

			if totalDistance > race[1] {
				currentTimeWiningStrategy++
			}
		}

		winingStrategies *= currentTimeWiningStrategy
		currentTimeWiningStrategy = 0
	}

	fmt.Println("winingStrategies:", winingStrategies)

	// 7 ms
	// 0 -> 0
	// 1 -> 6
	// 2 -> 10
	// 3 -> 12
	// 4 -> 12
	// 5 -> 10
	// 6 -> 6
	// 7 -> 0

	// 15 ms
	// 0 -> 0
	// 1 -> 14
	// 2 -> 26
	// 3 ->

}
