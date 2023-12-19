package y2023d17

import (
	"advent/utils/array"
	"advent/utils/funct"
	"advent/utils/input"
	"advent/utils/search"
	_ "embed"
	"fmt"
	"math"
	"sort"
	"strconv"
)

//go:embed test.txt
var test string

//go:embed input.txt
var actual string

func getNeighbours(current [2]int, cells [][]int) [][2]int {

	return [][2]int{}

}

type coords [2]int

type aStarNode struct {
	g    int
	h    int
	f    int
	node coords
}

func AStar(start, goal [2]int, grid [][]int) []coords {

	// List of nodes pending to be analyzed
	openList := []aStarNode{{
		node: start,
		g:    0,
		h:    0,
		f:    0,
	}}

	// List of nodes already visited
	// Should not be visited again
	closedList := []Node{}

	for len(openList) > 0 {

		// Pop node with lowest f (cost)
		currentNode := openList[0]
		openList = openList[1:]

		// Get neighbours
		neighbours := getNeighbours(currentNode.node, grid)

		successors := array.Map(neighbours, func(n [2]int) Node { return Node{node: n, parent: &currentNode} })

		// Compute successors
		for _, successor := range successors {

			// We found the goal
			if isSameCell(successor.node, goal) {
				return reconstructPath(successor)
			}

			// Compute G: distance to origin
			//successor.g = currentNode.g + calculateDistance(successor.node, currentNode.node)

			successor.g = currentNode.g + 1

			// Compute H: distance to goal
			// Using simple diagonal distance
			successor.h = calculateDistance(successor.node, goal)
			if math.IsNaN(successor.h) {
				fmt.Println("NaN:", successor.node, goal)
			}
			// Smaller is best.
			// It means we are closer to the goal and closer to the origin
			successor.f = successor.h + successor.g
			//fmt.Println(successor)

			// If the open list already has this node but with less f
			// (we reached it in a shorter path) we do not add this
			// If we find it with higher f, we replace it.
			// if it is not found, we can add it
			indexOfSuccessor := findIndex(openList, successor)
			indexOfClosedSuccessor := findIndex(closedList, successor)

			if indexOfClosedSuccessor == -1 { //|| closedList[indexOfClosedSuccessor].f > successor.f {
				if indexOfSuccessor == -1 {
					openList = append(openList, successor)
				} else if openList[indexOfSuccessor].f < successor.f {
					openList[indexOfSuccessor] = successor
				}
			}
		}

		//fmt.Println("old:", array.Map(openList, func(n Node) float64 { return n.f }))

		// Sort open list so the min f is at the beginning of the list
		sort.SliceStable(openList, func(i, j int) bool {
			return openList[i].f < openList[j].f
		})

		//fmt.Println("new:", array.Map(openList, func(n Node) float64 { return n.f }))

		closedList = append(closedList, currentNode)
	}

	return nil

}

func Run1() {
	cells := input.Get2DMatrix(test, "\n", "", func(s string) int { return funct.GetValue(strconv.Atoi(s)) })

	fmt.Println(cells)

	start := [2]int{0, 0}
	goal := [2]int{len(cells) - 1, len(cells[0]) - 1}

	path := search.AStar(start, goal, cells, getNeighbours)

	fmt.Println("Path", path)
}
