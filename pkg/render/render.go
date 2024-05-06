package render

import (
	"fmt"
)

// draws the entire board
func RenderBoard(array [][]int) {
	// top boarder
	for k := 0; k < len(array[0])+2; k++ {
		fmt.Printf("═")
	}
	fmt.Println()

	for i := 0; i < len(array); i++ {
		fmt.Printf("║")
		DrawLine(array[i])
		fmt.Printf("║\n")
	}

	for k := 0; k < len(array[0])+2; k++ {
		fmt.Printf("═")
	}
	fmt.Println()
}

// draws an individual sub array
func DrawLine(array []int) {

	for i := 0; i < len(array); i++ {
		if array[i] == 1 {
			fmt.Printf("█")
		} else {
			fmt.Printf(" ")
		}
	}
}
