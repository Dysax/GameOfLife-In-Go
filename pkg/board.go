package board

import (
	"fmt"
)

boardState := [5][6]int{
	{1,1,1,1,1,1,1},
	{1,1,1,1,1,1,1},
	{1,1,1,1,1,1,1},
	{1,1,1,1,1,1,1},
	{1,0,1,1,1,1,1},
	{1,1,1,1,1,1,1}
}
w := 5 int
h := 6 int

func deadState(w int, h int) [][]int {
	boardState := [w][h]int
	for i := 0; i <= w; i++ {
		for j := 0; j < h; j++ {

		}
	}
	return 
}
