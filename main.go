package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"time"

	"github.com/Dysax/GameOfLife/pkg/board"
	"github.com/Dysax/GameOfLife/pkg/render"
	"github.com/eiannone/keyboard"
)

func clearConsole() {
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func main() {
	clearConsole()
	err := keyboard.Open()
	if err != nil {
		panic(err)
	}
	defer keyboard.Close()

	ticker := time.NewTicker(33 * time.Millisecond)
	//ticker := time.NewTicker(time.Millisecond)
	done := make(chan bool)

	w, h := 80, 20

	boardState := board.RandomState(w, h)

	// for _, row := range boardState {
	// 	fmt.Println(row)
	// }
	// fmt.Println(len(boardState[0]))
	// fmt.Printf("board state array item %d\n", boardState[0][0])
	render.RenderBoard(boardState)
	//time.Sleep(time.Second)
	//newBoard := board.CalculateNewBoard(boardState)
	//render.RenderBoard(newBoard)

	go func() {
		for {
			select {
			case <-done:
				return
			case <-ticker.C:
				clearConsole()
				boardState = board.CalculateNewBoard(boardState)
				render.RenderBoard(boardState)
			}
		}
	}()

	go func() {
		for {
			char, key, err := keyboard.GetKey()
			if err != nil {
				done <- true
				return
			}
			if char == 'q' || key == keyboard.KeyEsc {
				done <- true
				return
			}
		}
	}()

	//time.Sleep(2500 * time.Millisecond)
	<-done
	ticker.Stop()
	fmt.Println("game over")

}
