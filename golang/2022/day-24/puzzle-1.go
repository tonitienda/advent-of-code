package main

import (
	"advent/utils/array"
	"advent/utils/input"
	"fmt"
	"math"
	"sort"
	"strings"
)

var Values = map[string]int{
	"#": 0,
	"^": 1,
	">": 2,
	"v": 4,
	"<": 8,
	".": 0,
}

var InverseValues = map[int]string{
	0: ".",
	1: "^",
	2: ">",
	4: "v",
	8: "<",
}

func getValue(str string) int {
	return Values[str]
}

func getRow(line string) []int {
	values := strings.Split(line, "")

	return array.Map(values, getValue)
}

type Cell struct {
	row int
	col int
}

type Node struct {
	node   Cell
	parent *Node
	minute int
	g      float64
	h      float64
	f      float64
}

var cachedNeighbours = map[int]map[Cell][]Cell{}

func findNeighboursAsOf(cell Cell, minutes int) []Cell {
	if cache, ok := cachedNeighbours[minutes]; ok {
		if value, ok2 := cache[cell]; ok2 {
			return value
		}
	} else {
		cachedNeighbours[minutes] = map[Cell][]Cell{}
	}

	board := getBoardAsOf(minutes)
	neighbours := findNeighbours(cell, board)
	cachedNeighbours[minutes][cell] = neighbours

	return neighbours

}

func findNeighbours(cell Cell, board [][]int) []Cell {
	neighbours := []Cell{}

	// Order this by the direction that seems optimal
	// Down and right

	// down
	if cell.row < len(board)-1 && board[cell.row+1][cell.col] == 0 {
		neighbours = append(neighbours, Cell{row: cell.row + 1, col: cell.col})
	}

	// right
	if cell.col < len(board[0])-1 && board[cell.row][cell.col+1] == 0 {
		neighbours = append(neighbours, Cell{row: cell.row, col: cell.col + 1})
	}

	// up
	if cell.row > 0 && board[cell.row-1][cell.col] == 0 {
		neighbours = append(neighbours, Cell{row: cell.row - 1, col: cell.col})
	}

	// left
	if cell.col > 0 && board[cell.row][cell.col-1] == 0 {
		neighbours = append(neighbours, Cell{row: cell.row, col: cell.col - 1})
	}

	return neighbours
}

var boardCache = map[int][][]int{}

func getBoardAsOf(minute int) [][]int {
	// Last minute cached
	keys := make([]int, 0, len(boardCache))
	for k := range boardCache {
		keys = append(keys, k)
	}

	lastMinute := array.Max(keys)
	// fmt.Println("lastMinute", lastMinute, "minute", minute)
	if lastMinute >= minute {
		return boardCache[minute]
	}

	board := boardCache[lastMinute]

	for i := lastMinute + 1; i <= minute+1; i++ {
		board = updateBoard(board, i)
		boardCache[i] = board
	}

	return boardCache[minute]
}

func updateBoard(board [][]int, minutes int) [][]int {

	if value, ok := boardCache[minutes]; ok {
		return value
	}

	newBoard := [][]int{}

	for row := 0; row < len(board); row++ {
		newBoard = append(newBoard, []int{})

		for col := 0; col < len(board[row]); col++ {
			newBoard[row] = append(newBoard[row], 0)

			if row == 0 || row == len(board)-1 || col == 0 || col == len(board[row])-1 {
				continue
			}

			//fmt.Println("newBoard[row][col]:", row, col, newBoard[row][col])
			// up
			if row < len(board)-2 {
				//fmt.Println("up - board[row+1][col]", board[row+1][col], "=>", Values["^"]&board[row+1][col])
				newBoard[row][col] += Values["^"] & board[row+1][col]
			} else {
				newBoard[row][col] += Values["^"] & board[1][col]
			}

			// down
			if row > 1 {
				newBoard[row][col] += Values["v"] & board[row-1][col]
			} else {
				//fmt.Println("down - board[len(board)-2][col]", board[len(board)-2][col], "=>", Values["v"]&board[len(board)-2][col])
				newBoard[row][col] += Values["v"] & board[len(board)-2][col]
			}

			// left
			if col < len(board[row])-2 {
				//fmt.Println("left - board[row][col+1]", board[row][col+1], "=>", Values["<"]&board[row][col+1])
				newBoard[row][col] += Values["<"] & board[row][col+1]
			} else {
				newBoard[row][col] += Values["<"] & board[row][1]
			}

			// right
			if col > 1 {
				newBoard[row][col] += Values[">"] & board[row][col-1]
			} else {
				//fmt.Println("right - board[row][len(board[row])-2]", board[row][len(board[row])-2], "=>", Values[">"]&board[row][len(board[row])-2])
				newBoard[row][col] += Values[">"] & board[row][len(board[row])-2]
			}

		}
	}

	//fmt.Println(newBoard)
	//fmt.Println("---")

	return newBoard
}

func printBoard(board [][]int, currentNode Cell) {
	for row := 0; row < len(board); row++ {
		for col := 0; col < len(board[row]); col++ {

			if row == currentNode.row && col == currentNode.col {
				fmt.Print("😶")
				continue
			}

			if row == 0 || row == len(board)-1 || col == 0 || col == len(board[row])-1 {
				fmt.Print("#")
				continue
			}

			value, ok := InverseValues[board[row][col]]

			if ok {
				fmt.Print(value)
			} else {
				fmt.Print("?")

			}
		}
		fmt.Println()
	}
}

