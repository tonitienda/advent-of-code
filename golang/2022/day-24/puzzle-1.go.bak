package main

import (
	"advent/utils/array"
	"advent/utils/input"
	"fmt"
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

func bfsPoi(start, goal Cell, cycle int) int {
	pendingNodes := []Cell{start}
	nextPendingNodes := map[Cell]bool{}

	minute := 0

	for {
		minute++
		for _, node := range pendingNodes {
			neighbours := findNeighboursAsOf(node, minute, cycle)

			for _, n := range neighbours {
				// We look for the cel that is over the goal
				// since the goal is in the margins and we will never
				// get it as a neighbour. And because we can only reach
				// the goal by going down from the cel on top
				if n.row == goal.row-1 && n.col == goal.col {
					// We would reach the goal in the next movement
					return minute + 1
				}

				nextPendingNodes[n] = true
			}
		}

		pendingNodes = []Cell{}

		for node, _ := range nextPendingNodes {
			pendingNodes = append(pendingNodes, node)
		}

		nextPendingNodes = map[Cell]bool{}
		//fmt.Printf("%d minutes, next pending nodes: %d\n", minute, len(pendingNodes))
	}
}

func main() {
	execType := "input"
	board := array.Map(input.GetLines(2022, 24, execType+".txt"), getRow)

	//fmt.Println(board)
	entrance := Cell{row: 0, col: 1}
	exit := Cell{row: len(board) - 1, col: len(board[0]) - 2}

	fmt.Println("E", entrance, "S", exit)

	boardCache[0] = board

	//fmt.Println(len(board)-2, len(board[0])-2)
	cycle := (len(board) - 2) * (len(board[0]) - 2)

	// TODO - calculate lcm of width and height of the board
	cycle = 600

	minutes := bfsPoi(entrance, exit, cycle)

	fmt.Println(minutes)

}
