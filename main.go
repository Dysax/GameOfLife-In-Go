package main

import (
	"fmt"

	"github.com/Dysax/GameOfLife/pkg/board"
)

func main() {
	w, h := 5, 5

	boardState := board.RandomState(w, h)

	for _, row := range boardState {
		fmt.Println(row)
	}
}