func isSameCell(n1, n2 Cell) bool {
	return n1.row == n2.row && n1.col == n2.col
}

func square(n int) int {
	return n * n
}

func calculateDistance(node1, node2 Cell) float64 {
	distance := math.Sqrt(float64(square(node2.row-node1.row) + square(node2.col-node1.col)))

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

func Astar(start Cell, goal Cell, board [][]int, minutes int) (int, bool) {
	// List of nodes pending to be analyzed

	openList := []Node{{
		node:   start,
		g:      0,
		h:      0,
		f:      0,
		minute: minutes,
	}}

	// List of nodes already visited
	// Should not be visited again
	closedList := []Node{}

	for len(openList) > 0 {
		// Pop node with lowest f (cost)
		currentNode := openList[0]
		minutes = currentNode.minute

		openList = openList[1:]

		fmt.Println(currentNode.node, "openList", len(openList))

		// Get neighbours
		// TODO - We can wait even if there are neighbours
		neighbours := findNeighboursAsOf(currentNode.node, minutes)
		minutes++
		for len(neighbours) == 0 {
			neighbours = findNeighboursAsOf(currentNode.node, minutes)
			minutes++
		}

		successors := array.Map(neighbours, func(n Cell) Node { return Node{node: n, parent: &currentNode, minute: minutes} })

		// Compute successors
		for _, successor := range successors {

			// We found the goal
			if isSameCell(successor.node, goal) {
				//return reconstructPath(successor)

				// TODO - Not sure about this
				return minutes, true
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

	return 0, false

}

// func findShortestPath(currentNode Cell, goal Cell, board [][]int, minutes int, visited map[int]map[Cell]bool) (int, bool) {
// 	fmt.Println("current node:", currentNode, "minutes", minutes)

// 	if currentNode.row == goal.row && currentNode.col == goal.col {
// 		fmt.Println("Arrived in ", minutes, "minutes")
// 		//panic("interruption")
// 		return minutes, true
// 	}

// 	// // Upper bound for minutes
// 	// if minutes > 1000 {
// 	// 	return 0, false
// 	// }
// 	if _, ok := visited[minutes]; !ok {
// 		visited[minutes] = map[Cell]bool{}
// 	}

// 	visited[minutes][currentNode] = true

// 	neighbours := findNeighbours(currentNode, board)
// 	newBoard := updateBoard(board, minutes)

// 	// A blizzard reached the expedition
// 	if newBoard[currentNode.row][currentNode.col] != 0 {
// 		fmt.Println("newBoard[currentNode.row][currentNode.col]", newBoard[currentNode.row][currentNode.col])
// 		fmt.Println(newBoard)
// 		printBoard(newBoard, currentNode)

// 		return 0, false
// 	}
// 	minutes++

// 	for len(neighbours) == 0 {
// 		neighbours = findNeighbours(currentNode, newBoard)
// 		newBoard = updateBoard(newBoard, minutes)
// 		minutes++
// 	}

// 	fmt.Println("Found neighbours")
// 	printBoard(newBoard, currentNode)
// 	fmt.Println("currentNode:", currentNode)
// 	fmt.Println("Neighbours:", neighbours)

// 	fmt.Println("Board value:", newBoard[neighbours[0].row][neighbours[0].col])
// 	fmt.Println()

// 	// If one of the neighbours is the Goal we do not need to visit the rest

// 	for _, neighbour := range neighbours {
// 		if neighbour.row == goal.row && neighbour.col == goal.col {
// 			return minutes + 1, true
// 		}
// 	}
// 	minMinutes := math.MaxInt
// 	atLeastOneVisited := false

// 	for _, neighbour := range neighbours {
// 		if !visited[minMinutes][neighbour] {
// 			takenMinutes, ok := findShortestPath(neighbour, goal, newBoard, minutes, visited)

// 			if ok {
// 				atLeastOneVisited = true

// 				if takenMinutes < minMinutes {
// 					minMinutes = takenMinutes
// 				}
// 			}
// 		}

// 	}
// 	return minMinutes, atLeastOneVisited
// }

func main() {
	execType := "test"
	board := array.Map(input.GetLines(2022, 24, execType+".txt"), getRow)

	fmt.Println(board)
	entrance := Cell{row: 0, col: 1}
	exit := Cell{row: len(board) - 1, col: len(board[0]) - 2}

	fmt.Println("E", entrance, "S", exit)

	// Visited means here, visit the same cell in the same minute. That will
	// lead to the same result
	//visited := map[int]map[Cell]bool{}

	// TODO - Not sure if this should be minute 1
	boardCache[0] = board

	//minutes, found := findShortestPath(entrance, exit, board, 0, visited)
	minutes, found := Astar(entrance, exit, board, 0)

	fmt.Println(minutes, found)

	// for i := 0; i < len(boardCache); i++ {
	// 	fmt.Println("Minute ", i)
	// 	printBoard(boardCache[i], Cell{row: 0, col: 1})
	// 	fmt.Println()
	// }

}
