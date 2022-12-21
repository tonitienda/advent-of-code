package main

import (
	"advent/utils/array"
	"advent/utils/funct"
	"advent/utils/input"
	"advent/utils/output"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func StrToInt(str string) int {
	return funct.GetValue(strconv.Atoi(str))
}

type Valve struct {
	id    string
	rate  int
	edges []Edge
}

type Edge struct {
	destId string
	cost   int
}

func (v Valve) String() string {
	return fmt.Sprintf("%s (%d) -> %s", v.id, v.rate, v.edges)

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
		id:    id,
		rate:  StrToInt(rate),
		edges: array.Map(connections, func(conn string) Edge { return Edge{destId: conn, cost: 1} }),
	}
}

func calculateMinDistance(valve1, valve2 string, index map[string]Valve, cost int, path []string) (int, []string, bool) {
	//fmt.Println(valve1, "==", valve2)
	if valve1 == valve2 {
		return cost, path, true
	}

	valve := index[valve1]

	//fmt.Println(tabs(len(path)), valve1, "=>", valve2, "children", valve.children, ":", path)

	children := array.Map(valve.edges, func(e Edge) string { return e.destId })
	if includes(children, valve2) {
		return cost + 1, append(path, valve2), true
	}

	minCost := int(math.Inf(1))
	minPath := []string{}

	atLeastOnePathFound := false
	for _, edge := range valve.edges {

		// Avoid cycles
		if includes(path, edge.destId) {
			continue
		}

		cost, newPath, ok := calculateMinDistance(edge.destId, valve2, index, cost+1, append(path, edge.destId))
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

func allVisited(index map[string]Valve, visited map[string]bool) bool {
	for key, _ := range index {
		if !visited[key] {
			return false
		}
	}

	return true
}

var iterations int = 0

func bruteForce(origin string, index map[string]Valve, availableMinutes int, accPressure int, visited map[string]bool) int {
	iterations++
	if availableMinutes <= 0 || allVisited(index, visited) {
		return accPressure
	}

	valve := index[origin]
	edges := valve.edges

	//fmt.Println("From", origin, "walking", len(reachableNodes), "nodes")

	maxPressure := 0

	for _, edge := range edges {
		if !visited[edge.destId] {

			// 1 minute for each node we visit to read destination
			minutesSpent := edge.cost
			pressureAcc := 0

			destNode := index[edge.destId]

			// This should always be true, since we removed the nodes that do not have value
			if destNode.rate > 0 {
				// 1m Opening the valve
				minutesSpent++
				pressureAcc = destNode.rate * (availableMinutes - minutesSpent)
			}
			//}
			visited[destNode.id] = true
			pressure := bruteForce(destNode.id, index, availableMinutes-minutesSpent, accPressure+pressureAcc, visited)
			visited[destNode.id] = false

			if pressure > maxPressure {
				maxPressure = pressure
			}
		}

	}

	//fmt.Println("Current max:", maxPressure, bestPath)
	return maxPressure
}

func getMarkDown(valves []Valve, index map[string]Valve) string {
	contents := "```mermaid\nflowchart LR\n"

	printed := map[string]bool{}
	for _, valve := range valves {
		for _, edge := range valve.edges {
			reverseEdgeId := edge.destId + "_" + valve.id
			edgeId := valve.id + "_" + edge.destId

			if !printed[reverseEdgeId] {
				contents += valve.id + " -- " + strconv.Itoa(edge.cost) + " --- " + edge.destId + "\n"
				printed[edgeId] = true
			}
		}
	}

	contents += "\n```\n\n"

	contents += "Total nodes: " + strconv.Itoa(len(valves)) + ".\n"

	return contents
}

func main() {

	execType := "test"
	data := input.GetLines(2022, 16, execType+".txt")

	valves := array.Map(data, getValve)

	index := makeValveMap(valves)

	fmt.Println(index)

	valves2 := []Valve{}

	// We only care about destinations that bring value
	for _, valve1 := range valves {
		if valve1.rate > 0 || valve1.id == "AA" {
			valve := Valve{id: valve1.id, rate: valve1.rate, edges: []Edge{}}
			for _, valve2 := range valves {
				if valve1.id != valve2.id && valve2.rate > 0 {
					_, minPath, _ := calculateMinDistance(valve1.id, valve2.id, index, 0, []string{})
					valve.edges = append(valve.edges, Edge{destId: valve2.id, cost: len(minPath)})
				}

			}
			valves2 = append(valves2, valve)
		}
	}

	originalValves := getMarkDown(valves, index)
	calculatedValves := getMarkDown(valves2, index)

	document := "### Original\n\n" + originalValves + "\n\n### New\n\n" + calculatedValves

	output.WriteContents(2022, 16, execType+".md", document)

	visited := map[string]bool{}
	visited["AA"] = true

	index2 := makeValveMap(valves2)

	fmt.Println("---")
	result := bruteForce("AA", index2, 30, 0, visited)
	fmt.Println("---")

	fmt.Println(result, "(", iterations, ")")

}
