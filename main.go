package main

import (
	"fmt"

	"github.com/Dysax/GameOfLife/pkg/board"
	"github.com/Dysax/GameOfLife/pkg/render"
)

func main() {
	w, h := 20, 10

	boardState := board.RandomState(w, h)

	for _, row := range boardState {
		fmt.Println(row)
	}
	// fmt.Println(len(boardState[0]))
	// fmt.Printf("board state array item %d\n", boardState[0][0])
	render.RenderBoard(boardState)
}
