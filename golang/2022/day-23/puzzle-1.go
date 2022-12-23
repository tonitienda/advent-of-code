package main

import (
	"advent/utils/array"
	"advent/utils/input"
	"fmt"
	"strings"
)

func StrToBool(str string) bool {
	return str == "#"
}

func northIsFree(cell Cell, board Board) bool {
	ne := Cell{row: cell.row - 1, col: cell.col + 1}
	n := Cell{row: cell.row - 1, col: cell.col}
	nw := Cell{row: cell.row - 1, col: cell.col - 1}

	return !(board[ne] || board[n] || board[nw])
}

func getNorthLocation(cell Cell) Cell {
	return Cell{row: cell.row - 1, col: cell.col}
}

func eastIsFree(cell Cell, board Board) bool {
	ne := Cell{row: cell.row - 1, col: cell.col + 1}
	e := Cell{row: cell.row, col: cell.col + 1}
	se := Cell{row: cell.row + 2, col: cell.col + 1}

	return !(board[ne] || board[e] || board[se])
}

func getEastLocation(cell Cell) Cell {
	return Cell{row: cell.row, col: cell.col + 1}
}

func westIsFree(cell Cell, board Board) bool {
	nw := Cell{row: cell.row - 1, col: cell.col - 1}
	w := Cell{row: cell.row, col: cell.col - 1}
	sw := Cell{row: cell.row + 2, col: cell.col - 1}

	return !(board[nw] || board[w] || board[sw])
}

func getWestLocation(cell Cell) Cell {
	return Cell{row: cell.row, col: cell.col - 1}
}

func southIsFree(cell Cell, board Board) bool {
	se := Cell{row: cell.row + 1, col: cell.col + 1}
	s := Cell{row: cell.row + 1, col: cell.col}
	sw := Cell{row: cell.row + 1, col: cell.col - 1}

	return !(board[se] || board[s] || board[sw])
}

func getSouthLocation(cell Cell) Cell {
	return Cell{row: cell.row + 1, col: cell.col}
}

type Cell struct {
	row int
	col int
}

type Board = map[Cell]bool

type NextMoves = map[Cell][]Cell

func main() {

	checks := []func(Cell, Board) bool{northIsFree, southIsFree, westIsFree, eastIsFree}
	updates := []func(Cell) Cell{getNorthLocation, getSouthLocation, getWestLocation, getEastLocation}

	execType := "test"
	data := array.Map(input.GetLines(2022, 23, execType+".txt"), func(line string) []bool { return array.Map(strings.Split(line, ""), StrToBool) })

	board := Board{}

	for row, _ := range data {
		for col, _ := range data[row] {
			if data[row][col] {
				board[Cell{row: row, col: col}] = true
				fmt.Println("len board:", len(board))
			}
		}
	}

	fmt.Println("Original elves", len(board))

	for i := 0; i < 10; i++ {

		nextMove := NextMoves{}

		for cell, _ := range board {
			for idx, check := range checks {
				if check(cell, board) {
					dest := updates[idx](cell)
					fmt.Println(cell, dest)

					nextMove[dest] = append(nextMove[dest], cell)
					break
				}
			}
		}

		// fmt.Println(board)
		// fmt.Println(nextMove)

		for _, elves := range nextMove {
			if len(elves) == 1 {
				elf := Cell{row: elves[0].row, col: elves[0].col}
				fmt.Println("before", len(board))
				delete(board, elf)
				fmt.Println("after", len(board))
			}
		}

		for dest, elves := range nextMove {
			if len(elves) == 1 {
				board[dest] = true
			}
		}
	}

	minRow := 9999
	maxRow := 0
	minCol := 9999
	maxCol := 0
	for cell, _ := range board {
		if cell.row > maxRow {
			maxRow = cell.row
		}
		if cell.row < minRow {
			minRow = cell.row
		}

		if cell.col > maxCol {
			maxCol = cell.col
		}
		if cell.col < minCol {
			minCol = cell.col
		}
	}

	fmt.Println(minRow, minCol, maxRow, maxCol)

	area := (maxRow - minRow) * (maxCol - minCol)
	fmt.Println("area", area)

	elves := 0
	for _, v := range board {
		if v {
			elves++
		}
	}
	fmt.Println("elves", elves)

	// board := Board{}

	// cell1 := Cell{row: 10, col: 10}

	// cell2 := Cell{row: 10, col: 10}

	// board[cell1] = true

	// fmt.Println(board[cell1])

	// fmt.Println(board[cell2])

}
