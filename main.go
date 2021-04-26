package main

import (
	"sudoku-solver/sudoku"
)

func main() {
	board := sudoku.Create(9, 1)
	board.Display()
	board.Solve()
	board.Display()
}
