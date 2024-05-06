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

	//ticker := time.NewTicker(33 * time.Millisecond)
	//ticker := time.NewTicker(67 * time.Millisecond)
	ticker := time.NewTicker(80 * time.Millisecond)
	done := make(chan bool)

	w, h := 180, 60

	boardState := board.RandomState(w, h)
	render.RenderBoard(boardState)

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

	<-done
	ticker.Stop()
	fmt.Println("game over")

}
