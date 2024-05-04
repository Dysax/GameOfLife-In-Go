package board

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
