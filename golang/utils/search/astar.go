package search

import (
	"advent/utils/array"
	"fmt"
	"math"
	"sort"
)

type Node struct {
	g      float64
	h      float64
	f      float64
	node   [2]int
	parent *Node
}

func reconstructPath(node Node) [][2]int {
	path := [][2]int{node.node}

	for node.parent != nil {
		path = append(path, node.parent.node)
		node = *node.parent
	}

	return path
}

func square(n int) int {
	return n * n
}

func calculateDistance(node1 [2]int, node2 [2]int) float64 {
	distance := math.Sqrt(float64(square(node2[0]-node1[0]) + square(node2[1]-node1[1])))

	return distance
}

func findIndex(list []Node, node Node) int {
	for idx, n := range list {
		if isSameCell(n.node, node.node) {
			return idx
		}
	}

	return -1
}

func isSameCell(n1, n2 [2]int) bool {
	return n1[0] == n2[0] && n1[1] == n2[1]
}

// // TODO - Use any if possible....
// // we do not really care about the contents of the grid
// func AStar[T any](start, goal [2]int, grid [][]T, getNeighbours func(node [2]int, grid [][]T) [][2]int) [][2]int {

// 	// List of nodes pending to be analyzed
// 	openList := []Node{{
// 		node: start,
// 		g:    0,
// 		h:    0,
// 		f:    0,
// 	}}

// 	// List of nodes already visited
// 	// Should not be visited again
// 	closedList := []Node{}

// 	for len(openList) > 0 {

// 		// Pop node with lowest f (cost)
// 		currentNode := openList[0]
// 		openList = openList[1:]

// 		// Get neighbours
// 		neighbours := getNeighbours(currentNode.node, grid)

// 		successors := array.Map(neighbours, func(n [2]int) Node { return Node{node: n, parent: &currentNode} })

// 		// Compute successors
// 		for _, successor := range successors {

// 			// We found the goal
// 			if isSameCell(successor.node, goal) {
// 				return reconstructPath(successor)
// 			}

// 			// Compute G: distance to origin
// 			//successor.g = currentNode.g + calculateDistance(successor.node, currentNode.node)

// 			successor.g = currentNode.g + 1

// 			// Compute H: distance to goal
// 			// Using simple diagonal distance
// 			successor.h = calculateDistance(successor.node, goal)
// 			if math.IsNaN(successor.h) {
// 				fmt.Println("NaN:", successor.node, goal)
// 			}
// 			// Smaller is best.
// 			// It means we are closer to the goal and closer to the origin
// 			successor.f = successor.h + successor.g
// 			//fmt.Println(successor)

// 			// If the open list already has this node but with less f
// 			// (we reached it in a shorter path) we do not add this
// 			// If we find it with higher f, we replace it.
// 			// if it is not found, we can add it
// 			indexOfSuccessor := findIndex(openList, successor)
// 			indexOfClosedSuccessor := findIndex(closedList, successor)

// 			if indexOfClosedSuccessor == -1 { //|| closedList[indexOfClosedSuccessor].f > successor.f {
// 				if indexOfSuccessor == -1 {
// 					openList = append(openList, successor)
// 				} else if openList[indexOfSuccessor].f < successor.f {
// 					openList[indexOfSuccessor] = successor
// 				}
// 			}
// 		}

// 		//fmt.Println("old:", array.Map(openList, func(n Node) float64 { return n.f }))

// 		// Sort open list so the min f is at the beginning of the list
// 		sort.SliceStable(openList, func(i, j int) bool {
// 			return openList[i].f < openList[j].f
// 		})

// 		//fmt.Println("new:", array.Map(openList, func(n Node) float64 { return n.f }))

// 		closedList = append(closedList, currentNode)
// 	}

// 	return nil

// }

type Coords [2]int

type CellCost [3]int

type getNeighbours[T any] func(current Coords, grid [][]T, path []Coords) []CellCost

func AStar[T any](start, goal Coords, grid [][]T, getNeighbours getNeighbours) [][2]int {

	// List of nodes pending to be analyzed
	openList := []Node{{
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
