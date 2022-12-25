package main

import (
	"advent/utils/array"
	"advent/utils/input"
	"fmt"
	"math"
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
	row    int
	col    int
	minute int
}

type Node struct {
	cell   Cell
	parent *Node
	g      float64
	h      float64
	f      float64
}

func (node Node) ToString() string {
	return fmt.Sprintf("{%d, %d}, m: %d", node.cell.row, node.cell.col, node.cell.minute)
}

var cachedNeighbours = map[int]map[Cell][]Cell{}

func findNeighboursAsOf(cell Cell, minutes, cycle int) []Cell {
	if cache, ok := cachedNeighbours[minutes]; ok {
		if value, ok2 := cache[cell]; ok2 {
			return value
		}
	} else {
		cachedNeighbours[minutes] = map[Cell][]Cell{}
	}

	board := getBoardAsOf(minutes, cycle)
	neighbours := findNeighbours(cell, board)

	withMinute := array.Map(neighbours, func(c Cell) Cell { c.minute = minutes + 1; return c })

	cachedNeighbours[minutes][cell] = withMinute

	return withMinute

}

func findNeighbours(cell Cell, board [][]int) []Cell {
	neighbours := []Cell{}

	// Order this by the direction that seems optimal
	// Down and right

	// down
	if cell.col > 0 && cell.col < len(board[cell.row]) {
		if cell.row < len(board)-2 && board[cell.row+1][cell.col] == 0 {
			neighbours = append(neighbours, Cell{row: cell.row + 1, col: cell.col})
		}
	}

	// right
	if cell.row > 0 && cell.row < len(board)-1 {
		if cell.col < len(board[0])-2 && board[cell.row][cell.col+1] == 0 {
			neighbours = append(neighbours, Cell{row: cell.row, col: cell.col + 1})
		}
	}

	// up
	if cell.col > 0 && cell.col < len(board[cell.row]) {
		if cell.row > 1 && board[cell.row-1][cell.col] == 0 {
			neighbours = append(neighbours, Cell{row: cell.row - 1, col: cell.col})
		}
	}

	// left
	if cell.row > 0 && cell.row < len(board)-1 {
		if cell.col > 1 && board[cell.row][cell.col-1] == 0 {
			neighbours = append(neighbours, Cell{row: cell.row, col: cell.col - 1})
		}
	}

	//self - We can wait in the same cell, for another minute
	// that can count as having itself as a neighbour
	if board[cell.row][cell.col] == 0 {
		neighbours = append(neighbours, Cell{row: cell.row, col: cell.col})
	}
	return neighbours
}

var boardCache = map[int][][]int{}

func getBoardAsOf(minute, cycle int) [][]int {
	if value, ok := boardCache[minute%cycle]; ok {
		return value
	}

	// Last minute cached
	keys := make([]int, 0, len(boardCache))
	for k := range boardCache {
		keys = append(keys, k)
	}

	lastMinute := array.Max(keys)
	board := boardCache[lastMinute]

	for i := lastMinute + 1; i <= minute%cycle; i++ {
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
				fmt.Print("E")
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
		if isSameCell(n.cell, node.cell) && n.cell.minute == node.cell.minute {
			return idx
		}
	}

	return -1
}

func bruteForce(start, goal Cell, cycle, minutes int, visited map[Cell]bool, path []Cell) (int, bool, []Cell) {
	currentCell := start

	// We can only get to the goal going down from the previous row
	// So if we reach {row-1, col} means that the next step is the goal
	if currentCell.row == goal.row-1 && currentCell.col == goal.col {
		return minutes + 1, true, append(path, goal)
	}

	neighbours := findNeighboursAsOf(currentCell, minutes, cycle)

	minMinutes := math.MaxInt
	found := false
	bestPath := []Cell{}

	for _, neighbour := range neighbours {

		// Check visited
		modCell := Cell{row: neighbour.row, col: neighbour.col, minute: neighbour.minute % cycle}

		if visited[modCell] {
			continue
		}

		visited[modCell] = true
		minutes2, ok, path2 := bruteForce(neighbour, goal, cycle, minutes+1, visited, append(path, neighbour))
		visited[modCell] = false

		if ok && minutes2 < minMinutes {
			found = true
			minMinutes = minutes2
			bestPath = path2
		}
	}

	return minMinutes, found, bestPath

}

