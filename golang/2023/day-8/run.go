package y2023d8

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

//go:embed test1.txt
var test1 string

//go:embed test0.txt
var test0 string

// This one is only usable for part 2
//
//go:embed test2.txt
var test2 string

func parseInput(text string) ([]rune, map[string][2]string) {
	lines := strings.Split(text, "\n\n")

	route := []rune(lines[0])

	nodes := map[string][2]string{}

	for _, line := range strings.Split(lines[1], "\n") {
		comps1 := strings.Split(line, " = ")

		root := comps1[0]

		cleanComp := strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(comps1[1], "(", ""), ")", ""), " ", "")
		children := strings.Split(cleanComp, ",")

		nodes[root] = [2]string{children[0], children[1]}
	}

	return route, nodes

}

func Run1() {

	origin := "AAA"
	destination := "ZZZ"
	route, nodes := parseInput(input)
	routeDefLength := len(route)
	fmt.Println(route, nodes)

	steps := 0
	currentNode := origin
	for {
		routeIdx := steps % routeDefLength
		nextDirection := route[routeIdx]

		if nextDirection == rune('L') {
			currentNode = nodes[currentNode][0]
		} else {
			currentNode = nodes[currentNode][1]
		}

		if currentNode == destination {
			break
		}
		steps++
	}

	fmt.Println("Steps:", steps+1)
}

func getOrigins(m map[string][2]string) []string {
	origins := []string{}

	for key, _ := range m {
		if key[2] == 'A' {
			origins = append(origins, key)
		}
	}

	return origins
}

func getDestinations(m map[string][2]string) []string {
	destinations := []string{}

	for key, _ := range m {
		if key[2] == 'Z' {
			destinations = append(destinations, key)
		}
	}

	return destinations
}

func Run2() {

	route, nodes := parseInput(input)
	routeDefLength := len(route)

	origins := getOrigins(nodes)
	fmt.Println(origins)
	destinations := getDestinations(nodes)
	fmt.Println(destinations)

	originsToDestinations := map[string]map[string]map[int]int{}
	for _, origin := range origins {
		originsToDestinations[origin] = map[string]map[int]int{}

		for _, destination := range destinations {
			originsToDestinations[origin][destination] = map[int]int{}
		}
	}

	for _, origin := range origins {
		currentNode := origin
		steps := 0
		for {

			// Loop until origin reaches itself again in index 0
			routeIdx := steps % routeDefLength
			nextDirection := route[routeIdx]

			if nextDirection == rune('L') {
				currentNode = nodes[currentNode][0]

			} else {
				currentNode = nodes[currentNode][1]
			}

			// This is a destination
			if currentNode[2] == 'Z' {
				// The same destination was already reached in this part of the route.
				// We are in a loop.
				// We do need to keep calculating
				if _, ok := originsToDestinations[origin][currentNode][routeIdx]; ok {
					break
				}
				originsToDestinations[origin][currentNode][routeIdx] = steps + 1
			}

			steps++

		}

	}

	numbers := []int{}
	total := 1
	for _, om := range originsToDestinations {
		for _, dm := range om {
			//indices := []string{}

			for _, steps := range dm {
				//indices = append(indices, fmt.Sprintf("%d = %d", i, steps))
				numbers = append(numbers, steps)
				total *= steps
			}

			// if len(indices) > 0 {
			// 	fmt.Println(o, "=>", d, "=", indices)
			// }
		}
	}

	//fmt.Println("numbers", numbers)
	//fmt.Printf("numbers2:%d, %d, %v\n", numbers[0], numbers[1], numbers[2:])
	//fmt.Println("total", total)
	fmt.Println("LCM", LCM(numbers[0], numbers[1], numbers[2:]...))

}

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}
