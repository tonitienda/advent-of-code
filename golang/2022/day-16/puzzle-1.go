package main

import (
	"advent/utils/array"
	"advent/utils/funct"
	"advent/utils/input"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func StrToInt(str string) int {
	return funct.GetValue(strconv.Atoi(str))
}

type Valve struct {
	id       string
	rate     int
	children []string
	isOpen   bool
}

func (v Valve) String() string {
	return fmt.Sprintf("%s (%d) -> %s", v.id, v.rate, v.children)

}

func getValve(line string) Valve {
	// Valve AA has flow rate=0; tunnels lead to valves DD, II, BB
	cleanStr := strings.ReplaceAll(line, "Valve ", "")
	cleanStr = strings.ReplaceAll(cleanStr, " has flow rate=", " ")

	cleanStr = strings.ReplaceAll(cleanStr, "; tunnels lead to valves ", " ")
	cleanStr = strings.ReplaceAll(cleanStr, "; tunnel leads to valve ", " ")

	cleanStr = strings.ReplaceAll(cleanStr, ", ", ",")

	parts := strings.Split(cleanStr, " ")

	id := parts[0]
	rate := parts[1]
	connections := strings.Split(parts[2], ",")

	return Valve{
		id:       id,
		rate:     StrToInt(rate),
		children: connections,
	}
}

func tabs(num int) string {
	tabs := ""

	for i := 0; i < num; i++ {
		tabs += "  "
	}

	return tabs
}

func calculateMinDistance(valve1, valve2 string, index map[string]Valve, cost int, path []string) (int, []string, bool) {
	//fmt.Println(valve1, "==", valve2)
	if valve1 == valve2 {
		return cost, path, true
	}

	valve := index[valve1]

	//fmt.Println(tabs(len(path)), valve1, "=>", valve2, "children", valve.children, ":", path)

	if includes(valve.children, valve2) {
		return cost + 1, append(path, valve2), true
	}

	minCost := int(math.Inf(1))
	minPath := []string{}

	atLeastOnePathFound := false
	for _, child := range valve.children {

		// Avoid cycles
		if includes(path, child) {
			continue
		}

		cost, newPath, ok := calculateMinDistance(child, valve2, index, cost+1, append(path, child))
		atLeastOnePathFound = atLeastOnePathFound || ok

		if ok {
			if cost < minCost {
				minCost = cost
				minPath = newPath
			}
		}
	}

	return minCost, minPath, atLeastOnePathFound
}

func makeValveMap(valves []Valve) map[string]Valve {
	index := map[string]Valve{}

	for _, node := range valves {
		index[node.id] = node
	}

	return index
}

func includes(arr []string, val string) bool {
	for _, v := range arr {
		if v == val {
			return true
		}
	}

	return false
}

func printValves(valves []Valve) {
	for _, valve := range valves {
		fmt.Println(valve)
	}
}

func printDistances(distances map[string][]Distance) {
	for key, dists := range distances {
		for _, dist := range dists {
			fmt.Println(key, "=>", dist.dest, dist.path)
		}
	}
}

type Distance struct {
	dest        string
	cost        int
	maxPressure int
	path        []string
}

func allVisited(index map[string]Valve, visited map[string]bool) bool {
	for key, _ := range index {
		if index[key].rate > 0 && !visited[key] {
			return false
		}
	}

	return true
}

func bruteForce(origin string, index map[string]Valve, distances map[string][]Distance, availableMinutes int, accPressure int, fullPath []string, visited map[string]bool) (int, []string) {

	if availableMinutes <= 0 || allVisited(index, visited) {
		return accPressure, fullPath
	}

	reachableNodes := distances[origin]
	//fmt.Println("From", origin, "walking", len(reachableNodes), "nodes")

	maxPressure := 0
	bestPath := []string{}

	for _, reachableNode := range reachableNodes {
		if !visited[reachableNode.dest] {

			// 1 minute for each node we visit to read destination
			minutesSpent := len(reachableNode.path)
			pressureAcc := 0

			node := index[reachableNode.dest]

			// We do not open valves for nodes already visited
			// The only have the valve open the first time
			if node.rate > 0 {
				// 1m Opening the valve
				minutesSpent++
				pressureAcc = node.rate * (availableMinutes - minutesSpent)
			}
			//}
			visited[node.id] = true
			pressure, path := bruteForce(reachableNode.dest, index, distances, availableMinutes-minutesSpent, accPressure+pressureAcc, append(fullPath, reachableNode.path...), visited)
			visited[node.id] = false

			if pressure > maxPressure {
				maxPressure = pressure
				bestPath = path
			}
		}

	}

	//fmt.Println("Current max:", maxPressure, bestPath)
	return maxPressure, bestPath
}

func main() {

	data := input.GetLines(2022, 16, "input.txt")

	valves := array.Map(data, getValve)

	index := makeValveMap(valves)

	fmt.Println(index)

	distances := map[string][]Distance{}

	for _, valve := range valves {
		distances[valve.id] = []Distance{}
	}

	// We only care about destinations that bring value
	for _, valve1 := range valves {
		for _, valve2 := range valves {
			if valve1.id != valve2.id && valve2.rate > 0 {
				minCost, minPath, _ := calculateMinDistance(valve1.id, valve2.id, index, 0, []string{})
				distances[valve1.id] = append(distances[valve1.id], Distance{dest: valve2.id, path: minPath, cost: minCost})

			}

		}

	}

	printDistances(distances)
	fmt.Println(len(distances))
	visited := map[string]bool{}
	visited["AA"] = true

	fmt.Println("---")
	result, path := bruteForce("AA", index, distances, 30, 0, []string{"AA"}, visited)
	fmt.Println("---")

	fmt.Println(result, path)

}
