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

	return newBoard
}

func aliveNeighbors(board [][]int, x, y int) int {
	count := 0
	if board[x][y] != 1 {
		return 0
	} else {
		if board[x][y] == board[x][y-1] {
			count++
		}
		if board[x][y] == board[x][y+1] {
			count++
		}
	}
	return count
}
