// Game Of Life - Entire Game is in Here
package main

import (
	"fmt"
	"time"
)

const (
	boardWidth  int = 10
	boardHeight     = 10
)

const (
	dead  string = "#"
	alive        = ">"
)

type board = [boardWidth][boardHeight]string

var state = board{
	{dead, dead, dead, dead, dead, dead, dead, dead, dead, dead},
	{dead, dead, dead, dead, dead, dead, dead, dead, dead, dead},
	{dead, dead, dead, dead, dead, dead, dead, dead, dead, dead},
	{dead, dead, alive, alive, dead, dead, dead, dead, dead, dead},
	{dead, dead, alive, alive, dead, dead, dead, dead, dead, dead},
	{dead, dead, dead, dead, alive, alive, dead, dead, dead, dead},
	{dead, dead, dead, dead, alive, alive, dead, dead, dead, dead},
	{dead, dead, dead, dead, dead, dead, dead, dead, dead, dead},
	{dead, dead, dead, dead, dead, dead, dead, dead, dead, dead},
	{dead, dead, dead, dead, dead, dead, dead, dead, dead, dead},
}

func main() {
	for true {
		drawState(state)
		state = updateState(state)
		time.Sleep(60 * time.Millisecond)
	}
}

func drawState(state board) {
	for _, values := range state {
		for _, cell := range values {
			fmt.Print(cell)
		}
		fmt.Print("\n")
	}
}

func updateState(state board) board {
	for row, values := range state {
		for col := range values {
			nAlive := countNeighbors(row, col, state)
			state[row][col] = updateCellState(nAlive, state[row][col])
		}
	}

	return state
}

func countNeighbors(row int, col int, state board) int {
	nAlive := 0
	if row+1 < boardHeight {
		if state[row+1][col] != dead {
			nAlive++
		}
	}

	if row-1 >= 0 {
		if state[row-1][col] != dead {
			nAlive++
		}
	}

	if col+1 < boardWidth {
		if state[row][col+1] != dead {
			nAlive++
		}
	}

	if col-1 >= 0 {
		if state[row][col-1] != dead {
			nAlive++
		}
	}

	if row+1 < boardHeight && col+1 < boardWidth {
		if state[row+1][col+1] != dead {
			nAlive++
		}
	}

	if row-1 >= 0 && col-1 >= 0 {
		if state[row-1][col-1] != dead {
			nAlive++
		}
	}

	if row+1 < boardHeight && col-1 >= 0 {
		if state[row+1][col-1] != dead {
			nAlive++
		}
	}

	if row-1 >= 0 && col+1 < boardWidth {
		if state[row-1][col+1] != dead {
			nAlive++
		}
	}

	return nAlive

}

func updateCellState(nAlive int, cellState string) string {
	if cellState == alive {
		if nAlive < 2 || nAlive > 3 {
			return dead
		}

		return alive
	}
	if nAlive == 3 {
		return alive
	}

	return dead

}
