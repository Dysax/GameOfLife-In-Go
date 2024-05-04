package main

import (
	"fmt"

	"github.com/Dysax/GameOfLife/pkg/board"
)

func main() {
	w, h := 5, 6
	boardState := board.DeadState(w, h)

	for _, row := range boardState {
		fmt.Println(row)
	}
}