// func Astar(start Cell, goal Cell, board [][]int, minutes int) (Node, int, bool) {
// 	// List of nodes pending to be analyzed

// 	openList := []Node{{
// 		cell: start,
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
// 		//minutes = currentNode.minute

// 		fmt.Println("minutes", minutes, "currentNode", currentNode, "openlist", len(openList), "closedList", len(closedList))

// 		neighbours := findNeighboursAsOf(currentNode.cell, minutes)

// 		successors := array.Map(neighbours, func(n Cell) Node {
// 			return Node{node: n, parent: &currentNode}
// 		})

// 		// fmt.Println("s", neighbours)
// 		// Compute successors
// 		for _, successor := range successors {

// 			// We found the goal
// 			if isSameCell(successor.node, goal) {
// 				return successor, minutes, true
// 			}

// 			// Compute G: distance to origin
// 			// if isSameCell(successor.node, currentNode.node) {
// 			// 	successor.g = currentNode.g
// 			// } else {
// 			successor.g = currentNode.g + 1
// 			// }

// 			// Compute H: distance to goal
// 			// Using simple diagonal distance
// 			successor.h = calculateDistance(successor.node, goal)
// 			// if math.IsNaN(successor.h) {
// 			// 	fmt.Println("NaN:", successor.node, goal)
// 			// }
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

// 			if indexOfClosedSuccessor == -1 {
// 				if indexOfSuccessor == -1 {
// 					openList = append(openList, successor)
// 				} else if openList[indexOfSuccessor].f < successor.f {
// 					openList[indexOfSuccessor] = successor
// 				}
// 			}
// 		}

// 		// Sort open list so the min f is at the beginning of the list
// 		sort.SliceStable(openList, func(i, j int) bool {
// 			return openList[i].f < openList[j].f
// 		})

// 		closedList = append(closedList, currentNode)

// 	}

// 	return Node{}, 0, false

// }

func reverse(nodes []Node) []Node {
	for i, j := 0, len(nodes)-1; i < j; i, j = i+1, j-1 {
		nodes[i], nodes[j] = nodes[j], nodes[i]
	}
	return nodes

}

func main() {
	execType := "input"
	board := array.Map(input.GetLines(2022, 24, execType+".txt"), getRow)

	fmt.Println(board)
	entrance := Cell{row: 0, col: 1}
	exit := Cell{row: len(board) - 1, col: len(board[0]) - 2}

	fmt.Println("E", entrance, "S", exit)

	boardCache[0] = board

	// board1 := getBoardAsOf(0)

	// board2 := getBoardAsOf()

	// printBoard(board1, entrance)
	// printBoard(board2, entrance)
	// TODO - Optimize if we can factor the num of rows and columns
	fmt.Println(len(board)-2, len(board[0])-2)
	cycle := (len(board) - 2) * (len(board[0]) - 2)
	cycle = 600
	visitedNodes := map[Cell]bool{}
	//_, minutes, found := Astar(entrance, exit, board, 0)
	_, found, path := bruteForce(entrance, exit, cycle, 1, visitedNodes, []Cell{entrance})

	// for idx, cell := range path {
	// 	fmt.Println("Minute", idx)
	// 	printBoard(getBoardAsOf(idx), cell)
	// 	fmt.Println()
	// }

	// Path contains the initial state, but we only want the steps
	// so we need to ignore the "initial state"
	fmt.Println(len(path)-1, found)

	// nodes := []Node{}

	// for node.parent != nil {
	// 	nodes = append(nodes, node)
	// 	node = *node.parent

	// 	if node.parent == nil {
	// 		nodes = append(nodes, node)
	// 	}
	// }

	// for _, node := range reverse(nodes) {
	// 	fmt.Println("Minute", node.minute)
	// 	printBoard(getBoardAsOf(node.minute), node.node)
	// 	fmt.Println()
	// }

	// fmt.Println(minutes, found)

}
