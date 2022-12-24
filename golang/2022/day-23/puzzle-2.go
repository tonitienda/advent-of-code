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

func northIsFree(elf Elf, board Board) bool {
	ne := Elf{row: elf.row - 1, col: elf.col + 1}
	n := Elf{row: elf.row - 1, col: elf.col}
	nw := Elf{row: elf.row - 1, col: elf.col - 1}

	return !(board[ne] || board[n] || board[nw])
}

func getNorthLocation(elf Elf) Elf {
	return Elf{row: elf.row - 1, col: elf.col}
}

func eastIsFree(elf Elf, board Board) bool {
	ne := Elf{row: elf.row - 1, col: elf.col + 1}
	e := Elf{row: elf.row, col: elf.col + 1}
	se := Elf{row: elf.row + 1, col: elf.col + 1}

	return !(board[ne] || board[e] || board[se])
}

func getEastLocation(elf Elf) Elf {
	return Elf{row: elf.row, col: elf.col + 1}
}

func westIsFree(elf Elf, board Board) bool {
	nw := Elf{row: elf.row - 1, col: elf.col - 1}
	w := Elf{row: elf.row, col: elf.col - 1}
	sw := Elf{row: elf.row + 1, col: elf.col - 1}

	return !(board[nw] || board[w] || board[sw])
}

func noOneAround(elf Elf, board Board) bool {
	return northIsFree(elf, board) && eastIsFree(elf, board) && westIsFree(elf, board) && southIsFree(elf, board)
}

func getWestLocation(elf Elf) Elf {
	return Elf{row: elf.row, col: elf.col - 1}
}

func southIsFree(elf Elf, board Board) bool {
	se := Elf{row: elf.row + 1, col: elf.col + 1}
	s := Elf{row: elf.row + 1, col: elf.col}
	sw := Elf{row: elf.row + 1, col: elf.col - 1}

	return !(board[se] || board[s] || board[sw])
}

func getSouthLocation(elf Elf) Elf {
	return Elf{row: elf.row + 1, col: elf.col}
}

type Elf struct {
	row int
	col int
}

type Board = map[Elf]bool

type NextMoves = map[Elf][]Elf

func printBoard(board Board) {
	minRow := 9999
	maxRow := 0
	minCol := 9999
	maxCol := 0
	for elf, _ := range board {
		if elf.row > maxRow {
			maxRow = elf.row
		}
		if elf.row < minRow {
			minRow = elf.row
		}

		if elf.col > maxCol {
			maxCol = elf.col
		}
		if elf.col < minCol {
			minCol = elf.col
		}
	}

	for row := minRow; row <= maxRow; row++ {
		for col := minCol; col <= maxCol; col++ {
			if board[Elf{row: row, col: col}] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func main() {

	checks := []func(Elf, Board) bool{northIsFree, southIsFree, westIsFree, eastIsFree}
	updates := []func(Elf) Elf{getNorthLocation, getSouthLocation, getWestLocation, getEastLocation}

	execType := "input"
	data := array.Map(input.GetLines(2022, 23, execType+".txt"), func(line string) []bool { return array.Map(strings.Split(line, ""), StrToBool) })

	board := Board{}

	for row, _ := range data {
		for col, _ := range data[row] {
			if data[row][col] {
				board[Elf{row: row, col: col}] = true
			}
		}
	}

	fmt.Printf("\n== Initial state ==\n")
	printBoard(board)
	round := 0
	for {
		round++
		elfMoved := false
		nextMove := NextMoves{}

		for elf, _ := range board {
			if !noOneAround(elf, board) {
				for idx, check := range checks {
					if check(elf, board) {
						dest := updates[idx](elf)
						nextMove[dest] = append(nextMove[dest], elf)
						break
					}
				}
			}
		}

		for _, elves := range nextMove {
			if len(elves) == 1 {
				elfMoved = true
				elf := Elf{row: elves[0].row, col: elves[0].col}
				delete(board, elf)
			}
		}

		if !elfMoved {
			break
		}

		for dest, elves := range nextMove {
			if len(elves) == 1 {
				board[dest] = true
			}
		}

		// fmt.Printf("\n== End of Round %d ==\n", i+1)
		// printBoard(board)

		check := checks[0]
		checks = append(checks[1:], check)

		update := updates[0]
		updates = append(updates[1:], update)
	}

	minRow := 9999
	maxRow := 0
	minCol := 9999
	maxCol := 0
	for elf, _ := range board {
		if elf.row > maxRow {
			maxRow = elf.row
		}
		if elf.row < minRow {
			minRow = elf.row
		}

		if elf.col > maxCol {
			maxCol = elf.col
		}
		if elf.col < minCol {
			minCol = elf.col
		}
	}

	fmt.Println(minRow, maxRow, minCol, maxCol)

	area := ((maxRow - minRow) + 1) * ((maxCol - minCol) + 1)
	fmt.Println("area", area)

	elves := 0
	for _, v := range board {
		if v {
			elves++
		}
	}
	fmt.Println("elves", elves)

	fmt.Println(area - elves)
	fmt.Println("Total rounds:", round)

	// board := Board{}

	// elf1 := Elf{row: 10, col: 10}

	// elf2 := Elf{row: 10, col: 10}

	// board[elf1] = true

	// fmt.Println(board[elf1])

	// fmt.Println(board[elf2])

}
