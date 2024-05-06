package board

import (
	"math/rand/v2"
)

// boardState := [5][6]int{
// 	{1,1,1,1,1,1,1},
// 	{1,1,1,1,1,1,1},
// 	{1,1,1,1,1,1,1},
// 	{1,1,1,1,1,1,1},
// 	{1,0,1,1,1,1,1},
// 	{1,1,1,1,1,1,1}
// }

func DeadState(w int, h int) [][]int {
	boardState := make([][]int, h)
	for i := range boardState {
		boardState[i] = make([]int, w)
	}
	return boardState
}

/*
RandomState takes a width and height and returns a 2D array
of ints either 1, to represent an alive cell , or 0, to represent a dead cell
*/
func RandomState(w int, h int) [][]int {
	var state [][]int = DeadState(w, h)
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			randomFloat := rand.Float64()
			if randomFloat >= 0.5 {
				state[i][j] = 0
			} else {
				state[i][j] = 1
			}
		}
	}
	return state
}

/*
Any live cell with 0 or 1 live neighbors becomes dead, because of underpopulation
any live cell with 2 or 3 live neighbors stays alive, because its neighborhood is correct
any live cell with more than 3 live neighbors becomes dead, because of overpopulation
any dead cell with exactly 3 live neighbors becomes alive by reproduction
*/
func CalculateNewBoard(board [][]int) [][]int {
	newBoard := make([][]int, len(board))
	for i := range newBoard {
		newBoard[i] = make([]int, len(board[i]))
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			neighbors := aliveNeighbors(board, i, j)
			if neighbors == 1 || neighbors == 0 {
				newBoard[i][j] = 0
			}
			if neighbors == 2 || neighbors == 3 {
				continue
			}
			if neighbors > 3 {
				newBoard[i][j] = 0
			}
			if board[i][j] == 0 && neighbors == 3 {
				newBoard[i][j] = 1
			}
		}
	}

	return newBoard
}

// this is aweful and likely the reason for the error. Fix this.
func aliveNeighbors(board [][]int, y, x int) int {
	var neighborArray = [][]int{
		{0, 1},
		{1, -1},
		{0, -1},
		{1, -1},
		{-1, 0},
		{1, 0},
		{-1, 1},
		{0, 1},
		{1, 1},
	}
	var count int = 0

	for _, n := range neighborArray {
		row := y + n[0]
		col := x + n[1]

		if row < 0 || row >= len(board) {
			continue
		}
		if col < 0 || col >= len(board) {
			continue
		}
		if board[row][col] == 1 {
			count++
		}
	}
	return count
}
