package board

import (
	"math/rand/v2"
)

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
	Conway's Game of Life Rules:

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

	for y := 0; y < len(board); y++ {
		for x := 0; x < len(board[y]); x++ {
			newBoard[y][x] = board[y][x]
			aliveNeighbors := countAliveNeighbors(board, y, x)
			if aliveNeighbors < 2 {
				newBoard[y][x] = 0
			}
			if aliveNeighbors >= 4 {
				newBoard[y][x] = 0
			}
			if aliveNeighbors == 3 {
				newBoard[y][x] = 1
			}
		}
	}
	return newBoard
}

// contents of neighborArray are added to each x and y, then if the result is out of bounds skip past it
func countAliveNeighbors(board [][]int, y, x int) int {
	var neighborArray = [][]int{
		{-1, -1},
		{0, -1},
		{1, -1},
		{-1, 0},
		{1, 0},
		{-1, 1},
		{0, 1},
		{1, 1},
	}

	count := 0
	for _, n := range neighborArray {
		col := x + n[0]
		row := y + n[1]

		if row < 0 || row >= len(board) {
			continue
		}
		if col < 0 || col >= len(board[0]) {
			continue
		}
		if board[row][col] == 1 {
			count++
		}
	}
	return count
}
