package main

import (
	"advent/utils/array"
	"fmt"
	"math"
	"sort"
)

type Cell struct {
	row int
	col int
}

type Node struct {
	cell   Cell
	g      int
	h      float64
	parent *Node
}

func getNeighbours(cell Cell) []Cell {
	return []Cell{
		{row: cell.row - 1, col: cell.col - 1},
		{row: cell.row - 1, col: cell.col},
		{row: cell.row - 1, col: cell.col + 1},
		{row: cell.row, col: cell.col - 1},
		{row: cell.row, col: cell.col + 1},
		{row: cell.row + 1, col: cell.col + 1},
		{row: cell.row + 1, col: cell.col},
		{row: cell.row + 1, col: cell.col + 1},
	}
}

func distance(c1, c2 Cell) float64 {
	return math.Sqrt(math.Pow(float64(c2.row)-float64(c1.row), 2) +
		math.Pow(float64(c2.col)-float64(c1.col), 2))
}

func findIndex(node Node, nodes []Node) int {
	for idx, node2 := range nodes {
		if node.cell.row == node2.cell.row && node.cell.col == node2.cell.col {
			return idx
		}
	}

	return -1
}

func aStar(start, goal Cell) Node {
	pendingNodes := []Node{{cell: start}}
	closedNodes := []Node{}

	for {
		current := pendingNodes[0]
		pendingNodes = pendingNodes[1:]

		neighbours := getNeighbours(current.cell)

		candidates := array.Map(neighbours, func(cell Cell) Node {
			return Node{
				cell:   cell,
				parent: &current,
				g:      current.g + 1,
				h:      distance(cell, goal),
			}
		})

		// If we found the goal, return it
		for _, candidate := range candidates {
			if candidate.cell.col == goal.col && candidate.cell.row == goal.row {
				return candidate
			}
		}

		for _, candidate := range candidates {
			closedIndex := findIndex(candidate, closedNodes)
			pendingIndex := findIndex(candidate, pendingNodes)

			// If it is not closed, see if we need to add it to the
			// pending list
			if closedIndex == -1 {
				// If it is not pending we add it,
				// If it is pending already we only
				// overwrite it if the candidate is better than the current
				// candidate
				if pendingIndex == -1 {
					pendingNodes = append(pendingNodes, candidate)
				} else {
					currentPending := pendingNodes[pendingIndex]

					if float64(currentPending.g)+currentPending.h > float64(candidate.g)+candidate.h {
						pendingNodes[pendingIndex] = candidate
					}
				}
			}
		}
		closedNodes = append(closedNodes, current)

		// Sort open list so the min f is at the beginning of the list
		sort.SliceStable(pendingNodes, func(i, j int) bool {
			f1 := pendingNodes[i].g + int(pendingNodes[i].h)
			f2 := pendingNodes[j].g + int(pendingNodes[j].h)
			return f1 < f2
		})
	}

}
func main() {
	start := Cell{row: 0, col: 0}
	goal := Cell{row: 99, col: 99}

	resolvedGoal := aStar(start, goal)

	fmt.Println(resolvedGoal)
	for resolvedGoal.parent != nil {
		fmt.Println(resolvedGoal.cell)
		resolvedGoal = *resolvedGoal.parent
	}
	fmt.Println(resolvedGoal.cell)

}
